package ui

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/mattn_diy/go-gtk/glib"
	"github.com/mattn_diy/go-gtk/gtk"
	"github.com/wrzfeijianshen/gui/TL/nrtools/gtk/serial"
	"unsafe"
	"github.com/wrzfeijianshen/gui/TL/nrtools/gtk/com"
)

var tabName chan string
var textview *gtk.TextView
var textBuffer *gtk.TextBuffer
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

func CreateBtn(notebook *gtk.Notebook, layout *gtk.Fixed, str string, width int,height int) ( *gtk.Button){
	prev := gtk.NewButtonWithLabel(str)
	prev.SetSizeRequest(80, 30)

	layout.Add(prev)
	layout.Put(prev, 5, 0)
	layout.Move(prev, width, height)
	return prev
}

func CreateLabel(notebook *gtk.Notebook, layout *gtk.Fixed, str string, width int, height int) (*gtk.Label){
	label := gtk.NewLabel(str)
	// prev.SetSizeRequest(80, 30)
	// label.ModifyFontSize(10)
	layout.Add(label)
	layout.Put(label, 5, 0)
	layout.Move(label, width, height)
	return  label
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
	port  := &serial.Port{}
	var c serial.Config

	fmt.Println("notebook: ", notebook, "layout ", layout)
	page := gtk.NewFrame(str)
	notebook.AppendPage(page, gtk.NewLabel(str))
	vbtn := gtk.NewFixed()

	lay := gtk.NewFixed()
	page.Add(lay) //把布局添加到主窗口中

	// 创建tab内容
	CreateLabel(notebook, lay, "串口:", 5, 20)

	CreateLabel(notebook, lay, "波特率:", 5, 50)
	CreateLabel(notebook, lay, "校验位:", 5, 80)
	CreateLabel(notebook, lay, "数据位:", 5, 110)
	CreateLabel(notebook, lay, "停止位:", 5, 140)
	boxName := CreateEditList(notebook, lay, "", 50, 20)
	for i := 1; i < 15; i++ {
		boxName.InsertText( i,"COM"+strconv.Itoa(i))
	}
	boxName.SetActive(0)

	// 波特率
	boxBaud := CreateEditList(notebook, lay, "", 50, 50)
	boxBaud.InsertText(0,"4800")
	boxBaud.InsertText(0,"4800")
	boxBaud.InsertText(1,"9600")
	boxBaud.InsertText( 2,"14400")
	boxBaud.InsertText(3,"38400")
	boxBaud.InsertText(4,"56000")
	boxBaud.InsertText(5,"57600")
	boxBaud.SetActive(0)

	// 校验位
	boxParity := CreateEditList(notebook, lay, "", 50, 80)
	boxParity.InsertText( 0,"NONE")
	boxParity.InsertText( 1,"ODD")
	boxParity.InsertText( 2,"EVEN")
	boxParity.InsertText(3,"MARK")
	boxParity.InsertText(4,"SPACE")
	boxParity.SetActive(0)
	// 数据位
	boxSize := CreateEditList(notebook, lay, "", 50, 110)
	boxSize.InsertText(0,"5")
	boxSize.InsertText(1,"6")
	boxSize.InsertText(2,"7")
	boxSize.InsertText(3,"8")
	boxSize.SetActive(3)

	// 停止位
	boxStopBits := CreateEditList(notebook, lay, "", 50, 140)
	boxStopBits.InsertText(0,"1")
	boxStopBits.InsertText(1,"1.5")
	boxStopBits.InsertText(2,"2")
	boxStopBits.SetActive(0)

	var strtemp string
	strtemp = "串口数据接收 : "
	lab := gtk.NewLabel(strtemp)
	lab.SetSizeRequest(300, 20)
	lay.Add(lab)
	lay.Put(lab, 5, 0)
	lay.Move(lab, 210, 0)

	frame2 := gtk.NewFrame("串口设置")
	frame2.SetSizeRequest(150,210)
	lay.Add(frame2)
	lay.Put(frame2, 5, 0)
	lay.Move(frame2, 0, 0)
	//framebox2 := gtk.NewVBox(false, 1)
	//frame2.Add(framebox2)

	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.SHADOW_IN)

	// 接收区
	var t gtk.TextTagTable

		textBuffer = gtk.NewTextBuffer(&t)
	textview = gtk.NewTextViewWithBuffer(*textBuffer)
	//textview = gtk.NewTextView()
	textview.SetSizeRequest(420,300)
	textview.SetBorderWidth(2)
	lay.Add(textview)
	lay.Put(textview, 5, 0)
	lay.Move(textview, 160, 30)
	swin.Add(textview)
	//textview.SetSensitive(false)

	// 发送区
	textviewSend := gtk.NewTextView()
	textviewSend.SetSizeRequest(340,100)
	textviewSend.SetBorderWidth(2)
	lay.Add(textviewSend)
	lay.Put(textviewSend, 5, 0)
	lay.Move(textviewSend, 160, 340)
	swin.Add(textviewSend)

	buffer := textviewSend.GetBuffer()
	var start, end gtk.TextIter
	buffer.GetStartIter(&start)
	buffer.Insert(&start, "Hello\n")

	SendBtn :=  CreateBtn(notebook, lay, "发送", 505,350)
	SendBtn.Clicked(func() {
		// 串口发送语句
		buffer.GetStartIter(&start)
		buffer.GetEndIter(&end)
		str := buffer.GetText(&start, &end,true)// 编辑框内容
		if port.Openflag() == false{
			return
		}
		port.Write([]byte(str))
	})

	// 按钮
	openBtn :=  CreateBtn(notebook, lay, "打开串口", 5,170)
	openBtnStat := false

	var startread, endread gtk.TextIter
	//var startreadup gtk.TextIter
	count := 0
	// 不断刷新接收区内容
	readCSer := func(chanread chan string){
		for ; ;  {
			if openBtnStat == false {
				return
			}
			buf := make([]byte, 2048)
			n, err := port.Read(buf[0:])
			if nil != err {
				continue
			}
			//fmt.Println("写了数据",string(buf[:n]))
			chanread <- string(buf[:n])

		}
	}
	goread := func (){
		var b1 gtk.ITextBuffer
		bufread := textview.GetBuffer()
		fmt.Println(b1)
		var chanread chan string
		chanread = make(chan string,1000)
		go readCSer(chanread)// 不断读取数据

		// 不断写数据
		for ; ;  {
			//fmt.Println(chanread,"buf : ", string(buf))
			//bufread.GetStartIter(&startreadup)
			if count >= 20{
				//time.Sleep(time.Second*1)
				count = 0
				//textview.CheckResize()
				//bufread.Delete(&startreadup,&endread)
				//startread = startreadup
				//endread = startreadup
				bufread.GetBounds(&startread,&endread)

				//bufread.GetStartIter(&startread)
				//bufread.GetEndIter(&endread)
				bufread.Delete(&startread,&endread)
				continue
			}

			bufread.GetBounds(&startread,&endread)
			//str3 := startread.GetText(&endread)
			//fmt.Println("str3",str3)
			//str2 := fmt.Sprintf("hello %d\n",count)
			count++
			//str2 := "hello\n"
			//bufread.SetText(str2)
			//bufread.GetStartIter(&startread)
			//str :=  <- chanread
			//fmt.Println(chanread,"数据读完  : ", str)

			bufread.Insert(&endread, <- chanread)
			//bufread.GetEndIter(&endread)
		}
	}

	tmp := ""
	openBtn.Clicked(func() {
		strtemp = "串口数据接收 : "

		if openBtnStat== false{
			openBtnStat = true
			// 1.串口信息配置

			c.Name = boxName.GetActiveText()
			c.Baud, _ = strconv.Atoi(boxBaud.GetActiveText())
			c.Size = com.StringTouint8(boxSize.GetActiveText())
			c.Parity = serial.ByteToParity(boxParity.GetActiveText()[0])
			c.StopBits = serial.StringStop(boxStopBits.GetActiveText())
			//fmt.Println(c.Name, c.Baud, string(c.Parity), c.StopBits, string(c.Size))
			var err error

			port,err = serial.OpenPort(&c)
			if err != nil {
				tmp = fmt.Sprintf("打开串口%s失败：%s。\n",c.Name, err)
				return
			}
			tmp = fmt.Sprintf("串口%s成功开启\n",c.Name)
			openBtn.SetLabel("关闭串口")

			// 禁用控件
			boxName.SetSensitive(false)
			boxBaud.SetSensitive(false)
			boxParity.SetSensitive(false)
			boxSize.SetSensitive(false)
			boxStopBits.SetSensitive(false)
			go goread()
		} else {
			openBtnStat = false

			openBtn.SetLabel("打开串口")
			boxName.SetSensitive(true)
			boxBaud.SetSensitive(true)
			boxParity.SetSensitive(true)
			boxSize.SetSensitive(true)
			boxStopBits.SetSensitive(true)

			err := port.Close()
			if err != nil {
				tmp = fmt.Sprintf("串口%s关闭失败：%s。\n",c.Name,err)
			}
			tmp = fmt.Sprintf("串口%s成功关闭\n",c.Name)
		}

		strtemp += tmp
		lab.SetLabel(strtemp)
	})

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

	prevPage := CreateBtn(notebook, layout, "上一个", 0,35)
	prevPage.Clicked(func() {
		notebook.PrevPage()
	})
	nextPage := CreateBtn(notebook, layout, "下一个", 90,35)
	nextPage.Clicked(func() {
		notebook.NextPage()
	})

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
