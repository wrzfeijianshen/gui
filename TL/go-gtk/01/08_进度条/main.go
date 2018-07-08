package main

import (
	"fmt"
	"os"

	"github.com/mattn_diy/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)
	builder := gtk.NewBuilder()
	builder.AddFromFile("bar.glade")
	window := gtk.WindowFromObject(builder.GetObject("window1"))
	window.SetSizeRequest(500, 500)

	bar := gtk.ProgressBarFromObject(builder.GetObject("progressbar1"))
	bar.SetFraction(0.5) // 设置进度

	bar.SetText("%50") // 设置进度条显示的文本
	// 获取进度
	value := bar.GetFraction()
	fmt.Println("value : ", value)

	window.Connect("destroy", gtk.MainQuit)
	window.ShowAll()
	gtk.Main()
}
