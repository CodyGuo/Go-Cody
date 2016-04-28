package main

import (
	"fmt"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type Foo struct {
	Index  int
	File   string
	Path   string
	Result string
	Remark string
}

type FooModel struct {
	walk.SortedReflectTableModelBase // This will use reflection to provide a sorted table model.
	items                            []*Foo
}

// We need to provide our model items like this.
func (m *FooModel) Items() interface{} {
	return m.items
}

func NewFooModel() *FooModel {
	m := new(FooModel)
	m.aa()
	return m
}

func main() {
	model := NewFooModel()
	var tv *walk.TableView
	var mw *walk.MainWindow

	MainWindow{
		Title:    "aa",
		AssignTo: &mw,
		Size:     Size{800, 800},
		Layout:   VBox{},
		Children: []Widget{
			PushButton{
				Text:      "Reset Rows",
				OnClicked: model.aa,
			},
			TableView{
				AssignTo: &tv,
				Columns: []TableViewColumn{
					{Title: "序号"}, // Changed Title to Name.
					{Title: "文件"},
					{Title: "文件路径"},
					{Title: "进度"},
					{Title: "备注"},
				},
				Model: model,
				OnCurrentIndexChanged: func() {
					fmt.Println("HELLO")
				},
			},
		},
	}.Run()
}

func (m *FooModel) aa() {
	m.items = make([]*Foo, 10)
	for i := range m.items {
		m.items[i] = &Foo{
			Index:  i,
			File:   "File",
			Path:   "Path",
			Result: "Result",
			Remark: "Remark",
		}
	}
	m.PublishRowsReset()
}
