package xui

import (
	"github.com/ying32/ui"
)

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

func (x *TXWindow) NameHorizontalSeparator(name string) *ui.Separator {
	c := x.NameControl(name)
	if c == nil {
		return nil
	}
	return c.(*ui.Separator)
}

func (x *TXWindow) NameVerticalSeparator(name string) *ui.Separator {
	c := x.NameControl(name)
	if c == nil {
		return nil
	}
	return c.(*ui.Separator)
}

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

func (x *TXWindow) NameMultilineEntry(name string) *ui.MultilineEntry {
	c := x.NameControl(name)
	if c == nil {
		return nil
	}
	return c.(*ui.MultilineEntry)
}
