package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	//"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2/layout"
	//"fyne.io/fyne/v2/dialog"

	"sortinGopher/permanentPath"
)

func main() {
	a := app.New()
	w := a.NewWindow("SortinGopher")
	w.Resize(fyne.NewSize(900, 600))

	input  := widget.NewEntry()
	input.SetPlaceHolder("Please enter the path where the zip file is located ... ")

	content := container.NewVBox(input, widget.NewButton("Save", func(){
		permanentPath.MkPathFile(input.Text)

		log.Println("Content was : ", input.Text)
	}))



	w.SetContent(content)
	w.ShowAndRun()
}


