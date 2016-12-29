package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	//create a window
	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Hello World Example")
	window.SetMinimumSize2(200, 200)

	//create a layout
	var layout = widgets.NewQVBoxLayout()

	//add the layout to the centralWidget
	var centralWidget = widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(layout)

	//create a button and connect the clicked signal
	var button = widgets.NewQPushButton2("Click me!", nil)
	button.ConnectClicked(func(flag bool) {
		widgets.QMessageBox_Information(nil, "OK", "You clicked me!", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})

	//add the button to the layout
	layout.AddWidget(button, 0, core.Qt__AlignCenter)

	//add the centralWidget to the window and show the window
	window.SetCentralWidget(centralWidget)
	window.Show()

	widgets.QApplication_Exec()
}
