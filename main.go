package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type pandionClient struct {
	window fyne.Window
	input  *widget.Entry
	output *widget.Label
}

// func (c *pandionClient) buildToolbar() *widget.Toolbar {
// 	return widget.NewToolbar(widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
// 		c.entry.SetText("")
// 	}),
// 		widget.NewToolbarSeparator(),
// 		widget.NewToolbarAction(theme.ContentCutIcon(), func() {
// 			e.cut()
// 		}),
// 		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
// 			e.copy()
// 		}),
// 		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {
// 			e.paste()
// 		}))
// }

func Show(app fyne.App) {
	window := app.NewWindow("Pandion TCP Client")
	// window.SetIcon(icon.TextEditorBitmap)

	input := widget.NewEntry()
	output := widget.NewLabel("")

	client := &pandionClient{
		window: window,
		input:  input,
		output: output,
	}

	// toolbar := client.buildToolbar()
	inputRow := widget.NewHBox(layout.NewSpacer(),
		widget.NewLabel("Response:"), output,
		widget.NewButton("Send", func() {

		}))
	// content := fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, status, nil, nil),
	// 	toolbar, status, widget.NewScrollContainer(entry))

	content := fyne.NewContainerWithLayout(layout.NewBorderLayout(client.output, inputRow, nil, nil), client.output, inputRow)

	window.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Exit", func() {
				client.output.SetText("Tadaa!")
			}),
		),
		fyne.NewMenu("Edit", fyne.NewMenuItem("Blah", func() {
			client.output.SetText("Blah")
		})),
		// fyne.NewMenu("Edit",
		// 	fyne.NewMenuItem("Cut", editor.cut),
		// 	fyne.NewMenuItem("Copy", editor.copy),
		// 	fyne.NewMenuItem("Paste", editor.paste),
		// ),
	))

	window.SetContent(content)
	window.Resize(fyne.NewSize(480, 320))
	window.Show()
}

func main() {
	app := app.New()
	Show(app)

	app.Run()
}
