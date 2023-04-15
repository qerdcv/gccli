package main

import (
	"context"
	"errors"
	"github.com/qerdcv/gccli/internal/gc"
	"github.com/qerdcv/gccli/internal/ui"
)

func main() {
	// app context
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	gcClient, err := gc.New(ctx)

	u, err := ui.New(gcClient)
	if err != nil {
		panic(err)
	}

	defer u.Close()

	if err = u.MainLoop(); err != nil && !errors.Is(err, ui.ErrQuit) {
		panic(err)
	}
}
