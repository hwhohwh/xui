package xui

import (
	"bytes"
	"io/ioutil"
	"reflect"

	xmldom "bitbucket.org/rj/xmldom-go"
	"github.com/andlabs/ui"
)

type TXWindow struct {
	Window       *ui.Window
	nameControls map[string]interface{}
}

func NewFromFile(xmlfile string, event interface{}) (*TXWindow, error) {
	bs, err := ioutil.ReadFile(xmlfile)
	if err != nil {
		return nil, err
	}
	return NewFormBytes(bs, event)
}

func NewFormBytes(xmlstr []byte, event interface{}) (*TXWindow, error) {
	doc, err := xmldom.ParseXml(bytes.NewReader(xmlstr))
	if err != nil {
		return nil, err
	}
	w := new(TXWindow)
	w.nameControls = make(map[string]interface{}, 0)
	root := doc.DocumentElement()
	if root != nil && root.NodeName() == "Window" {
		w.Window = w.buildWindow(root)
		if w.Window != nil {
			ctlroot := w.buildControls(root, nil, event)
			if ctlroot != nil {
				w.Window.SetChild(ctlroot)
			} else {
				panic("根控件不能为空！")
			}
		} else {
			panic("Window创建失败!")
		}
	} else {
		panic("xml不符合要求！")
	}
	return w, nil
}

func (x *TXWindow) buildWindow(node xmldom.Node) *ui.Window {
	attrs := newXmlAttrsMap(node)
	if attrs == nil {
		return nil
	}
	w := ui.NewWindow(attrs.Title(), attrs.Width(), attrs.Height(), attrs.HasMenu(), attrs.Center())
	if w != nil {
		w.SetMargined(attrs.Margined())
	}
	return w
}

func (x *TXWindow) appendControl(parent interface{}, child ui.Control, attrs *TXmlAttrs) {
	if parent == nil {
		return
	}
	switch getClassName(parent) {
	case "Box":
		box := parent.(*ui.Box)
		box.Append(child, attrs.Stretchy())

	case "Tab":
		tab := parent.(*ui.Tab)
		tab.Append(attrs.Text(), child)

	case "Group":
		group := parent.(*ui.Group)
		group.SetChild(child)

	case "Combobox":
		combox := parent.(*ui.Combobox)
		combox.Append(attrs.Text())
	}
}

