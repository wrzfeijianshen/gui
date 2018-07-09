// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"./nrtools"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/gxfont"
	"github.com/google/gxui/math"
	"github.com/google/gxui/samples/flags"
	"github.com/tarm/serial"
)

var (
	SerialPort        *serial.Port
	BoolSentSerialFor bool
	strSet1           string
	strSet2           string
	strSet3           string
)

func FuncSentSerialFor(str string) {
	BoolSentSerialFor = !BoolSentSerialFor
	str += "\r\n"
	for BoolSentSerialFor {
		nrtools.WriteSerial(SerialPort, str)
		time.Sleep(1 * time.Second)
	}
}

func FuncSentSerialForSet(str string) {
	BoolSentSerialFor = !BoolSentSerialFor
	strtmp := ""
	strSent := ""
	for BoolSentSerialFor {
		strSent = ""
		s := strings.Split(str, ",")
		//		fmt.Println(s)
		strSent = s[0]
		strSent += ","
		strSent += s[1]
		strSent += ","

		int0, _ := strconv.ParseFloat(s[2], 64)
		int1, _ := strconv.ParseFloat(s[4], 64)
		int2, _ := strconv.ParseFloat(strSet1, 64)
		int0 += int2
		int1 += int2
		strtmp = strconv.FormatFloat(int0, 'f', 4, 64)
		strtmp += ","
		strtmp += s[3]
		strtmp += ","
		strtmp += strconv.FormatFloat(int1, 'f', 4, 64)
		strtmp += ","
		strtmp += s[5]
		strtmp += ","
		strSent += strtmp

		int0, _ = strconv.ParseFloat(s[6], 64)
		int1, _ = strconv.ParseFloat(s[8], 64)
		int2, _ = strconv.ParseFloat(strSet2, 64)
		int0 += int2
		int1 += int2
		strtmp = strconv.FormatFloat(int0, 'f', 4, 64)
		strtmp += ","
		strtmp += s[7]
		strtmp += ","
		strtmp += strconv.FormatFloat(int1, 'f', 4, 64)
		strtmp += ","
		strtmp += s[9]
		strtmp += ","
		strSent += strtmp
		int0, _ = strconv.ParseFloat(s[10], 64)
		int1, _ = strconv.ParseFloat(s[12], 64)
		int2, _ = strconv.ParseFloat(strSet3, 64)
		int0 += int2
		int1 += int2
		strtmp = strconv.FormatFloat(int0, 'f', 4, 64)
		strtmp += ","
		strtmp += s[11]
		strtmp += ","
		strtmp += strconv.FormatFloat(int1, 'f', 4, 64)
		strtmp += ","
		strtmp += s[13]

		strSent += strtmp
		for i := 14; i < len(s); i++ {
			strSent += ","
			strSent += s[i]
		}
		//		fmt.Println(strSent)
		str = strSent
		nrtools.WriteSerial(SerialPort, strSent)
		time.Sleep(1 * time.Second)
	}
}

func addFrom( theme gxui.Theme){

}


