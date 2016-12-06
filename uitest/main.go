// uitest project main.go
package main

import (
	"fmt"

	"github.com/ying32/xui"

	"github.com/andlabs/ui"
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
	fmt.Println("check=", sender.Checked())
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

func loadmyui() {
	application := xui.NewApplication()
	event := &TXControlEvent{}
	if event == nil {
		panic("窗口事件创建失败!")
		return
	}
	w, err := xui.NewFromFile("ui.xml", event)
	if err != nil {
		fmt.Println("错误：", err)
		return
	}
	event.xw = w
	w.Window.OnClosing(func(window *ui.Window) bool {
		application.Quit()
		return true
	})
	w.Show()
	application.Run()

}

func main() {
	fmt.Println("Hello World!")
	loadmyui()

}
