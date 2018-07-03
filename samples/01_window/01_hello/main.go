// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/wrzfeijianshen/gui/gxui"
)

func appMain(driver gxui.Driver) {
}
func main() {
	gl.StartDriver(appMain)
}
