package UI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func renderPopupInfo(message string, window fyne.Window) {
	dialog.ShowInformation(
		"Внимание!",
		message,
		window,
	)
}
