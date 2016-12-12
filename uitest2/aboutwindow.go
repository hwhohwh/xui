package main

import (
	"github.com/ying32/ui"
	"github.com/ying32/xui"
)

type TAboutWindowEvent struct {
	wx *xui.TXWindow
}

func loadAboutWindow() {
	event := new(TAboutWindowEvent)
	aboutWindow, err := xui.NewFromFile(getCurrentDir()+"/about.xml", event)
	if err != nil {
		return
	}
	event.wx = aboutWindow
	aboutWindow.Window.OnClosing(aboutWindowOnClosing)
	aboutWindow.Show()
}

func aboutWindowOnClosing(w *ui.Window) bool {
	return true
}
