package main

import (
	"context"
	"fmt"
	"log"
	"os"

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
