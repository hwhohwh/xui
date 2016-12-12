package xui

import (
	"bytes"
	//	"fmt"
	"io/ioutil"
	"reflect"

	xmldom "bitbucket.org/rj/xmldom-go"
	"github.com/ying32/ui"
)

const (
	methodTypeClicked = iota + 0
	methodTypeChanged
	methodTypeSelected
	methodTypeToggled
)

type TXWindow struct {
	Window       *ui.Window
	event        interface{}
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
	w.event = event
	w.nameControls = make(map[string]interface{}, 0)
	root := doc.DocumentElement()
	if root != nil && root.NodeName() == "Window" {
		w.Window = w.buildWindow(root)
		if w.Window != nil {
			ctlroot := w.buildControls(root, nil)
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
	if attrs.HasMenu() {
		var i uint
		for i = 0; i < node.ChildNodes().Length(); i++ {
			menuNode := node.ChildNodes().Item(i)
			if menuNode.NodeName() == "Menus" {
				if menuNode.ChildNodes().Length() > 0 {
					x.buildMenus(menuNode, nil, nil)
				}
				node.RemoveChild(menuNode)
				break
			}
		}
	}
	w := ui.NewWindow(attrs.Title(), attrs.Left(), attrs.Top(), attrs.Width(), attrs.Height(), attrs.HasMenu())
	if w != nil {
		w.SetMargined(attrs.Margined())
		w.SetBorderless(attrs.Borderless())
		w.SetFullscreen(attrs.Fullscreen())
		if attrs.Center() {
			w.Center()
		}
	}

	return w
}

func (x *TXWindow) buildMenus(node xmldom.Node, w *ui.Window, menu *ui.Menu) {
	var i uint
	for i = 0; i < node.ChildNodes().Length(); i++ {
		menuNode := node.ChildNodes().Item(i)

		if menuNode.NodeType() != 1 {
			continue
		}
		attrs := newXmlAttrsMap(menuNode)
		switch menuNode.NodeName() {
		case "MenuAbout":
			if menu != nil {
				subm := menu.AppendAbout()
				m, ok := x.getMethod(attrs.Onclick())
				if ok {
					subm.OnClicked(func(sender *ui.MenuItem) {
						m.Func.Call([]reflect.Value{reflect.ValueOf(x.event), reflect.ValueOf(sender)})
					})
				}

				x.addNameControl(attrs.Name(), subm)
			}

		case "Menu":
			menu = ui.NewMenu(attrs.Text())

		case "MenuItem":
			if menu != nil {
				if attrs.Text() == "-" {
					menu.AppendSeparator()
				} else {
					subm := menu.Append(attrs.Text())
					//subm.SetChecked(attrs.Checked())
					if attrs.Enabled() {
						subm.Enable()
					} else {
						subm.Disable()
					}
					m, ok := x.getMethod(attrs.Onclick())
					if ok {
						subm.OnClicked(func(sender *ui.MenuItem) {
							m.Func.Call([]reflect.Value{reflect.ValueOf(x.event), reflect.ValueOf(sender)})
						})
					}
					x.addNameControl(attrs.Name(), subm)
				}
			}

		case "MenuQuit":
			if menu != nil {
				menu.AppendQuit()
			}

		case "MenuPreferences":
			if menu != nil {
				subm := menu.AppendPreferences()
				m, ok := x.getMethod(attrs.Onclick())
				if ok {
					subm.OnClicked(func(sender *ui.MenuItem) {
						m.Func.Call([]reflect.Value{reflect.ValueOf(x.event), reflect.ValueOf(sender)})
					})
				}
				x.addNameControl(attrs.Name(), subm)
			}

		case "MenuCheck":

			if menu != nil {
				subm := menu.AppendCheck(attrs.Text())
				subm.SetChecked(attrs.Checked())
				if attrs.Enabled() {
					subm.Enable()
				} else {
					subm.Disable()
				}
				m, ok := x.getMethod(attrs.Onclick())
				if ok {
					subm.OnClicked(func(sender *ui.MenuItem) {
						m.Func.Call([]reflect.Value{reflect.ValueOf(x.event), reflect.ValueOf(sender)})
					})
				}
				x.addNameControl(attrs.Name(), subm)
			}

		default:
			continue
		}
		if menuNode.HasChildNodes() {
			x.buildMenus(menuNode, w, menu)
		}
	}
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

func (x *TXWindow) getMethod(name string) (reflect.Method, bool) {
	if name == "" {
		var m reflect.Method
		return m, false
	}
	return reflect.TypeOf(x.event).MethodByName(name)
}

func (x *TXWindow) buildControls(node xmldom.Node, parent ui.Control) ui.Control {
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

			m, ok := x.getMethod(attrs.Onclick())
			if ok {
				btn.OnClicked(func(sender *ui.Button) {
					m.Func.Call([]reflect.Value{reflect.ValueOf(x.event), reflect.ValueOf(sender)})
				})
			}

			x.appendControl(parent, btn, attrs)
			x.addNameControl(attrs.Name(), btn)

		case "Entry":
			entry := ui.NewEntry()
			entry.SetReadOnly(attrs.ReadOnly())

			m, ok := x.getMethod(attrs.OnChanged())
			if ok {
				entry.OnChanged(func(sender *ui.Entry) {
					m.Func.Call([]reflect.Value{reflect.ValueOf(x.event), reflect.ValueOf(sender)})
				})
			}
			pcontrol = entry
			entry.SetText(attrs.Text())
			x.appendControl(parent, pcontrol, attrs)
			x.addNameControl(attrs.Name(), entry)

		case "MultilineEntry":
			mentry := ui.NewMultilineEntry(attrs.NonWrapping())
			mentry.SetReadOnly(attrs.ReadOnly())

			m, ok := x.getMethod(attrs.OnChanged())
			if ok {
				mentry.OnChanged(func(sender *ui.MultilineEntry) {
					m.Func.Call([]reflect.Value{reflect.ValueOf(x.event), reflect.ValueOf(sender)})
				})
			}
			pcontrol = mentry
			mentry.SetText(attrs.Text())
			x.appendControl(parent, pcontrol, attrs)
			x.addNameControl(attrs.Name(), mentry)

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

			m, ok := x.getMethod(attrs.OnToggled())
			if ok {
				chk.OnToggled(func(sender *ui.Checkbox) {
					m.Func.Call([]reflect.Value{reflect.ValueOf(x.event), reflect.ValueOf(sender)})
				})
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
			lParent = pcontrol
			x.addNameControl(attrs.Name(), group)

		case "Combobox":
			combox := ui.NewCombobox()

			m, ok := x.getMethod(attrs.OnSelected())
			if ok {
				combox.OnSelected(func(sender *ui.Combobox) {
					m.Func.Call([]reflect.Value{reflect.ValueOf(x.event), reflect.ValueOf(sender)})
				})
			}

			pcontrol = combox
			x.appendControl(parent, combox, attrs)
			x.addNameControl(attrs.Name(), combox)
			setCommAttr()
			x.buildControls(subnode, combox)
			combox.SetSelected(attrs.Selected())
			continue

		case "EditableCombobox":

			combox := ui.NewEditableCombobox()
			combox.SetText(attrs.Text())
			m, ok := x.getMethod(attrs.OnChanged())
			if ok {
				combox.OnChanged(func(sender *ui.EditableCombobox) {
					m.Func.Call([]reflect.Value{reflect.ValueOf(x.event), reflect.ValueOf(sender)})
				})
			}
			pcontrol = combox
			x.appendControl(parent, combox, attrs)
			x.addNameControl(attrs.Name(), combox)
			setCommAttr()
			x.buildControls(subnode, combox)
			continue

		case "TextItem":

			if parent != nil {
				switch getClassName(parent) {
				case "Combobox":
					parent.(*ui.Combobox).Append(attrs.Text())

				case "EditableCombobox":
					parent.(*ui.EditableCombobox).Append(attrs.Text())

				case "RadioButtons":
					parent.(*ui.RadioButtons).Append(attrs.Text())
				}

			}
			continue

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
			radio := ui.NewRadioButtons()

			if m, ok := x.getMethod(attrs.OnSelected()); ok {
				radio.OnSelected(func(sender *ui.RadioButtons) {
					m.Func.Call([]reflect.Value{reflect.ValueOf(x.event), reflect.ValueOf(sender)})
				})
			}
			pcontrol = radio
			x.appendControl(parent, pcontrol, attrs)
			x.addNameControl(attrs.Name(), pcontrol)

			//setCommAttr()
			x.buildControls(subnode, radio)
			radio.SetSelected(attrs.Selected())
			continue

		case "HorizontalSeparator":
			pcontrol = ui.NewHorizontalSeparator()
			x.appendControl(parent, pcontrol, attrs)
			x.addNameControl(attrs.Name(), pcontrol)

		case "VerticalSeparator":
			pcontrol = ui.NewVerticalSeparator()
			x.appendControl(parent, pcontrol, attrs)
			x.addNameControl(attrs.Name(), pcontrol)

		case "Slider":
			slider := ui.NewSlider(attrs.Min(), attrs.Max())
			slider.SetValue(attrs.IntValue())

			m, ok := x.getMethod(attrs.OnChanged())
			if ok {
				slider.OnChanged(func(sender *ui.Slider) {
					m.Func.Call([]reflect.Value{reflect.ValueOf(x.event), reflect.ValueOf(sender)})
				})
			}
			pcontrol = slider
			x.appendControl(parent, slider, attrs)
			x.addNameControl(attrs.Name(), pcontrol)

		case "Spinbox":
			spinbox := ui.NewSpinbox(attrs.Min(), attrs.Max())
			spinbox.SetValue(attrs.IntValue())

			m, ok := x.getMethod(attrs.OnChanged())
			if ok {
				spinbox.OnChanged(func(sender *ui.Spinbox) {
					m.Func.Call([]reflect.Value{reflect.ValueOf(x.event), reflect.ValueOf(sender)})
				})
			}

			pcontrol = spinbox
			x.appendControl(parent, spinbox, attrs)
			x.addNameControl(attrs.Name(), spinbox)

		default:
			continue
		}
		setCommAttr()
		if root == nil {
			root = lParent
		}
		if lParent != nil {
			x.buildControls(subnode, lParent)
		}
	}
	return root
}

func (x *TXWindow) Show() {
	x.Window.Show()
}
