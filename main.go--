package main

import (
	"context"
	"embed"
	"fmt"
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
	"github.com/wanneSimon/saya-app/conf"
	"github.com/wanneSimon/saya-app/env"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	fileOp := &env.FileOp{}

	// configs
	var rootPath string = GetCurrentAbPath()
	// var rootPath string = "D:\\Git_Repo\\saya"
	// appConfig := strings.Join(rootPath, filepath.Separator, "config", filepath.Separator, "saya.yml")
	sp := fmt.Sprintf("%c", filepath.Separator)
	var appConfigPath = rootPath + sp + "config" + sp + "saya.yml"
	configOps := conf.NewConfigOpsAndLoad(appConfigPath)

	appConfig := configOps.Get()

	fmt.Println("appconfig", appConfig)

	// Create application with options
	err := wails.Run(&options.App{
		Title:     appConfig.Title, //"saya-app",
		Width:     1024,
		Height:    768,
		MinWidth:  930,
		MinHeight: 490,
		Frameless: appConfig.Frameless,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		// BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 255},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			app.transparentWinOS(appConfig.Title)
			fileOp.SetContext(ctx)
		},
		Bind: []interface{}{
			app, fileOp, configOps,
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
