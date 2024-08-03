package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&myTheme{resourceNotoSansHansRegularTtf}) // 设置自定义主题
   
	w := a.NewWindow("Custom Theme with Icon Override")
	w.SetContent(container.NewVBox(
	widget.NewLabel("Hello, Fyne!"),
	widget.NewButtonWithIcon("Home", theme.HomeIcon(), func(){}),
))
   
	w.ShowAndRun()
}