package xui

import (
	"reflect"
	"strconv"
	"strings"
)

/// 从实列中获取类名，不含  *.ui
func getClassName(ctl interface{}) string {
	return strings.Replace(reflect.TypeOf(ctl).String(), "*ui.", "", -1)
}

// str to int
func atoi(s string) (r int) {
	r, _ = strconv.Atoi(s)
	return
}

// str to bool
func atob(s string) (r bool) {
	r, _ = strconv.ParseBool(s)
	return
}
