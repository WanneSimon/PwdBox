package main

import (
	"context"
	"embed"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	// "github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/wanneSimon/pwdbox/internal/conf"
	dataop "github.com/wanneSimon/pwdbox/internal/data-op"
	"github.com/wanneSimon/pwdbox/internal/env"
	"github.com/wanneSimon/pwdbox/internal/pwdbox"
)

//go:embed all:frontend/dist
var assets embed.FS

//	func main() {
//		pwdbox.TestAES()
//	}
var sp = fmt.Sprintf("%c", filepath.Separator)
var ConfigFolder string = "config" + sp

func main() {
	initLog()
	// Create an instance of the app structure
	app := NewApp()
	fileOp := env.FileOp{} // 文件操作

	dbop := pwdbox.DbOp{} // 数据库操作

	// configs
	var rootPath string = GetCurrentAbPath()
	// var rootPath string = "D:\\Git_Repo\\saya"
	// appConfig := strings.Join(rootPath, filepath.Separator, "config", filepath.Separator, "saya.yml")
	// sp := fmt.Sprintf("%c", filepath.Separator)
	var appConfigPath = rootPath + sp + ConfigFolder + "saya.yml"
	configOps := conf.NewConfigOpsAndLoad(appConfigPath) // 配置操作

	appConfig := configOps.Get() // 配置

	pwdTool := pwdbox.PwdTool{}     // 加密工具
	dataOutOp := dataop.DataOutOp{} // 数据导出

	// fmt.Println("appconfig", appConfig)

	// Create application with options
	err := wails.Run(&options.App{
		Title: appConfig.Title, //"pwdbox",
		// Width:     1024,
		// Height:    768,
		Width:     930,
		Height:    660,
		MinWidth:  930,
		MinHeight: 660,
		Frameless: appConfig.Frameless,
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: NewFileLoader(),
		},
		// BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 255},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			app.transparentWinOS(appConfig.Title)
			fileOp.SetContext(ctx)

			pwdbox.InitSqlite(appConfig.Pwdbox)
		},
		Bind: []interface{}{
			app, configOps, &fileOp, &dbop,
			&(pwdbox.PlatformServiceInstance), &pwdbox.AccountServiceInstance,
			&pwdTool, &dataOutOp,
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: appConfig.Debug,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			// WindowIsTranslucent:               true,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "userdata",
			Theme:                             windows.SystemDefault,
			CustomTheme: &windows.ThemeSettings{
				DarkModeTitleBar:   windows.RGB(20, 20, 20),
				DarkModeTitleText:  windows.RGB(200, 200, 200),
				DarkModeBorder:     windows.RGB(20, 0, 20),
				LightModeTitleBar:  windows.RGB(200, 200, 200),
				LightModeTitleText: windows.RGB(20, 20, 20),
				LightModeBorder:    windows.RGB(200, 200, 200),
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

// 获取当前路径
func GetCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// 日志
func initLog() {
	consoleOut := os.Stdout
	file, err := os.OpenFile("logs/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	var outWriters []io.Writer
	if err == nil {
		outWriters = []io.Writer{consoleOut, file}
	} else {
		outWriters = []io.Writer{consoleOut}
	}

	_, err2 := os.Stat("logs")
	if err2 != nil {
		if os.IsNotExist(err2) {
			os.MkdirAll("logs", os.ModePerm)
		}
	}

	multiWriter := io.MultiWriter(outWriters...)
	log.SetOutput(multiWriter)
}
