package main

import (
	. "github.com/lxn/walk/declarative"
)

func main() {

	mainWindow.Run()

}

var LableHello = Label{
	Text: "Hello girl......",
}

var widget = []Widget{
	LableHello,
}

var mainWindow = MainWindow{

	Title:    "Walk",
	MinSize:  Size{400, 400},
	Layout:   VBox{},
	Children: widget,
	ToolBar: ToolBar{
		ButtonStyle: ToolBarButtonImageBeforeText,
		Items: []MenuItem{
			Menu{
				Text:  "New A",
				Image: "img/document-new.png",
				Items: []MenuItem{
					Action{
						Text: "A",
					},
					Action{
						Text: "B",
					},
					Action{
						Text: "C",
					},
				},
			},
			Action{
				Text:    "Special",
				Image:   "img/system-shutdown.png",
				Enabled: Bind("isSpecialMode && enabledCB.Checked"),
			},
		},
	},
}
