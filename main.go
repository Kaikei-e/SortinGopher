package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	//"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2/layout"
	//"fyne.io/fyne/v2/dialog"

	"sortinGopher/imagesClassifier"
	"sortinGopher/permanentPath"
	"sortinGopher/unzipper"
)

func main() {
	a := app.New()
	w := a.NewWindow("SortinGopher")
	w.Resize(fyne.NewSize(750, 450))

	input  := widget.NewEntry()
	input.SetPlaceHolder("Please enter the path where the zip file is located ... ")

	content := container.NewVBox(input, widget.NewButton("Save", func(){
		permanentPath.MkPathFile(input.Text)

		log.Println("Content was : ", input.Text)
	}))

	execButton := widget.NewButton("Perform ZIP decompression and classification", func ()  {
		unzipper.SortZipFile(input.Text)
		imagesClassifier.FilesClassifier(input.Text)

	})

	w.SetContent(fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		content,
		layout.NewSpacer(),
		execButton,
	),

)


	w.ShowAndRun()
}


