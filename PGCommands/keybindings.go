package main

import "github.com/jroimartin/gocui"

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("commands", gocui.KeyEnter, gocui.ModNone, executeCommand); err != nil {
		return err
	}

	if err := g.SetKeybinding("commands", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}

	if err := g.SetKeybinding("commands", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	return nil
}
