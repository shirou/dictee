package main

import (
	"github.com/shirou/dictee/dictionary"
	"github.com/therecipe/qt/core"
)

const (
	DictTabName = int(core.Qt__UserRole) + 1<<iota
)

type DictTabModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*DictTab               `property:"dictTab"`

	_ func(*DictTab) `slot:"addDictTab"`
	_ func(row int)  `slot:"removeDictTab"`
}

type DictTab struct {
	core.QObject

	_ string `property:"dictName"`

	Dictionary dictionary.Dictionary
}

func init() {
	DictTab_QRegisterMetaType()
}

func (m *DictTabModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		DictName: core.NewQByteArray2("dictName", len("dictName")),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddDictTab(m.addDictTab)
	m.ConnectRemoveDictTab(m.removeDictTab)
}

func (m *DictTabModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.DictTab()) {
		return core.NewQVariant()
	}

	p := m.DictTab()[index.Row()]

	switch role {
	case DictName:
		return core.NewQVariant14(p.DictName())
	default:
		return core.NewQVariant()
	}
}

func (m *DictTabModel) rowCount(parent *core.QModelIndex) int {
	return len(m.DictTab())
}

func (m *DictTabModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *DictTabModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *DictTabModel) addDictTab(p *DictTab) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.DictTab()), len(m.DictTab()))
	m.SetDictTab(append(m.DictTab(), p))
	m.EndInsertRows()
}

func (m *DictTabModel) removeDictTab(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetDictTab(append(m.DictTab()[:row], m.DictTab()[row+1:]...))
	m.EndRemoveRows()
}

func (m *DictTabModel) Get(index int) *DictTab {
	return m.DictTab()[index]
}
