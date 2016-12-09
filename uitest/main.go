// uitest project main.go
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ying32/ui"
	"github.com/ying32/xui"
)

type TXControlEvent struct {
	xw *xui.TXWindow
}

func (x *TXControlEvent) Test1(sender *ui.Button) {
	ui.MsgBox(x.xw.Window, "测试", "事件")
}

func (x *TXControlEvent) Test2(sender *ui.Button) {
	chk := x.xw.NameCheckbox("chk1")
	if chk != nil {
		chk.SetChecked(!chk.Checked())
	}
}

func (x *TXControlEvent) Testchk(sender *ui.Checkbox) {
	fmt.Println("checked=", sender.Checked())
}

func (x *TXControlEvent) TestChanged(sender *ui.Entry) {
	fmt.Println("OnChanged:", sender.Text())
}

func (x *TXControlEvent) TestSlider(sender *ui.Slider) {
	slider1 := x.xw.NameSlider("slider1")
	if slider1 != nil {
		slider1.SetValue(sender.Value())
	}
	progressbar1 := x.xw.NameProgressBar("progressbar1")
	if progressbar1 != nil {
		progressbar1.SetValue(sender.Value())
	}
	fmt.Println("slider.value=", sender.Value())
}

func (x *TXControlEvent) TestSpinbox(sender *ui.Spinbox) {
	fmt.Println("spibox.value=", sender.Value())
}

func (x *TXControlEvent) TestRadioSel(sender *ui.RadioButtons) {
	fmt.Println("radiobuttons selected=", sender.Selected())
}

func (x *TXControlEvent) TestMenu1Click(sender *ui.MenuItem) {
	fmt.Println("menuitem单击")
}

func (x *TXControlEvent) TestMemuOpen(sender *ui.MenuItem) {
	s := ui.OpenFile(x.xw.Window)
	if s != "" {
		if edit1 := x.xw.NameEntry("edit1"); edit1 != nil {
			edit1.SetText(s)
		}
	}
}

func (x *TXControlEvent) TestMenuSave(sender *ui.MenuItem) {
	ui.SaveFile(x.xw.Window)
}

func loadmyui() {

	event := &TXControlEvent{}
	if event == nil {
		panic("窗口事件创建失败!")
		return
	}

	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	w, err := xui.NewFromFile(path+"/ui.xml", event)
	//w, err := xui.NewFormBytes([]byte(uixmlstr), event)
	if err != nil {
		fmt.Println("错误：", err)
		return
	}
	event.xw = w
	w.Window.OnClosing(func(window *ui.Window) bool {
		ui.Quit()
		return true
	})
	w.Window.OnContentSizeChanged(func(window *ui.Window) {
		fmt.Println("OnContentSizeChanged")
	})
	w.Show()
}

func main() {
	xui.Application.Init()
	loadmyui()
	xui.Application.Run()
}
