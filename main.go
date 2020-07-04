package main

import (
	"fmt"
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/ataboo/pandion-tcp-client/client"
	"github.com/ataboo/pandion-tcp-client/linebuffer"
)

var tcpClient *client.TCPClient

func drawLineBuffer(responseText *widget.TextGrid, lineBuffer *linebuffer.LineBuffer, scrollView *widget.ScrollContainer) {
	text := ""

	for i := lineBuffer.Count() - 1; i >= 0; i-- {
		line, err := lineBuffer.Get(i)
		if err != nil {
			log.Fatal("failed to get line")
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

func initTCPClient(address string) error {
	tcpClient = client.NewTCPClient(address)
	response, err := sendCommandText("testing!")
	if err != nil {
		return err
	}

	fmt.Printf("Response: %s\n", response)
	tcpClient.Close()

	return nil
}

func promptForTCPConnection(w fyne.Window, addressInput *widget.Entry) {
	showTCPConnectDialog(w, func(confirm bool) {
		if confirm {
			err := initTCPClient(addressInput.Text)
			if err != nil {
				dialog.ShowError(err, w)
			}
		} else {
			tcpClient = nil
		}
	}, addressInput)
}

func sendCommandText(command string) (response string, err error) {
	if tcpClient == nil {
		return "", fmt.Errorf("client is not initialized")
	}

	if err := tcpClient.Connect(); err != nil {
		return "", err
	}

	if err := tcpClient.WriteString(command); err != nil {
		return "", err
	}

	response, err = tcpClient.ReadString()

	return response, nil
}

func main() {
	buffer := linebuffer.NewLineBuffer(100)

	a := app.New()

	w := a.NewWindow("Pandion TCP Client")
	w.SetIcon(resourcePandionIconMintPng)

	responseText := widget.NewTextGrid()
	scrollingResponse := widget.NewVScrollContainer(responseText)

	addressInput := widget.NewEntry()
	addressInput.PlaceHolder = "127.0.0.1:3001"

	commandInput := newEnterEntry()
	commandInput.OnEnter = func() {
		buffer.Push("> " + commandInput.Text)

		response, err := sendCommandText(commandInput.Text)
		if err != nil {
			promptForTCPConnection(w, addressInput)
			dialog.ShowError(fmt.Errorf("failed to send command:\n%s", err), w)
			return
		}

		buffer.Push("< " + response)

		commandInput.SetText("")
		drawLineBuffer(responseText, buffer, scrollingResponse)
	}
	commandForm := widget.NewForm(widget.NewFormItem("Command", commandInput))
	content := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, commandForm, nil, nil), commandForm, scrollingResponse)

	w.SetContent(content)

	w.Resize(fyne.NewSize(480, 320))

	promptForTCPConnection(w, addressInput)

	w.ShowAndRun()
}
