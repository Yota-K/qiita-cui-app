package ui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// キーバインドの設定
func Keybindings(t *widgets.Tree) {
	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "k":
			t.ScrollUp()
		case "j":
			t.ScrollDown()
		case "E":
			t.ExpandAll()
		case "C", "h":
			t.CollapseAll()
		case "<Enter>", "l":
			t.ToggleExpand()
		}

		ui.Render(t)
	}
}
