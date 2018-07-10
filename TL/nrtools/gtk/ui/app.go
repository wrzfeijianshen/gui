package ui

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"github.com/wrzfeijianshen/nrtools/gtk/serial"
	"unsafe"
	"github.com/wrzfeijianshen/nrtools/gtk/com"
)

var tabName chan string

func CreateWindowEdit() *gtk.Window {
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetDefaultSize(320, 200)
	layout := gtk.NewFixed()
	window.Add(layout) //把布局添加到主窗口中
	label := gtk.NewLabel("请输入标签名称:")
	label.SetSizeRequest(300, 40)
	// label.ModifyFontSize(20)
	layout.Add(label)
	layout.Put(label, 0, 0)  //设置按钮在容器的位置
	layout.Move(label, 0, 0) //设置按钮在容器的位置

	window.SetPosition(gtk.WIN_POS_CENTER)
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		fmt.Println("destroy pending...")
		gtk.MainQuit()
	}, "foo")

	//获取entry控件
	entry := gtk.NewEntry()

	entry.SetText("示例")                           //设置内容
	fmt.Println("entry text = ", entry.GetText()) //获取内容
	// entry.ModifyFontSize(20)                      //修改字体大小
	entry.SetSizeRequest(300, 40)
	layout.Put(entry, 0, 50)  //设置按钮在容器的位置
	layout.Move(entry, 0, 50) //设置按钮在容器的位置
	layout.Add(entry)

	//信号处理，当用户在文本输入控件内部按回车键时引发activate信号
	entry.Connect("activate", func() {
		fmt.Println("entry text = ", entry.GetText()) //获取内容
	})

	btn := gtk.NewButton()
	btn.SetLabel("确定添加")

	btn.Connect("clicked", func() {
		str := entry.GetText()
		fmt.Println(str)
		tabName = make(chan string, 200)
		tabName <- str
		window.Hide()
	})
	btn.SetSizeRequest(80, 30)
	layout.Put(btn, 0, 110)  //设置按钮在容器的位置
	layout.Move(btn, 0, 110) //设置按钮在容器的位置
	layout.Add(btn)

	window.Add(layout)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		fmt.Println("destroy pending...")
		window.Hide()
	}, "foo")
	window.ShowAll()
	return window
}
func CreateWindow() *gtk.Window {
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetDefaultSize(700, 300)
	layout := gtk.NewFixed()
	window.Add(layout) //把布局添加到主窗口中

	vbox := gtk.NewVBox(false, 1)
	CreateMenuAndToolbar(window, vbox)
	window.Add(vbox)
	layout.Put(vbox, 0, 0)  //设置按钮在容器的位置
	layout.Move(vbox, 0, 0) //设置按钮在容器的位置

	CreateNotbook(window, layout)
	return window
}

func CreateBtn(notebook *gtk.Notebook, layout *gtk.Fixed, str string, width int,height int, btnfun interface{}, args ...interface{}) ( *gtk.Button){
	prev := gtk.NewButtonWithLabel(str)
	prev.SetSizeRequest(80, 30)
	prev.Clicked(func() {
		if len(args) > 1 {
			btnfun.(func(...interface{}))(args)
		} else if len(args) == 1 {
			btnfun.(func(interface{}))(args[0])
		} else {
			btnfun.(func())()
		}
	})
	layout.Add(prev)
	layout.Put(prev, 5, 0)
	layout.Move(prev, width, height)
	return prev
}

func CreateLabel(notebook *gtk.Notebook, layout *gtk.Fixed, str string, width int, height int) {
	label := gtk.NewLabel(str)
	// prev.SetSizeRequest(80, 30)
	// label.ModifyFontSize(10)
	layout.Add(label)
	layout.Put(label, 5, 0)
	layout.Move(label, width, height)
}

func CreateEditList(notebook *gtk.Notebook, layout *gtk.Fixed, str string, width int, height int) *gtk.ComboBoxText {
	box := gtk.NewComboBoxTextWithEntry()
	box.SetSizeRequest(80, 20)
	// prev.SetSizeRequest(80, 30)
	// box.ModifyFontSize(10)
	// box.SetTextColumn(0)
	layout.Add(box)
	layout.Put(box, 5, 0)
	layout.Move(box, width, height)
	return box
}

