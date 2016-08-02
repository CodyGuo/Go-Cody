package main

import (
	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func() {
		name := ui.NewEntry()
		button := ui.NewButton("Greet")
		greeting := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Enter your name:"), true)
		box.Append(name, true)
		box.Append(button, true)
		box.Append(greeting, true)
		window := ui.NewWindow("Hello", 200, 100, true)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			greeting.SetText("Hello, " + name.Text() + "!")
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
