package main

import (
	"runtime"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
	"github.com/go-gl/gl/v4.6-core/gl"
)
// 参考博文 https://blog.csdn.net/meltlao2/article/details/80854788
func init(){
	runtime.LockOSThread()
}
func main() {
	// 初始化glfw库
	if err := glfw.Init(); err != nil {
		//panic(err)
		log.Println(" glfw.Init err : ",err)
		return
	}

	// 创建一个窗口和它的OpenGL上下文

	// gflw 创建opengl上下文可能是任何版本的，可设置opengl的最低版本
	glfw.WindowHint(glfw.ContextVersionMajor,2) // opengl主版本号
	glfw.WindowHint(glfw.ContextVersionMinor,0) // opengl 副版本号
	glfw.WindowHint(glfw.OpenGLProfile,glfw.OpenGLAnyProfile)
	// 参数1: 窗口宽度 2 窗口高度 3 窗口标题 4 显示模式，5 窗口共享资源
	window,err := glfw.CreateWindow(500,500,"Hello world",nil,nil)
	if err != nil {
		log.Println(" glfw.CreateWindow err : ",err)
		glfw.Terminate()// 没有创建会返回nil
		return
	}
	//终止GLFW
	//当你使用完毕后，请在程序退出前及时终止它
	defer glfw.Terminate()

	//  设置当前窗口上下文
	// 使用openglAPI之前，必须设置好当前opengl上下文，此上下文一直保持，除非设置另外一个上下文或者此窗口被销毁
	window.MakeContextCurrent()

	initOpenGL()
	// 循环，直到用户关闭窗口

	for !window.ShouldClose()  {
		draw()
		// 渲染
		//gl.Clear(gl.DEPTH_BUFFER_BIT)
		// 交换缓冲区 在windows上个更新内容
		window.SwapBuffers()

		// 轮询事件
		glfw.PollEvents()
	}
	//终止GLFW
	//当你使用完毕后，请在程序退出前及时终止它
	glfw.Terminate()
}

func initOpenGL() {
	if err := gl.Init(); err != nil{
		log.Println(" gl.Init() err : ",err)

	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version",version)
}

func draw(){
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}