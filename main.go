package main

import (
	"context"
	"fmt"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

// App 应用结构体
type App struct {
	ctx context.Context
}

// NewApp 创建应用实例
func NewApp() *App {
	return &App{}
}

// OnStartup 应用启动时调用
func (a *App) OnStartup(ctx context.Context) {
	a.ctx = ctx
}

// OnShutdown 应用关闭时调用
func (a *App) OnShutdown(ctx context.Context) {
}

// OnDomReady DOM准备就绪时调用
func (a *App) OnDomReady(ctx context.Context) {
}

// Greet 问候方法示例
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, Welcome to Go-Trade!", name)
}

// main 主函数
func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:  "Go-Trade 校园闲置交易",
		Width:  1280,
		Height: 800,
		AssetServer: &assetserver.Options{
			Assets: nil,
		},
		BackgroundColour: &options.RGBA{R: 245, G: 247, B: 250, A: 1},
		OnStartup:        app.OnStartup,
		OnShutdown:       app.OnShutdown,
		OnDomReady:       app.OnDomReady,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			Theme:                windows.SystemDefault,
		},
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarDefault(),
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
