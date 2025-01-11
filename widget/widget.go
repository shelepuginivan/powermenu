// Package widget provides the core power menu widget.
package widget

import (
	"os/exec"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/powermenu/config"
)

type Widget struct {
	*gtk.Box

	// Index of the active option.
	active int

	// Total number of options.
	total int

	onEscape func()
	children []*gtk.Image
	commands []*config.Command
}

// New returns new power menu widget.
func New() (*Widget, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 12)
	if err != nil {
		return nil, err
	}

	w := &Widget{
		Box:    box,
		active: 0,
		total:  0,
	}

	return w, nil
}

// Increment increments index of the active option.
func (w *Widget) Increment() {
	w.active = (w.active + 1) % w.total
}

// Decrement decrements index of the active option.
func (w *Widget) Decrement() {
	w.active = (w.active + w.total - 1) % w.total
}

func (w *Widget) OnEscape(cb func()) {
	w.onEscape = cb
}

func (w *Widget) SetActive(i int) {
	w.active = max(min(i, w.total-1), 0)
}

func (w *Widget) AddOption(img *gtk.Image, cmd *config.Command) {
	w.total++
	w.PackEnd(img, true, true, 0)

	w.children = append(w.children, img)
	w.commands = append(w.commands, cmd)
}

func (w *Widget) ExecuteCommand() {
	cmd := w.commands[w.active]

	exec.Command(cmd.Name, cmd.Args...).Run()
}

func (w *Widget) OnKeyPress(_ interface{}, ev *gdk.Event) {
	event := &gdk.EventKey{
		Event: ev,
	}

	switch event.KeyVal() {
	case gdk.KEY_Escape:
		w.onEscape()
	case gdk.KEY_Left, gdk.KEY_h:
		w.Decrement()
	case gdk.KEY_Right, gdk.KEY_l:
		w.Increment()
	case gdk.KEY_1:
		w.SetActive(0)
	case gdk.KEY_2:
		w.SetActive(1)
	case gdk.KEY_3:
		w.SetActive(2)
	case gdk.KEY_4:
		w.SetActive(3)
	case gdk.KEY_5:
		w.SetActive(4)
	case gdk.KEY_6:
		w.SetActive(5)
	case gdk.KEY_Return:
		w.ExecuteCommand()
	}
}