func (x *TXWindow) buildControls(node xmldom.Node, parent ui.Control, event interface{}) ui.Control {
	if !node.HasChildNodes() {
		return nil
	}
	var i uint
	var root, pcontrol ui.Control
	var attrs *TXmlAttrs
	lParent := parent
	root = nil

	setCommAttr := func() {
		if pcontrol != nil {
			if attrs.Enabled() {
				pcontrol.Enable()
			} else {
				pcontrol.Disable()
			}
			if attrs.Visible() {
				pcontrol.Show()
			} else {
				pcontrol.Hide()
			}
		}
	}

	for i = 0; i < node.ChildNodes().Length(); i++ {
		subnode := node.ChildNodes().Item(i)
		if subnode.NodeType() != 1 {
			continue
		}
		pcontrol = nil
		attrs = newXmlAttrsMap(subnode)
		switch subnode.NodeName() {
		case "Button":
			btn := ui.NewButton(attrs.Text())
			pcontrol = btn
			eventname := attrs.Onclick()
			if eventname != "" {
				m, ok := reflect.TypeOf(event).MethodByName(eventname)
				if ok {
					btn.OnClicked(func(sender *ui.Button) {
						m.Func.Call([]reflect.Value{reflect.ValueOf(event), reflect.ValueOf(sender)})
					})
				}
			}
			x.appendControl(parent, btn, attrs)
			x.addNameControl(attrs.Name(), btn)

		case "Entry":
			entry := ui.NewEntry()
			entry.SetReadOnly(attrs.ReadOnly())
			eventname := attrs.OnChanged()
			if eventname != "" {
				m, ok := reflect.TypeOf(event).MethodByName(eventname)
				if ok {
					entry.OnChanged(func(sender *ui.Entry) {
						m.Func.Call([]reflect.Value{reflect.ValueOf(event), reflect.ValueOf(sender)})
					})
				}
			}
			pcontrol = entry
			entry.SetText(attrs.Text())
			x.appendControl(parent, pcontrol, attrs)
			x.addNameControl(attrs.Name(), entry)

		case "HorizontalBox":
			box := ui.NewHorizontalBox()
			box.SetPadded(attrs.Padded())
			pcontrol = box
			x.appendControl(parent, box, attrs)
			lParent = pcontrol
			x.addNameControl(attrs.Name(), box)

		case "VerticalBox":
			box := ui.NewVerticalBox()
			box.SetPadded(attrs.Padded())
			pcontrol = box
			x.appendControl(parent, box, attrs)
			lParent = pcontrol
			x.addNameControl(attrs.Name(), box)

		case "Label":
			lbl := ui.NewLabel(attrs.Text())
			pcontrol = lbl
			x.appendControl(parent, lbl, attrs)
			x.addNameControl(attrs.Name(), lbl)

		case "Checkbox":
			chk := ui.NewCheckbox(attrs.Text())
			chk.SetChecked(attrs.Checked())
			eventname := attrs.OnToggled()
			if eventname != "" {
				m, ok := reflect.TypeOf(event).MethodByName(eventname)
				if ok {
					chk.OnToggled(func(sender *ui.Checkbox) {
						m.Func.Call([]reflect.Value{reflect.ValueOf(event), reflect.ValueOf(sender)})
					})
				}
			}
			pcontrol = chk
			x.appendControl(parent, pcontrol, attrs)
			x.addNameControl(attrs.Name(), chk)

		case "Tab":
			tab := ui.NewTab()
			//tab.SetMargined(attrs.Margined())
			pcontrol = tab
			x.appendControl(parent, tab, attrs)
			lParent = pcontrol
			x.addNameControl(attrs.Name(), pcontrol)

		case "Group":
			group := ui.NewGroup(attrs.Text())
			pcontrol = group
			group.SetMargined(attrs.Margined())
			x.appendControl(parent, group, attrs)
			lParent = parent
			x.addNameControl(attrs.Name(), group)

		case "Combobox":
			combox := ui.NewCombobox()

			eventname := attrs.OnSelected()
			if eventname != "" {
				m, ok := reflect.TypeOf(event).MethodByName(eventname)
				if ok {
					combox.OnSelected(func(sender *ui.Combobox) {
						m.Func.Call([]reflect.Value{reflect.ValueOf(event), reflect.ValueOf(sender)})
					})
				}
			}
			pcontrol = combox
			x.appendControl(parent, combox, attrs)
			x.addNameControl(attrs.Name(), combox)
			setCommAttr()
			x.buildControls(subnode, combox, event)
			combox.SetSelected(attrs.Selected())
			continue

		case "CombItem":
			//x.appendControl(parent, nil, attrs)
			if parent != nil {
				parent.(*ui.Combobox).Append(attrs.Text())
			}

		case "DatePicker":
			pcontrol := ui.NewDatePicker()
			x.appendControl(parent, pcontrol, attrs)
			x.addNameControl(attrs.Name(), pcontrol)

		case "DateTimePicker":
			pcontrol = ui.NewDateTimePicker()
			x.appendControl(parent, pcontrol, attrs)
			x.addNameControl(attrs.Name(), pcontrol)

		case "TimePicker":
			pcontrol = ui.NewTimePicker()
			x.appendControl(parent, pcontrol, attrs)
			x.addNameControl(attrs.Name(), pcontrol)

		case "ProgressBar":
			probar := ui.NewProgressBar()
			probar.SetValue(attrs.IntValue())
			pcontrol = probar
			x.appendControl(parent, probar, attrs)
			x.addNameControl(attrs.Name(), probar)

		case "RadioButtons":
			pcontrol := ui.NewRadioButtons()
			x.appendControl(parent, pcontrol, attrs)
			x.addNameControl(attrs.Name(), pcontrol)

		case "HorizontalSeparator":
			pcontrol = ui.NewHorizontalSeparator()
			x.appendControl(parent, pcontrol, attrs)
			x.addNameControl(attrs.Name(), pcontrol)

		case "Slider":
			slider := ui.NewSlider(attrs.Min(), attrs.Max())
			slider.SetValue(attrs.IntValue())
			eventname := attrs.OnChanged()
			if eventname != "" {
				m, ok := reflect.TypeOf(event).MethodByName(eventname)
				if ok {
					slider.OnChanged(func(sender *ui.Slider) {
						m.Func.Call([]reflect.Value{reflect.ValueOf(event), reflect.ValueOf(sender)})
					})
				}
			}
			pcontrol = slider
			x.appendControl(parent, slider, attrs)
			x.addNameControl(attrs.Name(), pcontrol)

		case "Spinbox":
			spinbox := ui.NewSpinbox(attrs.Min(), attrs.Max())
			spinbox.SetValue(attrs.IntValue())
			eventname := attrs.OnChanged()
			if eventname != "" {
				m, ok := reflect.TypeOf(event).MethodByName(eventname)
				if ok {
					spinbox.OnChanged(func(sender *ui.Spinbox) {
						m.Func.Call([]reflect.Value{reflect.ValueOf(event), reflect.ValueOf(sender)})
					})
				}
			}
			pcontrol = spinbox
			x.appendControl(parent, spinbox, attrs)
			x.addNameControl(attrs.Name(), spinbox)

		default:
		}
		setCommAttr()
		if root == nil {
			root = lParent
		}
		if lParent != nil {
			x.buildControls(subnode, lParent, event)
		}
	}
	return root
}

