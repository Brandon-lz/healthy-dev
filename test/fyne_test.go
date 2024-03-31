package main


import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

func TestMain() {
	a := app.New()

	w := a.NewWindow("Health Reminder")
	w.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox(
		// 在此添加你的提醒内容，例如一个文本标签或图像等
		// 这里只是一个示例
		// fyne.NewMenuItemSeparator()("记得每隔三十分钟进行休息和伸展运动！"),
		container.NewCenter(container.NewDocTabs(container.NewTabItem("健康提醒", container.NewLabel("记得每隔三十分钟进行休息和伸展运动！"))))
	)

	w.SetContent(content)
	w.Show()

	// 每隔30分钟弹出提醒
	ticker := time.NewTicker(30 * time.Minute)
	go func() {
		for range ticker.C {
			showHealthReminder(w)
		}
	}()

	a.Run()
}

func showHealthReminder(w fyne.Window) {
	dialog.ShowInformation("健康提醒", "记得每隔三十分钟进行休息和伸展运动！", w)
}