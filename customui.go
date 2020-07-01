package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

/// Entry widget with a handler for the enter key.
type enterEntry struct {
	widget.Entry
	OnEnter func()
}

func newEnterEntry() *enterEntry {
	entry := &enterEntry{}
	entry.ExtendBaseWidget(entry)

	return entry
}

func (e *enterEntry) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyReturn:
		e.OnEnter()
	default:
		e.Entry.KeyDown(key)
	}
}

func showTCPConnectDialog(window fyne.Window, callback func(bool)) {
	addressInput := widget.NewEntry()
	addressInput.PlaceHolder = "192.168.1.1:2345"
	content := widget.NewForm(widget.NewFormItem("Address", addressInput))

	dialog.NewCustomConfirm("TCP Connect", "Connect", "Cancel", content, callback, window).Show()
}
