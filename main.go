package main

import (
	"fmt"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	os.Setenv("FYNE_SCALE", "1.5")

	a := app.New()

	w := a.NewWindow("Hello")
	w.SetIcon(resourcePandionIconMintPng)

	commandInput := newEnterEntry()
	commandInput.OnEnter = func() {
		fmt.Println("Got: " + commandInput.Text)
		commandInput.SetText("")
	}

	responseText := widget.NewTextGrid()
	responseText.SetText("First Row?\nsecond row?")

	commandForm := widget.NewForm(widget.NewFormItem("Command", commandInput))

	if deskCanvas, ok := w.Canvas().(desktop.Canvas); ok {
		deskCanvas.SetOnKeyDown(func(ev *fyne.KeyEvent) {
			fmt.Println("KeyDown: " + string(ev.Name))
		})
		deskCanvas.SetOnKeyUp(func(ev *fyne.KeyEvent) {
			fmt.Println("KeyUp  : " + string(ev.Name))
		})
	}

	content := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, commandForm, nil, nil), commandForm, responseText)

	w.SetContent(content)
	w.Resize(fyne.NewSize(480, 320))

	showTCPConnectDialog(w, func(confirm bool) {
		if confirm {
			fmt.Println("Yes!")
		} else {
			fmt.Println("No!")
		}
	})

	w.ShowAndRun()
}
