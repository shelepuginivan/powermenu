package main

import (
	"sort"

	"github.com/dlasky/gotk3-layershell/layershell"
	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/powermenu/config"
	"github.com/shelepuginivan/powermenu/icons"
	"github.com/shelepuginivan/powermenu/must"
	"github.com/shelepuginivan/powermenu/styles"
	"github.com/shelepuginivan/powermenu/widget"
)

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	win := must.Must(gtk.WindowNew(gtk.WINDOW_TOPLEVEL))
	cfg := config.Resolve()

	// Properly handle transparency.
	win.SetAppPaintable(true)
	screen := win.GetScreen()
	visual, err := screen.GetRGBAVisual()
	if err == nil {
		win.SetVisual(visual)
	}

	// Initialize layershell for window.
	layershell.InitForWindow(win)
	layershell.SetNamespace(win, "powermenu")

	layershell.SetLayer(win, layershell.LAYER_SHELL_LAYER_OVERLAY)

	// Set anchors based on the resolved configuration.
	layershell.SetAnchor(win, layershell.LAYER_SHELL_EDGE_TOP, cfg.Anchors.Top)
	layershell.SetAnchor(win, layershell.LAYER_SHELL_EDGE_LEFT, cfg.Anchors.Left)
	layershell.SetAnchor(win, layershell.LAYER_SHELL_EDGE_RIGHT, cfg.Anchors.Right)
	layershell.SetAnchor(win, layershell.LAYER_SHELL_EDGE_BOTTOM, cfg.Anchors.Bottom)

	// Set margins based on the resolved configuration.
	layershell.SetMargin(win, layershell.LAYER_SHELL_EDGE_TOP, cfg.Margins.Top)
	layershell.SetMargin(win, layershell.LAYER_SHELL_EDGE_LEFT, cfg.Margins.Left)
	layershell.SetMargin(win, layershell.LAYER_SHELL_EDGE_RIGHT, cfg.Margins.Right)
	layershell.SetMargin(win, layershell.LAYER_SHELL_EDGE_BOTTOM, cfg.Margins.Bottom)

	// Load application CSS.
	styles.Load()

	// Create a new label widget to show in the window.
	widget := must.Must(widget.New())
	widget.SetSpacing(cfg.Layout.Spacing)
	widget.SetOrientationFromString(cfg.Layout.Orientation)
	widget.OnEscape(gtk.MainQuit)

	// Configure window.
	win.SetName("powermenu")
	win.SetTitle("powermenu")
	win.Connect("destroy", gtk.MainQuit)
	win.Connect("key-press-event", widget.OnKeyPress)
	win.Add(widget)

	options := []*config.Option{
		cfg.Hibernate,
		cfg.Lock,
		cfg.Logout,
		cfg.PowerOff,
		cfg.Reboot,
		cfg.Suspend,
	}

	sort.Slice(options, func(i, j int) bool {
		return options[i].Order < options[j].Order
	})

	for _, option := range options {
		if option.Enabled {
			icons.ExistsOrWrite(option.Icon, option.GetDefaultIcon())
			img := must.Must(gtk.ImageNewFromFile(option.Icon))
			widget.AddOption(img, option.Command)
		}
	}

	// Allow keyboard input.
	layershell.SetKeyboardMode(win, layershell.LAYER_SHELL_KEYBOARD_MODE_ON_DEMAND)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.
	// This blocks until gtk.MainQuit() is run.
	gtk.Main()
}
