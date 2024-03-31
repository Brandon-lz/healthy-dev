package main


import "github.com/lxn/walk"

func main() {
	mw,err := walk.NewMainWindow()
	if err!= nil {
		panic(err)
	}
	mw.SetTitle("Healthy Dev")
	size := mw.Size()
	size.Width = 800
	size.Height = 600
	mw.SetSize(size)
	mw.Show()
}