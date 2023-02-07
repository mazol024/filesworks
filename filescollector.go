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
	dumpdir := binding.NewString()
	dumpdir.Set(mypwd + pathsep + "allyourfiles" + pathsep)
	infopan := binding.NewString()
	infopan.Set("Information ...")

	al := widget.NewLabel("From source folder")
	ae := widget.NewEntryWithData(startdir)
	bl := widget.NewLabel("To destination folder")
	be := widget.NewEntryWithData(dumpdir)
	txts := widget.NewEntryWithData(infopan)
	txts.Text = "Files copied: "

	button1 := widget.NewButton("Collect files",
		func() {
			x, _ := startdir.Get()
			y, _ := dumpdir.Get()
			ll := scancopy(x, y, pathsep)
			infopan.Set(ll)
		})
	topblock := container.NewGridWithColumns(2,
		container.NewVBox(al, ae, container.NewGridWithColumns(4, button1)),
		container.NewVBox(bl, be, widget.NewLabel("")),
	)

	textblock := container.NewMax(txts)
	comby := container.NewGridWithColumns(1, topblock, textblock)
	w.SetContent(
		comby,
	)
	w.Resize(fyne.Size{Width: 640, Height: 320})
	w.ShowAndRun()

}
