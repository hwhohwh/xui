# xui
xui是一个基于[andlabs的ui库](https://github.com/andlabs/ui)的，使用xml布局来替代手动创建，目前仅打算支持Mac OS X,
目前ui库已经转为由[本人自己维护的版本了](https://github.com/ying32/ui)。   

## 使用例程


#### xml布局文件
```xml  

<?xml encoding="utf-8" version="1.0" ?>
<Window width="600" height="700" center="true" hasmenu="true" title="这是一个测试" margined="true">
	<Menus>
	     <Menu text="文件">
		     <MenuAbout />
			 <MenuQuit />
			 <MenuPreferences />
		     <MenuItem text="新建(F)" onclick="TestMenu1Click" />
			 <MenuItem text="-" />
			 <MenuItem text="打开(O)..." onclick="TestMemuOpen" />
			 <MenuItem text="保存(S)..." onclick="TestMenuSave"  checked="true"/>
			 <MenuCheck text="测试一个选项" checked="true"  enabled="false"/>
		 </Menu>
	</Menus>
	<HorizontalBox>
		<Tab>
		     <VerticalBox text="TAB1">
			     <Label text="我是标签" />
				 <Entry name="edit1" text="默认文本" onchanged="TestChanged"/>
		         <Button name="test" text="按钮3" onclick="Test1" />
				 <Combobox name="comb1" onselected="TestSelected" selected="0">
				     <CombItem text="Item1"/>
					 <CombItem text="Item2"/>
					 <CombItem text="Item3"/>
				 </Combobox>
				 <Checkbox name="chk2" text="选项2" checked="true" ontoggled="Testchk" />
				 <DatePicker />
				 <DateTimePicker />
				 <TimePicker />
				 <ProgressBar name="progressbar1" />
				 <HorizontalSeparator />
				 <Slider value="30" name="slider1" onchanged="TestSlider" />
				 <Spinbox value="50" name="spinbox1"  onchanged="TestSpinbox" />
		     </VerticalBox>
		     <VerticalBox text="TAB2">
		         <Group text="选项组">
		             <Button name="test" text="按钮4" onclick="Test2"/>
			         <Checkbox name="chk1" text="选项1" checked="true" ontoggled="Testchk" />
			     </Group>
				 <RadioButtons name="radio1" selected="0" onselected="TestRadioSel">
				    <RadioItem text="选项1" />
					<RadioItem text="选项2" />
				 </RadioButtons>
				 <MultilineEntry text="这是一个测试"></MultilineEntry>
		     </VerticalBox>
		</Tab>
	</HorizontalBox>
</Window>

```   

#### Go代码  

```go  

// uitest project main.go
package main

import (
	"fmt"

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
	application := xui.NewApplication()
	application.Init()

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
	w.Window.OnContentSizeChanged(func(window *ui.Window) {
		fmt.Println("OnContentSizeChanged")
	})
	w.Show()
	application.Run()

}

func main() {
	fmt.Println("Hello World!")
	loadmyui()

}


```
 
