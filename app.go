package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	sysRuntime "runtime"
	"syscall"

	"github.com/lxn/win"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Context() context.Context {
	return a.ctx
}

func (a *App) Exit() {
	log.Println("Exit...")
	os.Exit(0)
}

func (a *App) Minimises() {
	runtime.WindowMinimise(a.ctx)
}

// 让窗口完全透明化
func (a *App) transparentWinOS(title string) {
	ti, _ := syscall.UTF16PtrFromString(title)
	hwnd := win.FindWindow(nil, ti)
	// hwnd := win.FindWindow(nil, syscall.UTF16PtrFromString(title))
	win.SetWindowLong(hwnd, win.GWL_EXSTYLE, win.GetWindowLong(hwnd, win.GWL_EXSTYLE)|win.WS_EX_LAYERED)
}

// 不同平台启动指令不同 https://www.lmlphp.com/user/365130/article/item/8221378
var OpenLinkCommands = map[string]string{
	"windows": "start /c",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func (a *App) OpenUrl(uri string) error {
	// runtime.GOOS获取当前平台
	gos := sysRuntime.GOOS
	run, ok := OpenLinkCommands[gos]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", gos)
	}

	var cmd *exec.Cmd
	if gos == "windows" {
		cmd = exec.Command("cmd", run+" "+uri)
	} else {
		cmd = exec.Command(run, uri)
	}
	return cmd.Run()
}
