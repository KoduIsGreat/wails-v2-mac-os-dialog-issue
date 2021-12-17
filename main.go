package main

import (
	"embed"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"

	"github.com/wailsapp/wails/v2/pkg/options/mac"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/src
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "test",
		Width:             720,
		Height:            570,
		MinWidth:          720,
		MinHeight:         570,
		MaxWidth:          1280,
		MaxHeight:         740,
		Menu: menu.NewMenuFromItems(menu.SubMenu("File", menu.NewMenuFromItems(
			menu.Text("Save", keys.CmdOrCtrl("s"), func(_ *menu.CallbackData) {
				path, err := runtime.SaveFileDialog(app.ctx, runtime.SaveDialogOptions{
					DefaultDirectory:           "",
					DefaultFilename:            "test.yaml",
					Title:                      "Save",
					Filters:                    nil,
					ShowHiddenFiles:            false,
					CanCreateDirectories:       false,
					TreatPackagesAsDirectories: false,
				})
				runtime.LogInfo(app.ctx, path)
				if err != nil {
					runtime.LogError(app.ctx, err.Error())
				}
				runtime.LogInfo(app.ctx, path)
			}),
			menu.Separator(),
			menu.Text("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
				runtime.Quit(app.ctx)
			}),
		))),
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		RGBA:              &options.RGBA{R: 33, G: 37, B: 43, A: 255},
		Assets:            assets,
		LogLevel:          logger.DEBUG,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnShutdown:        app.shutdown,
		Bind: []interface{}{
			app,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
		},
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHiddenInset(),
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "Vanilla Template",
				Message: "Part of the Wails projects",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
