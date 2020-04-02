package main

import (
	"log"

	"github.com/jannson/miniblink"
	"github.com/lxn/win"
)

func showLogo() {
	//启动小图标
	//获取屏幕大小
	logoWin := miniblink.NewWebView(true, true,
		117, 133,
		int(win.GetSystemMetrics(win.SM_CXSCREEN)/5*4),
		int(win.GetSystemMetrics(win.SM_CYSCREEN)/5))
	log.Println("create logo win ok")

	logoWin.LoadURL(uiAddress + "index.html#/logo")
	logoWin.HideDockIcon()
	logoWin.ShowWindow()

	//注入打开主窗口函数
	logoWin.Inject("OpenMainWin", showMain)
}
