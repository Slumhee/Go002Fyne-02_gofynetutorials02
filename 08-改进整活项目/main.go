package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("01")
	customFont := fyne.NewStaticResource("NotoSansHans.ttf", loadFont("NotoSansHans-Regular.ttf"))
	a.Settings().SetTheme(&myTheme{font: customFont})

	w := a.NewWindow("高端检测器")

	btn := widget.NewButton("检测是否开机", func(){
	showAdDialog(w)
	})

	w.SetContent(container.NewVBox(btn))
	w.Resize(fyne.Size{Width: 514, Height: 114})
	w.CenterOnScreen()
	w.ShowAndRun()
}

func  showAdDialog(w fyne.Window){
	entry := widget.NewPasswordEntry()
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "请输入密码继续检测", Widget: entry},
		},
	OnSubmit: func() {
		if entry.Text == "114514"{
			showLoadingDialog(w)
		}else{
			dialog.ShowInformation("错误", "密码错误, 请关注B站InkkaPlum频道三连得到正确密码", w)
		}
	},
	}
	adContent := container.NewVBox(
		widget.NewLabel("请关注B站InkkaPlum频道三连得到正确密码并输入以继续检测"),
		form,
	)
	adDialog := dialog.NewCustom("注意!", "取消", adContent, w)
	adDialog.Show()
}

func showLoadingDialog(w fyne.Window){
	progress := widget.NewProgressBarInfinite()
		progressContainer := container.NewVBox(progress)

		loadingDialog := dialog.NewCustom("正在检测...", "取消", progressContainer, w)
		loadingDialog.Show()

	go func(){	time.Sleep(20 * time.Second)
		loadingDialog.Hide()
		dialog.ShowInformation("结果", "电脑是开机的", w)}()
	}
