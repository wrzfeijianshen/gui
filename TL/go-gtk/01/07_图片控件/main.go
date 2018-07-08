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
	builder.AddFromFile("image.glade")
	window := gtk.WindowFromObject(builder.GetObject("window1"))
	// 获取image控件
	image1 := gtk.ImageFromObject(builder.GetObject("image1"))
	// 获取image控件大小
	var w, h int
	image1.GetSizeRequest(&w, &h)
	fmt.Println(w, h)
	// 创建pixbuf，指定大小（宽度和高度），image有多大就设置多大
	// 最后一个参数false代表不保存图片原来的尺寸
	pixbuf1, _ := gdkpixbuf.NewPixbufFromFileAtScale("image/face.png", w, h, false)
	// image设置pixbuf
	image1.SetFromPixbuf(pixbuf1)
	// pixbuf1使用完毕，需要释放资源
	pixbuf1.Unref()

	window.Connect("destroy", gtk.MainQuit)
	window.ShowAll()
	gtk.Main()
}
