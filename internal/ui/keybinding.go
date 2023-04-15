package ui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/qerdcv/gccli/internal/ui/view"
)

func (ui *UI) keybindings() error {
	if err := ui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, ui.quit); err != nil {
		return fmt.Errorf("gui set keybindings: %w", err)
	}

	if err := ui.SetKeybinding(view.Weekly.String(), gocui.KeyArrowRight, gocui.ModNone, ui.quit); err != nil {
		return fmt.Errorf("gui set keybindings: %w", err)
	}

	return nil
}
