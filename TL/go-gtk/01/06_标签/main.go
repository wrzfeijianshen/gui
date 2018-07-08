package main

import (
	"fmt"
	"os"

	"github.com/mattn_diy/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()
	builder.AddFromFile("label.glade")

	window := gtk.WindowFromObject(builder.GetObject("window1"))

	//获取label控件
	labelOne := gtk.LabelFromObject(builder.GetObject("label1"))

	fmt.Println("labelOne = ", labelOne.GetText()) // 获取label内容
	labelOne.SetText("你大爷")                        //设置内容

	//按窗口关闭按钮，自动触发"destroy"信号
	window.Connect("destroy", gtk.MainQuit)
	window.Show()

	gtk.Main()
}
