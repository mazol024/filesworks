package main

import (
	"os"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var pathsep string

func main() {
	switch runtime.GOOS {
	case "windows":
		pathsep = "\\"
	default:
		pathsep = "/"

	}

	myApp := app.New()
	w := myApp.NewWindow("Files collector")
	startdir := binding.NewString()
	mypwd, _ := os.Getwd()
	startdir.Set(mypwd)
	// startdir.Set("sorce dir ")
	dumpdir := binding.NewString()
	dumpdir.Set(mypwd + pathsep + "allyourfiles" + pathsep)
	infopan := binding.NewString()
	infopan.Set("Information ...")

	al := widget.NewLabel("From source folder")
	// al := widget.NewLabelWithData(startdir)
	ae := widget.NewEntryWithData(startdir)
	bl := widget.NewLabel("To destination folder")
	// bl := widget.NewLabelWithData(dumpdir)
	be := widget.NewEntryWithData(dumpdir)
	txts := widget.NewEntryWithData(infopan)

	button1 := widget.NewButton("Run Collect",
		func() {
			// a, _ := startdir.Get()
			// output, _ := os.ReadDir(a)
			// longlist := ""
			// for _, i := range output {
			// 	longlist = longlist + "\n" + i.Name()
			// }
			// infopan.Set(" Button pressed \n" + longlist)
			x, _ := startdir.Get()
			y, _ := dumpdir.Get()
			ll := scancopy(x, y, pathsep)
			infopan.Set(ll)
		})
	button1.Resize(fyne.NewSize(50, 20))
	bgridh := container.NewGridWithColumns(4, button1)
	bgridv := container.NewGridWithRows(4, bgridh)

	topblock := container.NewGridWithColumns(2,
		container.NewVBox(al, ae),
		container.NewVBox(bl, be),
	)
	bottomblock := container.NewMax(txts)
	comby := container.NewGridWithColumns(1, topblock, bottomblock, bgridv)
	w.SetContent(
		comby,
	)
	w.Resize(fyne.Size{Width: 640, Height: 480})
	w.ShowAndRun()

}
