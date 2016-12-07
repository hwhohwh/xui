package xui

import "github.com/ying32/ui"

type TApplication struct {
}

func NewApplication() *TApplication {
	return &TApplication{}
}

func (self *TApplication) Run() {
	ui.OnShouldQuit(func() bool {
		return true
	})
	ui.Main(func() {

	})
}

func (self *TApplication) Quit() {
	ui.Quit()
}
