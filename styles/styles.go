// Package styles provides CSS styles for powermenu.
package styles

import (
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

var stylesPath = filepath.Join(xdg.ConfigHome, "powermenu", "style.css")

// Load loads stylesheet and applies it for the application.
func Load() {
	// Load CSS stylesheet from the configuration directory.
	provider, _ := gtk.CssProviderNew()
	provider.LoadFromPath(stylesPath)

	// Apply styles for the application.
	screen, _ := gdk.ScreenGetDefault()
	gtk.AddProviderForScreen(screen, provider, 1)
}
