package main

import (
	"github.com/therecipe/qt/core"

	"github.com/shirou/dictee/dictionary"
)

const (
	TitleName = int(core.Qt__UserRole) + 1<<iota
	DictName
)

type TitleModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_      map[int]*core.QByteArray `property:"roles"`
	titles []*Title                 `property:"title"`

	_ func(*Title)  `slot:"addTitle"`
	_ func(row int) `slot:"removeTitle"`
}

type Title struct {
	core.QObject

	_ string `property:"titleName"` // if just "title", qtmoc hangs
	_ string `property:"dictName"`

	original dictionary.Title
}

func init() {
	Title_QRegisterMetaType()
}

func (m *TitleModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		TitleName: core.NewQByteArray2("titleName", len("titleName")),
		DictName:  core.NewQByteArray2("dictName", len("dictName")),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddTitle(m.addTitle)
	m.ConnectRemoveTitle(m.removeTitle)
}

func (m *TitleModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Title()) {
		return core.NewQVariant()
	}

	p := m.Title()[index.Row()]

	switch role {
	case TitleName:
		return core.NewQVariant14(p.TitleName())
	case DictName:
		return core.NewQVariant14(p.DictName())
	default:
		return core.NewQVariant()
	}
}

func (m *TitleModel) rowCount(parent *core.QModelIndex) int {
	return len(m.Title())
}

func (m *TitleModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *TitleModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *TitleModel) addTitle(p *Title) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Title()), len(m.Title()))
	m.SetTitle(append(m.Title(), p))
	m.EndInsertRows()
}

func (m *TitleModel) removeTitle(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetTitle(append(m.Title()[:row], m.Title()[row+1:]...))
	m.EndRemoveRows()
}

func (m *TitleModel) Get(index int) *Title {
	return m.Title()[index]
}
