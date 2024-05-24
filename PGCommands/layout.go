package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("commands", 0, 0, int(0.2*float32(maxX)), maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Commands"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		for cmd := range commands {
			fmt.Fprintln(v, cmd)
		}

		if _, err := g.SetCurrentView("commands"); err != nil {
			log.Panicln(err)
		}
	}

	if v, err := g.SetView("console", int(0.2*float32(maxX)), 10, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Console"
	}

	if v, err := g.SetView("description", int(0.2*float32(maxX)), 0, maxX-1, 9); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Description"
	}

	updateDescription(g)

	return nil
}
