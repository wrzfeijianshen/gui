package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn_diy/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)
	builder := gtk.NewBuilder()
	builder.AddFromFile("btn.glade")
	window := gtk.WindowFromObject(builder.GetObject("window1"))
	button1 := gtk.ButtonFromObject(builder.GetObject("button1"))
	button2 := gtk.ButtonFromObject(builder.GetObject("button2"))

	button1.SetLabel("图片") // 设置文本信息
	fmt.Println("button1 txt : ", button1.GetLabel())
	button1.SetSensitive(false) // 变灰色不可按
	var w, h int
	button2.GetSizeRequest(&w, &h)
	fmt.Println("button2 size w/h : ", w, "/", h)
	// 创建pixbuf，指定大小（宽度和高度
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("image/head.png", w-10, h-10, false)
	// pixbuf新建image
	image := gtk.NewImageFromPixbuf(pixbuf)

	// 释放pixbuf资源
	pixbuf.Unref()
	// 按钮设置image
	button2.SetImage(image)
	// 按钮信号处理
	button2.Connect("clicked", func() {
		fmt.Println("按钮2被按下")
	})

	window.Connect("destroy", gtk.MainQuit)
	window.ShowAll()
	gtk.Main()
}
