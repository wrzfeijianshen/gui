package main

import (
	"fmt"
	"os"

	"github.com/mattn_diy/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)
	builder := gtk.NewBuilder()         // 新建Builder
	builder.AddFromFile("1_test.glade") // 读取glade文件
	window := gtk.WindowFromObject(builder.GetObject("window1"))
	b1 := gtk.ButtonFromObject(builder.GetObject("btn1"))
	b2 := gtk.ButtonFromObject(builder.GetObject("button2"))

	// 信号处理
	b1.Connect("clicked", func() {
		fmt.Println("button txt : ", b1.GetLabel())
	})

	b2.Connect("clicked", func() {
		fmt.Println("button txt : ", b2.GetLabel())
		gtk.MainQuit() // 关闭窗口
	})
	//按窗口关闭按钮，自动触发"destroy"信号
	window.Connect("destroy", gtk.MainQuit)
	window.ShowAll()
	gtk.Main()
}