func CreateNotpage(notebook *gtk.Notebook, layout *gtk.Fixed, str string) {
	fmt.Println("notebook: ", notebook, "layout ", layout)
	page := gtk.NewFrame(str)
	notebook.AppendPage(page, gtk.NewLabel(str))
	vbtn := gtk.NewFixed()

	lay := gtk.NewFixed()
	page.Add(lay) //把布局添加到主窗口中

	// 创建tab内容
	CreateLabel(notebook, lay, "串口:", 0, 0)
	CreateLabel(notebook, lay, "波特率:", 0, 25)
	CreateLabel(notebook, lay, "校验位:", 0, 50)
	CreateLabel(notebook, lay, "数据位:", 0, 75)
	CreateLabel(notebook, lay, "停止位:", 0, 100)
	boxName := CreateEditList(notebook, lay, "", 50, 0)
	for i := 1; i < 15; i++ {
		boxName.InsertText( i,"COM"+strconv.Itoa(i))
	}
	boxName.SetActive(0)

	// 波特率
	boxBaud := CreateEditList(notebook, lay, "", 50, 25)
	boxBaud.InsertText(0,"4800")
	boxBaud.InsertText(0,"4800")
	boxBaud.InsertText(1,"9600")
	boxBaud.InsertText( 2,"14400")
	boxBaud.InsertText(3,"38400")
	boxBaud.InsertText(4,"56000")
	boxBaud.InsertText(5,"57600")
	boxBaud.SetActive(0)

	// 校验位
	boxParity := CreateEditList(notebook, lay, "", 50, 50)
	boxParity.InsertText( 0,"NONE")
	boxParity.InsertText( 1,"ODD")
	boxParity.InsertText( 2,"EVEN")
	boxParity.InsertText(3,"MARK")
	boxParity.InsertText(4,"SPACE")
	boxParity.SetActive(0)
	// 数据位
	boxSize := CreateEditList(notebook, lay, "", 50, 75)
	boxSize.InsertText(0,"5")
	boxSize.InsertText(1,"6")
	boxSize.InsertText(2,"7")
	boxSize.InsertText(3,"8")
	boxSize.SetActive(3)

	// 停止位
	boxStopBits := CreateEditList(notebook, lay, "", 50, 100)
	boxStopBits.InsertText(0,"1")
	boxStopBits.InsertText(1,"1.5")
	boxStopBits.InsertText(2,"2")
	boxStopBits.SetActive(0)

	port  := &serial.Port{}

	btnfunc := func() {
		// 1.串口信息配置
		var c serial.Config
		c.Name = boxName.GetActiveText()
		c.Baud, _ = strconv.Atoi(boxBaud.GetActiveText())
		c.Size = com.StringTouint8(boxSize.GetActiveText())
		c.Parity = serial.ByteToParity(boxParity.GetActiveText()[0])
		c.StopBits = serial.StringStop(boxStopBits.GetActiveText())
		fmt.Println(c.Name, c.Baud, string(c.Parity), c.StopBits, string(c.Size))
		var err error
		port,err = serial.OpenPort(&c)
		if err != nil {
			fmt.Println("打开串口失败：", err)
			return
		}
		//port.Read("aa")
		fmt.Println("串口成功开启：")
	}

	btncolosfunc := func() {
		// 1.串口信息配置
		err := port.Close()
		if err != nil {
			fmt.Println("串口关闭失败：",err)
		}
		com.StrtoByte("aa")
	}
	// 按钮
	 CreateBtn(notebook, lay, "打开串口", 0,125, btnfunc)
	//btn.SetTooltipText("关闭串口")
	CreateBtn(notebook, lay, "关闭串口", 0,160, btncolosfunc)
	page.Add(vbtn)
	notebook.ShowAll()
}
func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func CreateNotbook(w *gtk.Window, layout *gtk.Fixed) {

	notebook := gtk.NewNotebook()
	notebook.SetSizeRequest(600, 600)
	CreateNotpage(notebook, layout, "串口")

	CreateBtn(notebook, layout, "上一个", 0,35, notebook.PrevPage)
	CreateBtn(notebook, layout, "下一个", 90,35, notebook.NextPage)

	prev := gtk.NewButtonWithLabel("创建示例")
	prev.SetSizeRequest(80, 30)
	flag := false
	win := &gtk.Window{}
	str := "示例"
	prev.Clicked(func() {
		w.ShowAll()
		if flag == false {
			win = CreateWindowEdit()
			flag = true
		}
	})
	layout.Add(prev)
	layout.Put(prev, 5, 0)
	layout.Move(prev, 180, 35)

	showbtn := gtk.NewButtonWithLabel("显示创建")
	showbtn.SetSizeRequest(80, 30)
	showbtn.Clicked(func() {
		if flag == true {
			str = <-tabName
			fmt.Println(str)
			CreateNotpage(notebook, layout, str)
			w.ShowAll()
		}
		layout.Show()
		w.ShowAll()
	})
	layout.Add(showbtn)
	layout.Put(showbtn, 5, 0)
	layout.Move(showbtn, 270, 35)

	w.Add(notebook)
	layout.Put(notebook, 0, 0)   //设置按钮在容器的位置
	layout.Move(notebook, 0, 80) //设置按钮在容器的位置

}
func CreateMenuAndToolbar(w *gtk.Window, vbox *gtk.VBox) {
	action_group := gtk.NewActionGroup("my_group")
	ui_manager := CreateUIManager()
	accel_group := ui_manager.GetAccelGroup()
	w.AddAccelGroup(accel_group)
	AddFileMenuActions(action_group)
	AddEditMenuActions(action_group)
	ui_manager.InsertActionGroup(action_group, 0)
	menubar := ui_manager.GetWidget("/MenuBar")
	vbox.PackStart(menubar, false, false, 0)
	eventbox := gtk.NewEventBox()
	vbox.PackStart(eventbox, false, false, 0)
}

