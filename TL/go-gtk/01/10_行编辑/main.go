package main

import (
	"fmt"
	"os"

	"github.com/mattn_diy/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)
	builder := gtk.NewBuilder()
	builder.AddFromFile("entry.glade")
	window := gtk.WindowFromObject(builder.GetObject("window1"))
	window.SetSizeRequest(500, 500)
	entry := gtk.EntryFromObject(builder.GetObject("entry1"))

	entry.SetText("77877你猜")                      // 设置内容
	fmt.Println("entry text = ", entry.GetText()) // 获取内容
	// entry.SetVisibility(false)                  // 设置不可见字符，即密码模式
	// entry.SetEditable(false)                    // 只读，不可编辑
	entry.ModifyFontSize(30) //修改字体大小

	// 信号处理，当用户在文本输入控件内部按回车键时引发activate信号
	entry.Connect("activate", func() {
		fmt.Println("entry text = ", entry.GetText()) // 获取内容
	})

	window.Connect("destroy", gtk.MainQuit)
	window.ShowAll()
	gtk.Main()
}
