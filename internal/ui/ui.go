package ui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/qerdcv/gccli/internal/gc"
	"github.com/qerdcv/gccli/internal/ui/view"
)

type UI struct {
	*gocui.Gui

	gcClient *gc.CalendarClient
}

func New(gcClient *gc.CalendarClient) (*UI, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return nil, fmt.Errorf("gocui new gui: %w", err)
	}

	ui := &UI{
		g,
		gcClient,
	}

	g.SetManagerFunc(ui.layout)

	if err = ui.keybindings(); err != nil {
		return nil, fmt.Errorf("ui set keybindings: %w", err)
	}

	return ui, nil
}

func (ui *UI) layout(g *gocui.Gui) error {
	events, err := ui.gcClient.GetEvents()
	if err != nil {
		return fmt.Errorf("google calendar get events: %w", err)
	}

	return view.NewWeeklyView().Layout(g, events)
}

func (ui *UI) quit(g *gocui.Gui, v *gocui.View) error {
	return ErrQuit
}
