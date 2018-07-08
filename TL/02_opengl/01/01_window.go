package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
)

// GLFW 需要在其被初始化之后的线程里被调用
func init() {
	runtime.LockOSThread()
}

func main() {

	// 初始化glfw 调用c底层的C.glfwInit()进行初始化
	err := glfw.Init()
	if err != nil {
		// 初始化失败返回
		fmt.Println(err)
		return
	}

	// 使用完毕后 终止glfw
	defer glfw.Terminate()

	// GFLW创建的OpenGL上下文可能是任何版本的,设置最低版本
	Version := 3
	if Version < 3 {
		glfw.WindowHint(glfw.ContextVersionMajor, 2) // OpenGL主版本号
		glfw.WindowHint(glfw.ContextVersionMinor, 1) // OpenGL副版本号
	} else {
		glfw.WindowHint(glfw.ContextVersionMajor, 3)                // OpenGL主版本号
		glfw.WindowHint(glfw.ContextVersionMinor, 2)                // OpenGL副版本号
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile) // 3.2
	}

	// 创建窗口 宽度 高度 标题 显示模式(nil为窗口化) 窗口共享资源
	window, err := glfw.CreateWindow(300, 300, "window", nil, nil)
	if err != nil {
		// 创建失败返回
		fmt.Println(err)
		return
	}

	// 使用OpenGL API之前,必须创建opengl上下文
	window.MakeContextCurrent()

	// 检查窗口是否关闭
	for !window.ShouldClose() {
		// 此时窗口还在运行

		// 交换缓冲区，即在window上更新内容
		// c/c++中 glfwSwapBuffers(window);
		// -->glfwSwapBuffers-->flfw.SwapBuffers()
		// 参数为window，所以window函数中也实现了SwapBuffers接口函数
		window.SwapBuffers()

		// 轮询事件 glfwPollEvents();
		glfw.PollEvents()
	}
}
