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

func (self *TApplication) Init() {
	err := ui.InitApp()
	if err != nil {
		panic(err)
	}
}

func (self *TApplication) Quit() {
	ui.Quit()
}
