// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/wrzfeijianshen/gui/gxui/mixins"

	"github.com/wrzfeijianshen/gui/gxui"
	"github.com/wrzfeijianshen/gui/drivers/gl"
	"github.com/wrzfeijianshen/gui/gxui/math"
)
type Window struct {
	mixins.Window
}
func appMain(driver gxui.Driver) {
	w := &Window{}
	//
	w.Window.Init(w, driver, 100, 200, "hello")
	w.SetSize(math.Size{200,200})
}
func main() {
	gl.StartDriver(appMain)
}
