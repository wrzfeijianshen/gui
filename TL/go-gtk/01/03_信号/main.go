package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-gtk/glib"

	"github.com/mattn_diy/go-gtk/gtk"
)

var count int

func HandleButton(ctx *glib.CallbackContext) {
	arg := ctx.Data()   // 获取用户传递的参数
	p, ok := arg.(*int) // 类型断言
	if ok {
		fmt.Println("count = : ", *p)
		*p = 222
	}
	fmt.Println("按钮b1 按下了")
	//gtk.MainQuit() //关闭gtk程序
}
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

	// 绑定信号
	count = 1
	b1.Connect("pressed", HandleButton, &count)

	// 推荐匿名函数
	b2.Connect("pressed", func() {
		fmt.Println("b2 按下了，count : ", count)

	})

	window.ShowAll() // 显示所有的控件
	gtk.Main()       // 主事件循环
}
