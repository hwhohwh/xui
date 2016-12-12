package main

import (
	"fmt"

	"github.com/ying32/ui"
	"github.com/ying32/xui"
)

type TMainWindowEvent struct {
	wx *xui.TXWindow
}

var mainWindow *xui.TXWindow

func loadMainWindowUI() {
	event := new(TMainWindowEvent)
	mainWindow, err := xui.NewFromFile(getCurrentDir()+"/main.xml", event)
	if err != nil {
		panic(err)
	}
	event.wx = mainWindow
	mainWindow.Window.OnClosing(mainWindowOnCloseing)
	mainWindow.Show()
}

func mainWindowOnCloseing(w *ui.Window) bool {
	ui.Quit()
	return true
}

func (wx *TMainWindowEvent) MenuClickAbout(sender *ui.MenuItem) {
	fmt.Println("click about")
	loadAboutWindow()
}
