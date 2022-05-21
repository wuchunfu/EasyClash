package main

import (
	"embed"
	"fmt"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"easyclash/app"
	"easyclash/pkg/notify"
	"easyclash/pkg/stdpath"
	"easyclash/pkg/utils"
)

//go:embed build/appicon.png
var icon []byte

//go:embed frontend/dist
var ui embed.FS

//go:embed .easy_clash
var ruleSet embed.FS

var appName = "EasyClash"

func main() {
	path, _ := stdpath.AppDataLocation(appName)
	err := utils.SaveDir(ruleSet, path, false)
	if err != nil {
		_ = notify.Notification("EasyClash", "", "init conf dir error", "")
		log.Fatal(fmt.Errorf("AppDataLocation dir %s init err %v", path, err))
		return
	}

	_app := app.NewApp(path)
	opt := appOptions()
	opt.OnStartup = _app.Startup
	opt.OnDomReady = _app.DomReady
	opt.OnShutdown = _app.Shutdown
	opt.OnBeforeClose = _app.OnBeforeClose
	opt.Bind = []interface{}{
		_app,
	}

	err = wails.Run(opt)
	if err != nil {
		log.Fatal(err)
	}
}

func appOptions() *options.App {
	return &options.App{
		Title:             appName,
		Width:             1024,
		Height:            768,
		MinWidth:          985,
		MinHeight:         768,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: true,
		RGBA:              &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		Assets:            ui,
		LogLevel:          logger.DEBUG,
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  true,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameVibrantLight,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   appName,
				Message: "clash config manager",
				Icon:    icon,
			},
		},
	}
}
