package dictee

import (
	"fmt"
	"log"

	"github.com/shirou/dictee/dictionary"
)

type Dictee struct {
	ConfigRoot   string                  `json:"-"`
	Dictionaries []dictionary.Dictionary `json:"dictionaries"`
}

func (d *Dictee) SearchDictonary() {

}

func (d *Dictee) Search(word string) {
	ch := make(chan dictionary.Title, 100)

	for _, dict := range d.Dictionaries {
		go func() {
			t, err := dict.Search(word)
			if err != nil {
				fmt.Errorf("search error:%v", err)
				return
			}
			for _, tt := range t {
				ch <- tt
			}
			close(ch)
		}()
	}

	for title := range ch {
		fmt.Println(title)
	}

}

func (d *Dictee) MakeIndex() {
	for _, dict := range d.Dictionaries {
		log.Printf("create index start: %v", dict)
		dict.MakeIndex(d.ConfigRoot)
	}
}
