package main

import (
	"bytes"
	"errors"
	"image/png"
	"log"

	ui "blink-demo/ui/bin"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/jannson/walk"
	"github.com/raintean/blink"
)

func notifyIcon(mw *walk.MainWindow) *walk.NotifyIcon {
	imgb, err := ui.Asset("app.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _ := png.Decode(bytes.NewReader(imgb))
	if img == nil {
		log.Fatal(errors.New("decode error"))
	}
	icon, err := walk.NewIconFromImage(img)
	//icon, err := walk.Resources.Icon("img/rc.ico")
	if err != nil {
		log.Fatal(err)
	}

	// Create the notify icon and make sure we clean it up on exit.
	ni, err := walk.NewNotifyIcon(mw)
	if err != nil {
		log.Fatal(err)
	}

	// Set the icon and a tool tip text.
	if err := ni.SetIcon(icon); err != nil {
		log.Fatal(err)
	}
	if err := ni.SetToolTip("Click for info or use the context menu to exit."); err != nil {
		log.Fatal(err)
	}

	// When the left mouse button is pressed, bring up our balloon.
	ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button != walk.LeftButton {
			return
		}

		if err := ni.ShowCustom(
			"Walk NotifyIcon Example",
			"There are multiple ShowX methods sporting different icons."); err != nil {

			log.Fatal(err)
		}
	})

	// We put an exit action into the context menu.
	exitAction := walk.NewAction()
	if err := exitAction.SetText("E&xit"); err != nil {
		log.Fatal(err)
	}
	exitAction.Triggered().Attach(func() { walk.App().Exit(0) })
	if err := ni.ContextMenu().Actions().Add(exitAction); err != nil {
		log.Fatal(err)
	}

	// The notify icon is hidden initially, so we have to make it visible.
	if err := ni.SetVisible(true); err != nil {
		log.Fatal(err)
	}

	// Now that the icon is visible, we can bring up an info balloon.
	if err := ni.ShowInfo("Walk NotifyIcon Example", "Click the icon to show again."); err != nil {
		log.Fatal(err)
	}

	return ni
}

func main() {
	walk.AddDispatchHook(blink.DispatchBlinkMessage)

	mw, err := walk.NewMainWindow()
	if err != nil {
		log.Fatal(err)
	}

	blink.SetDebugMode(true)

	err = blink.PreInitBlink(func(f func()) {
		mw.Synchronize(f)
	})
	if err != nil {
		log.Fatal(err)
	}

	//挂载嵌入资源到虚拟文件系统
	blink.RegisterFileSystem("app", &assetfs.AssetFS{
		Asset:     ui.Asset,
		AssetDir:  ui.AssetDir,
		AssetInfo: ui.AssetInfo,
	})

	ni := notifyIcon(mw)
	if ni != nil {
		defer ni.Dispose()
	}

	//显示logo小图标
	go showLogo()

	mw.Run()
}
