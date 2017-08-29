//go:generate qtmoc
package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"

	"github.com/shirou/dictee"
	"github.com/shirou/dictee/dictionary"
)

var model *TitleModel
var dicteeSingle dictee.Dictee
var config *dictee.Config

type QmlBridge struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(text string) `signal:"onSearchInputFinished"`
	_ func()            `signal:"onBackButtonClicked"`
	_ func()            `signal:"onForwardButtonClicked"`
	_ func(index int)   `signal:"onTitleClicked"`
	_ func(text string) `signal:"setViewText"`

	titleCh chan *Title
}

func (bridge *QmlBridge) init() {
	bridge.ConnectOnSearchInputFinished(func(word string) {
		fmt.Println("go:", word)

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		ch := make(chan dictionary.Title, 100)
		go dicteeSingle.SearchAll(ctx, word, ch)
		for {
			select {
			case title, ok := <-ch:
				if !ok {
					return
				}
				p := NewTitle(nil)
				p.SetTitleName(title.Title)
				p.original = title
				bridge.titleCh <- p
			case <-ctx.Done():
				if ctx.Err() != nil {
					fmt.Println("ERR", ctx.Err())
					return
				}
			}
		}

	})
	bridge.ConnectOnForwardButtonClicked(func() {
		fmt.Println("go: forward")
	})
	bridge.ConnectOnBackButtonClicked(func() {
		fmt.Println("go: back")
	})
	bridge.ConnectOnTitleClicked(func(index int) {
		title := model.Get(index)
		fmt.Println("go: title", title.TitleName())
		original := title.original
		g, err := original.Dictionary.Get(original)
		if err != nil {
			fmt.Println(errors.Wrap(err, "dict get error"))
			return
		}
		fmt.Println(g.ToRich())
		bridge.SetViewText(g.ToRich())
	})
}

func main() {
	var err error
	config, dicteeSingle, err = dictee.ReadConfig("")
	if err != nil {
		panic(err)
	}

	//enable high dpi scaling
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)
	core.QCoreApplication_SetOrganizationName("QtProject")
	core.QCoreApplication_SetApplicationName("dictee")
	core.QCoreApplication_SetApplicationVersion("1.0.0")
	quickcontrols2.QQuickStyle_SetStyle("material")

	model = NewTitleModel(nil)

	//create a moc bridge
	bridge := NewQmlBridge(nil)
	bridge.titleCh = make(chan *Title, 100)

	var app = qml.NewQQmlApplicationEngine(nil)
	// register to QmlBridge struct to QML side.
	app.RootContext().SetContextProperty("QmlBridge", bridge)
	app.RootContext().SetContextProperty("TitleModel", model)

	app.Load(core.NewQUrl3(filepath.Join("qml/main.qml"), 0))
	//	app.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	go func() {
		for {
			select {
			case title, ok := <-bridge.titleCh:
				if !ok {
					fmt.Println("bridge.titleCh closed")
					return
				}
				model.AddTitle(title)
			}
		}
	}()

	gui.QGuiApplication_Exec()
}
