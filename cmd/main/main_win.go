package main

import (
	"github.com/jannson/miniblink"
)

var mainWin *miniblink.WebView

func showMain() {
	if mainWin == nil {
		//窗口不存在,新建一个
		mainWin = miniblink.NewWebView(false, true, 1366, 920)
		mainWin.LoadURL("http://127.0.0.1:8896")
		mainWin.SetWindowIcon(uiAddress + "app.icon")
		mainWin.SetWindowTitle("易有云")
		mainWin.MoveToCenter()
		mainWin.ShowWindow()
		mainWin.ToTop()

		//当窗口被销毁的时候,变量=nil
		mainWin.On("destroy", func(_ *miniblink.WebView) {
			mainWin = nil
		})
	} else {
		//窗口实例存在,则提到前台
		mainWin.RestoreWindow()
		mainWin.MoveToCenter()
		mainWin.ToTop()
	}
}
