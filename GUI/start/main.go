package main

import (
	"bufio"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"os"
)

func main() {
	fileToBeUploaded := "download.png"
	file, err := os.Open(fileToBeUploaded)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	// read file into bytes
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes) // <--------------- here!

	file.Close()

	iconress := fyne.NewStaticResource("icon", bytes)

	App.SetIcon(iconress)

	w1.SetContent(widget.NewVBox(
		//lbl1,
		//cb1,
		//btn1,
		tb1,
		btn2,
		lbl2,
	))

	w1.ShowAndRun()

}

var App = app.New()
var w1 = App.NewWindow("Hello")
var w2 = App.NewWindow("message")

//labels:
var lbl1 = widget.NewLabel("first time App run")
var lbl2 = widget.NewLabel("")

//checkboxes:
var cb1 = widget.NewCheck("True or false", func(bol bool) {
	lbl1.SetText(fmt.Sprintf("%v", bol))
})

//buttons:
var btn1 = widget.NewButton("Quit", func() {
	App.Quit()
})
var btn2 = widget.NewButton("show", btn2Action)

//textBoxes:
var tb1 = widget.NewEntry()

func btn2Action() {
	w2.SetContent(widget.NewHBox(
		widget.NewLabel(tb1.Text),
	))

	w2.ShowAndRun()
}
