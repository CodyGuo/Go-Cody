package main

import (
	"fmt"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type Foo struct {
	Index int
	Data  string
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
					{Name: "Index"}, // Changed Title to Name.
					{Name: "Data"},
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
			Index: i,
			Data:  "aa",
		}
	}
	m.PublishRowsReset()
}
