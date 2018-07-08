package main

import (
	"os"

	"github.com/mattn_diy/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)
	builder := gtk.NewBuilder()
	builder.AddFromFile("window.glade")
	window := gtk.WindowFromObject(builder.GetObject("window1"))

	window.SetSizeRequest(300, 300)        // 设置窗口大小
	window.SetTitle("窗口")                  // 设置标题
	window.SetIconFromFile("face.png")     // 设置icon
	window.SetResizable(false)             // 设置不可伸缩
	window.SetPosition(gtk.WIN_POS_CENTER) // 设置居中显示

	// 关闭时触发"destroy"
	window.Connect("destroy", gtk.MainQuit)
	window.ShowAll()
	gtk.Main()
}
