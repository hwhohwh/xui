package xui

import (
	xmldom "bitbucket.org/rj/xmldom-go"
)

type TXmlAttrs struct {
	attrs map[string]string
}

func newXmlAttrsMap(node xmldom.Node) (res *TXmlAttrs) {
	if node.NodeType() == 3 {
		return nil
	}
	res = new(TXmlAttrs)
	res.attrs = make(map[string]string, 0)
	var i uint
	for i = 0; i < node.Attributes().Length(); i++ {
		attr := node.Attributes().Item(i)
		if _, ok := res.attrs[attr.NodeName()]; !ok {
			res.attrs[attr.NodeName()] = attr.NodeValue()
		}
	}
	return
}

func (x *TXmlAttrs) HasAttr(name string) bool {
	if _, ok := x.attrs[name]; ok {
		return true
	}
	return false
}

func (x *TXmlAttrs) Get(name string) string {
	return x.GetDef(name, "")
}

func (x *TXmlAttrs) GetDef(name, def string) string {
	if v, ok := x.attrs[name]; ok {
		return v
	}
	return def
}

func (x *TXmlAttrs) ToInt(name string) int {
	return atoi(x.Get(name))
}

func (x *TXmlAttrs) ToIntDef(name string, def int) int {
	v := x.GetDef(name, "")
	if v == "" {
		return def
	}
	return atoi(v)
}

func (x *TXmlAttrs) ToBool(name string) bool {
	return atob(x.Get(name))
}

func (x *TXmlAttrs) ToBoolDef(name string, def bool) bool {
	v := x.GetDef(name, "")
	if v == "" {
		return def
	}
	return atob(v)
}

func (x *TXmlAttrs) Title() string {
	return x.Get("title")
}

func (x *TXmlAttrs) Text() string {
	return x.Get("text")
}

func (x *TXmlAttrs) Name() string {
	return x.Get("name")
}

func (x *TXmlAttrs) Onclick() string {
	return x.Get("onclick")
}

func (x *TXmlAttrs) Top() int {
	return x.ToInt("top")
}

func (x *TXmlAttrs) Left() int {
	return x.ToInt("left")
}

func (x *TXmlAttrs) Width() int {
	return x.ToInt("width")
}

func (x *TXmlAttrs) Height() int {
	return x.ToInt("height")
}

func (x *TXmlAttrs) Center() bool {
	return x.ToBool("center")
}

func (x *TXmlAttrs) HasMenu() bool {
	return x.ToBool("hasmenu")
}

func (x *TXmlAttrs) Margined() bool {
	return x.ToBoolDef("margined", true)
}

func (x *TXmlAttrs) Enabled() bool {
	return x.ToBoolDef("enabled", true)
}

func (x *TXmlAttrs) Visible() bool {
	return x.ToBoolDef("visible", true)
}

func (x *TXmlAttrs) Checked() bool {
	return x.ToBool("checked")
}

func (x *TXmlAttrs) OnToggled() string {
	return x.Get("ontoggled")
}

func (x *TXmlAttrs) OnChanged() string {
	return x.Get("onchanged")
}

func (x *TXmlAttrs) ReadOnly() bool {
	return x.ToBoolDef("readonly", false)
}

func (x *TXmlAttrs) Padded() bool {
	return x.ToBool("padded")
}

func (x *TXmlAttrs) Stretchy() bool {
	return x.ToBoolDef("stretchy", true)
}

func (x *TXmlAttrs) Selected() int {
	return x.ToIntDef("selected", -1)
}

func (x *TXmlAttrs) OnSelected() string {
	return x.Get("onselected")
}

func (x *TXmlAttrs) IntValue() int {
	return x.ToInt("value")
}

func (x *TXmlAttrs) Min() int {
	return x.ToInt("min")
}

func (x *TXmlAttrs) Max() int {
	return x.ToIntDef("max", 100)
}
