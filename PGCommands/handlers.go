package main

import (
	"fmt"
	"os/exec"

	"github.com/jroimartin/gocui"
)

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		lines := v.BufferLines()
		newCy := cy + 1
		if newCy >= len(lines)-1 {
			v.SetCursor(cx, 0)
		} else {
			v.SetCursor(cx, newCy)
		}
		updateDescription(g)
	}

	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		lines := v.BufferLines()
		newCy := cy - 1
		if newCy < 0 {
			lastIndex := len(lines) - 2
			v.SetCursor(cx, lastIndex)
		} else {
			v.SetCursor(cx, newCy)
		}
		updateDescription(g)
	}

	return nil
}

func updateDescription(g *gocui.Gui) error {
	commandsView, err := g.View("commands")
	if err != nil {
		return err
	}
	descriptionView, err := g.View("description")
	if err != nil {
		return err
	}

	_, cy := commandsView.Cursor()
	line, err := commandsView.Line(cy)
	if err != nil {
		line = ""
	}

	descriptionView.Clear()
	if cmd, ok := commands[line]; ok {
		fmt.Fprintln(descriptionView, cmd.Description)
		fmt.Fprintln(descriptionView)
		fmt.Fprintln(descriptionView, "Command: "+cmd.Cmd)
	} else {
		fmt.Fprintln(descriptionView, "No description available")
	}

	return nil
}

func executeCommand(g *gocui.Gui, v *gocui.View) error {
	commandsView, err := g.View("commands")
	if err != nil {
		return err
	}
	consoleView, err := g.View("console")
	if err != nil {
		return err
	}

	_, cy := commandsView.Cursor()
	line, err := commandsView.Line(cy)
	if err != nil {
		return err
	}

	command, ok := commands[line]
	if !ok {
		fmt.Fprintln(consoleView, "Unknown command")
		return nil
	}

	consoleView.Clear()
	cmd := exec.Command("sh", "-c", command.Cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(consoleView, "Error:", err)
	}
	fmt.Fprintln(consoleView, string(out))

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
