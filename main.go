package main

import (
	"log"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/ataboo/pandion-tcp-client/client"
	"github.com/ataboo/pandion-tcp-client/linebuffer"
)

func drawLineBuffer(responseText *widget.TextGrid, lineBuffer *linebuffer.LineBuffer, scrollView *widget.ScrollContainer) {
	text := ""

	for i := lineBuffer.Count() - 1; i >= 0; i-- {
		line, err := lineBuffer.Get(i)
		if err != nil {
			log.Fatal("Failed to get line")
		}

		if i < lineBuffer.Count()-1 {
			text += "\n"
		}

		text += line
	}

	responseText.SetText(text)
	scrollView.Refresh()
	scrollView.Scrolled(&fyne.ScrollEvent{
		DeltaY: -100,
	})
}

func main() {
	var tcpClient *client.TcpClient
	buffer := linebuffer.NewLineBuffer(100)

	a := app.New()

	w := a.NewWindow("Pandion TCP Client")
	w.SetIcon(resourcePandionIconMintPng)

	responseText := widget.NewTextGrid()
	scrollingResponse := widget.NewVScrollContainer(responseText)

	commandInput := newEnterEntry()
	commandInput.OnEnter = func() {
		tcpClient.Send(commandInput.Text)
		buffer.Push(commandInput.Text)
		commandInput.SetText("")

		drawLineBuffer(responseText, buffer, scrollingResponse)
	}
	commandForm := widget.NewForm(widget.NewFormItem("Command", commandInput))
	content := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, commandForm, nil, nil), commandForm, scrollingResponse)

	w.SetContent(content)

	addressInput := widget.NewEntry()
	addressInput.PlaceHolder = "127.0.0.1:3001"

	showTCPConnectDialog(w, func(confirm bool) {
		if confirm {
			tcpClient = client.NewTcpClient()
			err := tcpClient.Connect(addressInput.Text)
			if err != nil {
				println("Connect failed")
				os.Exit(1)
			}

			tcpClient.StartReadPump(func(line string) {
				buffer.Push(line)
				drawLineBuffer(responseText, buffer, scrollingResponse)
			})

			// fmt.Println("Yes!")
		} else {
			// fmt.Println("No!")
		}
	}, addressInput)

	w.Resize(fyne.NewSize(480, 320))

	w.ShowAndRun()
}
