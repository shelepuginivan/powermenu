package main

import (
	"log"

	"github.com/dlasky/gotk3-layershell/layershell"
	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/powermenu/config"
	"github.com/shelepuginivan/powermenu/styles"
)

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	config := config.Resolve()

	// Initialize layershell for window.
	layershell.InitForWindow(win)
	layershell.SetNamespace(win, "powermenu")

	// Set anchors based on the resolved configuration.
	layershell.SetAnchor(win, layershell.LAYER_SHELL_EDGE_TOP, config.Anchors.Top)
	layershell.SetAnchor(win, layershell.LAYER_SHELL_EDGE_LEFT, config.Anchors.Left)
	layershell.SetAnchor(win, layershell.LAYER_SHELL_EDGE_RIGHT, config.Anchors.Right)
	layershell.SetAnchor(win, layershell.LAYER_SHELL_EDGE_BOTTOM, config.Anchors.Bottom)

	// Set margins based on the resolved configuration.
	layershell.SetMargin(win, layershell.LAYER_SHELL_EDGE_TOP, config.Margins.Top)
	layershell.SetMargin(win, layershell.LAYER_SHELL_EDGE_LEFT, config.Margins.Left)
	layershell.SetMargin(win, layershell.LAYER_SHELL_EDGE_RIGHT, config.Margins.Right)
	layershell.SetMargin(win, layershell.LAYER_SHELL_EDGE_BOTTOM, config.Margins.Bottom)

	// Configure window.
	win.SetTitle("powermenu")
	win.Connect("destroy", gtk.MainQuit)

	// Load application CSS.
	styles.Load()

	// Create a new label widget to show in the window.
	l, err := gtk.LabelNew("Hello, gotk3!")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	// Add the label to the window.
	win.Add(l)

	// Set the default window size.
	win.SetDefaultSize(800, 30)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.
	// This blocks until gtk.MainQuit() is run.
	gtk.Main()
}