func CreateUIManager() *gtk.UIManager {
	UI_INFO := `
<ui>
  <menubar name='MenuBar'>
    <menu action='FileMenu'>
      <menu action='FileNew'>
        <menuitem action='FileNewStandard' />
        <menuitem action='FileNewFoo' />
        <menuitem action='FileNewGoo' />
      </menu>
      <separator />
      <menuitem action='FileQuit' />
    </menu>
    <menu action='EditMenu'>
      <menuitem action='EditCopy' />
      <menuitem action='EditPaste' />
      <menuitem action='EditSomething' />
    </menu>
    <menu action='ChoicesMenu'>
      <menuitem action='ChoiceOne'/>
      <menuitem action='ChoiceTwo'/>
      <menuitem action='ChoiceThree'/>
      <separator />
      <menuitem action='ChoiceToggle'/>
    </menu>
  </menubar>
  <toolbar name='ToolBar'>
    <toolitem action='FileNewStandard' />
    <toolitem action='FileQuit' />
  </toolbar>
  <popup name='PopupMenu'>
    <menuitem action='EditCopy' />
    <menuitem action='EditPaste' />
    <menuitem action='EditSomething' />
  </popup>
</ui>
`
	ui_manager := gtk.NewUIManager()
	ui_manager.AddUIFromString(UI_INFO)
	return ui_manager
}

func OnMenuFileNewGeneric() {
	fmt.Println("A File|New menu item was selected.")
}

func OnMenuFileQuit() {
	fmt.Println("quit app...")
	gtk.MainQuit()
}

func OnMenuOther(ctx *glib.CallbackContext) {
	v := reflect.ValueOf(ctx.Target())
	if v.Kind() == reflect.Ptr {
		fmt.Printf("Item %s(%p) was selected", v.Elem(), v.Interface())
		fmt.Println()
		if w, ok := v.Elem().Interface().(gtk.IWidget); ok {
			v := reflect.ValueOf(ctx.Target())
			v2 := v.Elem()
			fmt.Println(v.Kind(), v2.Kind())
			fmt.Println("Menu item ", w.GetName(), " was selected")
		}
	}
}

func AddFileMenuActions(action_group *gtk.ActionGroup) {
	action_group.AddAction(gtk.NewAction("FileMenu", "File", "", ""))

	action_filenewmenu := gtk.NewAction("FileNew", "", "", gtk.STOCK_NEW)
	action_group.AddAction(action_filenewmenu)

	action_new := gtk.NewAction("FileNewStandard", "_New",
		"Create a new file", gtk.STOCK_NEW)
	action_new.Connect("activate", OnMenuFileNewGeneric)
	action_group.AddActionWithAccel(action_new, "")

	action_new_foo := gtk.NewAction("FileNewFoo", "New Foo",
		"Create new foo", gtk.STOCK_NEW)
	action_new_foo.Connect("activate", OnMenuFileNewGeneric)
	action_group.AddAction(action_new_foo)

	action_new_goo := gtk.NewAction("FileNewGoo", "_New Goo",
		"Create new goo", gtk.STOCK_NEW)
	action_new_goo.Connect("activate", OnMenuFileNewGeneric)
	action_group.AddAction(action_new_goo)

	action_filequit := gtk.NewAction("FileQuit", "", "", gtk.STOCK_QUIT)
	action_filequit.Connect("activate", OnMenuFileQuit)
	action_group.AddActionWithAccel(action_filequit, "")
}

func AddEditMenuActions(action_group *gtk.ActionGroup) {
	action_group.AddAction(gtk.NewAction("EditMenu", "Edit", "", ""))

	action_editcopy := gtk.NewAction("EditCopy", "", "", gtk.STOCK_COPY)
	action_editcopy.Connect("activate", OnMenuOther)
	action_group.AddActionWithAccel(action_editcopy, "")

	action_editpaste := gtk.NewAction("EditPaste", "", "", gtk.STOCK_PASTE)
	action_editpaste.Connect("activate", OnMenuOther)
	action_group.AddActionWithAccel(action_editpaste, "")

	action_editsomething := gtk.NewAction("EditSomething", "Something", "", "")
	action_editsomething.Connect("activate", OnMenuOther)
	action_group.AddActionWithAccel(action_editsomething, "<control><alt>S")
}

func AppMain() {
	gtk.Init(nil)
	window := CreateWindow()
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		fmt.Println("destroy pending...")
		gtk.MainQuit()
	}, "foo")
	window.ShowAll()
	gtk.Main()
}
