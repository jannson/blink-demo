package main

import "github.com/raintean/blink"

var mainWin *blink.WebView

func showMain() {
	if mainWin == nil {
		//窗口不存在,新建一个
		mainWin = blink.NewWebView(false, 1366, 920)
		mainWin.LoadURL("http://192.168.100.1:8899")
		mainWin.SetWindowIcon(uiAddress + "app.icon")
		mainWin.SetWindowTitle("易有云")
		mainWin.MoveToCenter()
		mainWin.ShowWindow()
		mainWin.ToTop()

		//当窗口被销毁的时候,变量=nil
		mainWin.On("destroy", func(_ *blink.WebView) {
			mainWin = nil
		})
	} else {
		//窗口实例存在,则提到前台
		mainWin.RestoreWindow()
		mainWin.MoveToCenter()
		mainWin.ToTop()
	}
}