func (x *TXWindow) addNameControl(name string, control interface{}) {
	if name == "" || control == "" {
		return
	}
	if _, ok := x.nameControls[name]; !ok {
		x.nameControls[name] = control
	}
}

func (x *TXWindow) NameControl(name string) interface{} {
	if v, ok := x.nameControls[name]; ok {
		return v
	}
	return nil
}

func (x *TXWindow) NameButton(name string) *ui.Button {
	c := x.NameControl(name)
	if c == nil {
		return nil
	}
	return c.(*ui.Button)
}

func (x *TXWindow) NameCheckbox(name string) *ui.Checkbox {
	c := x.NameControl(name)
	if c == nil {
		return nil
	}
	return c.(*ui.Checkbox)
}

func (x *TXWindow) NameEntry(name string) *ui.Entry {
	c := x.NameControl(name)
	if c == nil {
		return nil
	}
	return c.(*ui.Entry)
}

func (x *TXWindow) NameHorizontalBox(name string) *ui.Box {
	c := x.NameControl(name)
	if c == nil {
		return nil
	}
	return c.(*ui.Box)
}

func (x *TXWindow) NameVerticalBox(name string) *ui.Box {
	c := x.NameControl(name)
	if c == nil {
		return nil
	}
	return c.(*ui.Box)
}

func (x *TXWindow) NameLabel(name string) *ui.Label {
	c := x.NameControl(name)
	if c == nil {
		return nil
	}
	return c.(*ui.Label)
}

func (x *TXWindow) NameTab(name string) *ui.Tab {
	c := x.NameControl(name)
	if c == nil {
		return nil
	}
	return c.(*ui.Tab)
}

func (x *TXWindow) NameGroup(name string) *ui.Group {
	c := x.NameControl(name)
	if c == nil {
		return nil
	}
	return c.(*ui.Group)
}

func (x *TXWindow) NameCombobox(name string) *ui.Combobox {
	c := x.NameControl(name)
	if c == nil {
		return nil
	}
	return c.(*ui.Combobox)
}

func (x *TXWindow) NameDatePicker(name string) *ui.DateTimePicker {
	c := x.NameControl(name)
	if c == nil {
		return nil
	}
	return c.(*ui.DateTimePicker)
}

func (x *TXWindow) NameDateTimePicker(name string) *ui.DateTimePicker {
	c := x.NameControl(name)
	if c == nil {
		return nil
	}
	return c.(*ui.DateTimePicker)
}

func (x *TXWindow) NameTimePicker(name string) *ui.DateTimePicker {
	c := x.NameControl(name)
	if c == nil {
		return nil
	}
	return c.(*ui.DateTimePicker)
}

func (x *TXWindow) NameProgressBar(name string) *ui.ProgressBar {
	c := x.NameControl(name)
	if c == nil {
		return nil
	}
	return c.(*ui.ProgressBar)
}

func (x *TXWindow) NameRadioButtons(name string) *ui.RadioButtons {
	c := x.NameControl(name)
	if c == nil {
		return nil
	}
	return c.(*ui.RadioButtons)
}

//func (x *TXWindow) NameHorizontalSeparator(name string) *ui.h {
//	c := x.NameControl(name)
//	if c == nil {
//		return nil
//	}
//	return c.(*ui.HorizontalSeparator)
//}

func (x *TXWindow) NameSlider(name string) *ui.Slider {
	c := x.NameControl(name)
	if c == nil {
		return nil
	}
	return c.(*ui.Slider)
}

func (x *TXWindow) NameSpinbox(name string) *ui.Spinbox {
	c := x.NameControl(name)
	if c == nil {
		return nil
	}
	return c.(*ui.Spinbox)
}

func (x *TXWindow) Show() {
	x.Window.Show()
}
