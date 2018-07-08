package main

import (
	"os"

	"github.com/mattn_diy/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)

	// 主窗口
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("控件")
	window.SetSizeRequest(300, 300)

	// 创建布局
	layout := gtk.NewFixed()
	// 创建按钮
	b1 := gtk.NewButton()
	b1.SetLabel("b1嘿嘿")
	b2 := gtk.NewButtonWithLabel("b2哈哈")
	b2.SetSizeRequest(100, 80)

	// window窗口添加布局，布局添加容器
	window.Add(layout)
	layout.Put(b1, 10, 10)
	layout.Move(b1, 20, 20) // 移动位置时，先put
	layout.Put(b2, 50, 50)

	window.ShowAll() // 显示所有的控件

	gtk.Main() // 主事件循环
}
