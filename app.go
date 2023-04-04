package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/wanneSimon/saya-app/conf"
)

// App struct
type App struct {
	ctx    context.Context
	config conf.AppConfig
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

func (a *App) Config() conf.AppConfig {
	return a.config
}

func (a *App) SaveConfig(ac conf.AppConfig) bool {
	if ac.SaveAppConfig() {
		a.config = ac
		return true
	}
	return false
}
