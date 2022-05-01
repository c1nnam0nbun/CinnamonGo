package window

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

type window struct {
	glfw.Window
}

type Window interface {
	Update() bool
}

func New(desc Descriptor) Window {
	var w window
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	window, err := glfw.CreateWindow(int(desc.Width), int(desc.Height), desc.Title, nil, nil)
	if err != nil {
		panic(err)
	}

	w.Window = *window
	return &w
}

func (w *window) Update() bool {
	if !w.ShouldClose() {
		w.SwapBuffers()
		glfw.PollEvents()
		return true
	}
	return false
}

type Descriptor struct {
	Width   uint32
	Height  uint32
	Title   string
	IsVsync bool
}

func DefaultDescriptor() Descriptor {
	return Descriptor{
		Width:   1280,
		Height:  720,
		Title:   "Cinnamon",
		IsVsync: true,
	}
}