func appMain(driver gxui.Driver) {
	BoolSentSerialFor = false
	theme := flags.CreateTheme(driver)
	layout := theme.CreateLinearLayout()
	layout.SetSizeMode(gxui.Fill)

	font, err := driver.CreateFont(gxfont.Default, 20)
	if err != nil {
		panic(err)
	}
	fontedit, err := driver.CreateFont(gxfont.Default, 15)
	if err != nil {
		panic(err)
	}
	window := theme.CreateWindow(320, 800, "NR串口工具")
	window.SetBackgroundBrush(gxui.CreateBrush(gxui.Black)) //背景色

	label := theme.CreateLabel()
	label.SetFont(font)
	label.SetText("Serial:")
	label.SetColor(gxui.White)
	layout.AddChild(label)

	//window.AddChild(label)
	edit := theme.CreateCodeEditor()
	edit.SetFont(fontedit)
	edit.SetText("COM10")
	edit.SetTabWidth(400)
	edit.SetDesiredWidth(600)
	//edit.SetSize(math.Size{100,300})
	//edit.SetMargin(math.Spacing{20,20,20,20})
	//edit.SetPadding(math.Spacing{10,20,200,100})
	layout.AddChild(edit)

	buttonState := map[gxui.Button]func() bool{}

	update := func() {
		for button, f := range buttonState {
			button.SetChecked(f())
		}
	}

	button := func(name string, action func(), isSelected func() bool) gxui.Button {

		b := theme.CreateButton()
		b.SetText(name)
		b.OnClick(func(gxui.MouseEvent) { action(); update() })
		layout.AddChild(b)
		buttonState[b] = isSelected
		return b
	}

	button("OpenSerial",
		func() {
			str := edit.Text()

			var flag bool
			SerialPort, flag = nrtools.OpenSerial(str)
			if flag != true {
				fmt.Printf("%s%s%s\n", "Open Serial ", str, " failure")
			} else {
				fmt.Printf("%s%s%s\n", "Open Serial ", str, " successful")
			}
		},
		func() bool { return true },
	)

	editSent := theme.CreateTextBox()
	editSent.SetFont(fontedit)
	editSent.SetText("$ECINS,103808,")
	editSent.SetDesiredWidth(300)
	layout.AddChild(editSent)
	labe2 := theme.CreateLabel()
	labe2.SetFont(font)
	labe2.SetText("GPS:")
	labe2.SetColor(gxui.White)
	layout.AddChild(labe2)

	editSent0 := theme.CreateTextBox()
	editSent0.SetFont(fontedit)
	editSent0.SetText("3000.5378,N,12204.7185,E,")
	editSent0.SetDesiredWidth(300)
	layout.AddChild(editSent0)

	labe3 := theme.CreateLabel()
	labe3.SetFont(font)
	labe3.SetText("BeiDou:")
	labe3.SetColor(gxui.White)
	layout.AddChild(labe3)

	editSent1 := theme.CreateTextBox()
	editSent1.SetFont(fontedit)
	editSent1.SetText("2959.1233,N,12206.2385,E,")
	editSent1.SetDesiredWidth(300)
	layout.AddChild(editSent1)

	editSent2 := theme.CreateTextBox()
	editSent2.SetFont(fontedit)
	editSent2.SetText("3052.1782,N,12308.5193,E,")
	editSent2.SetDesiredWidth(300)
	layout.AddChild(editSent2)

	editSent3 := theme.CreateTextBox()
	editSent3.SetFont(fontedit)
	editSent3.SetText("327.68,325.19,242.68,235.27,")
	editSent3.SetDesiredWidth(300)
	layout.AddChild(editSent3)

	editSent4 := theme.CreateTextBox()
	editSent4.SetFont(fontedit)
	editSent4.SetText("018.56,019.63,")
	editSent4.SetDesiredWidth(300)
	layout.AddChild(editSent4)

	editSent5 := theme.CreateTextBox()
	editSent5.SetFont(fontedit)
	editSent5.SetText("1712,128,r,1752,127,t,")
	editSent5.SetDesiredWidth(300)
	layout.AddChild(editSent5)

	editSent6 := theme.CreateTextBox()
	editSent6.SetFont(fontedit)
	editSent6.SetText("-021,T,026,H,10298,P,00325.3,140305,+08,")
	editSent6.SetDesiredWidth(300)
	layout.AddChild(editSent6)

	editSent7 := theme.CreateTextBox()
	editSent7.SetFont(fontedit)
	editSent7.SetText("1,1,1,1,1,1,1,1")
	editSent7.SetDesiredWidth(300)
	layout.AddChild(editSent7)

	//window.AddChild(edit)

	//	btnOpen := theme.CreateButton()
	//	btnOpen.SetText("OpenSerial:")
	//	layout.AddChild(btnOpen)

	button("SentSerial",
		func() {
			str := editSent.Text()
			str += editSent0.Text()
			str += editSent1.Text()
			str += editSent2.Text()
			str += editSent3.Text()
			str += editSent4.Text()
			str += editSent5.Text()
			str += editSent6.Text()
			str += editSent7.Text()
			nrtools.WriteSerial(SerialPort, str)
		},
		func() bool { return true },
	)

	button("SentSerialFor",
		func() {
			str := editSent.Text()
			str += editSent0.Text()
			str += editSent1.Text()
			str += editSent2.Text()
			str += editSent3.Text()
			str += editSent4.Text()
			str += editSent5.Text()
			str += editSent6.Text()
			str += editSent7.Text()
			go FuncSentSerialFor(str)
		},
		func() bool { return true },
	)
	editSet1 := theme.CreateTextBox()
	editSet1.SetFont(fontedit)
	editSet1.SetText("0.0001")
	editSet1.SetDesiredWidth(300)
	layout.AddChild(editSet1)
	editSet2 := theme.CreateTextBox()
	editSet2.SetFont(fontedit)
	editSet2.SetText("0.0010")
	editSet2.SetDesiredWidth(300)
	layout.AddChild(editSet2)
	editSet3 := theme.CreateTextBox()
	editSet3.SetFont(fontedit)
	editSet3.SetText("0.0100")
	editSet3.SetDesiredWidth(300)
	layout.AddChild(editSet3)
	button("SentSerialForADD",
		func() {
			strSet1 = editSet1.Text()
			strSet2 = editSet2.Text()
			strSet3 = editSet3.Text()

			str := editSent.Text()
			str += editSent0.Text()
			str += editSent1.Text()
			str += editSent2.Text()
			str += editSent3.Text()
			str += editSent4.Text()
			str += editSent5.Text()
			str += editSent6.Text()
			str += editSent7.Text()

			go FuncSentSerialForSet(str)
		},
		func() bool { return true },
	)

	update()
	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(layout)
	window.OnClose(driver.Terminate)
	window.SetPadding(math.Spacing{L: 10, T: 10, R: 10, B: 10})
}

func main() {
	gl.StartDriver(appMain)
}
