package main

import (
	"os"

	"github.com/mattn_diy/go-gtk/gtk"
)

func main() {

	gtk.Init(&os.Args) // gtk 初始化

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL) // 创建窗口
	window.SetPosition(gtk.WIN_POS_CENTER)       // 窗口居中
	window.SetTitle("hello gtk")                 // 设置标题
	window.SetSizeRequest(300, 300)              // 设置窗口宽度和高度

	window.Show() // 显示窗口
	gtk.Main()    // 主事件循环
}
