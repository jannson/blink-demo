package main

import (
	"log"

	ui "blink-demo/ui/bin"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/jannson/walk"
	"github.com/raintean/blink"
)

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

	//显示logo小图标
	go showLogo()

	mw.Run()
}
