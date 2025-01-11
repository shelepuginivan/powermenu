// Package widget provides the core power menu widget.
package widget

import (
	"os/exec"
	"sync"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/powermenu/config"
)

// Widget is the core widget of powermenu that provides power management
// options and handles keyboard input.
type Widget struct {
	*gtk.Box

	mu sync.Mutex

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

	ctx, err := w.GetStyleContext()
	if err != nil {
		return nil, err
	}

	w.SetName("container")
	ctx.AddClass("container")

	return w, nil
}

// Increment increments index of the active option.
func (w *Widget) Increment() {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.unmarkActive()
	w.active = (w.active + 1) % w.total
	w.markActive()
}

// Decrement decrements index of the active option.
func (w *Widget) Decrement() {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.unmarkActive()
	w.active = (w.active + w.total - 1) % w.total
	w.markActive()
}

func (w *Widget) SetActive(i int) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.unmarkActive()
	w.active = max(min(i, w.total-1), 0)
	w.markActive()
}

// OnEscape sets callback that is triggered when Escape is pressed.
func (w *Widget) OnEscape(cb func()) {
	w.onEscape = cb
}

// markActive marks the current option as active by adding the respective CSS
// class.
func (w *Widget) markActive() {
	ctx, err := w.children[w.active].GetStyleContext()
	if err != nil {
		return
	}

	ctx.AddClass("active")
}

// unmarkActive unmarks the current option as active by removing the respective
// CSS class.
func (w *Widget) unmarkActive() {
	ctx, err := w.children[w.active].GetStyleContext()
	if err != nil {
		return
	}

	ctx.RemoveClass("active")
}

// AddOption adds option to the [Widget].
func (w *Widget) AddOption(img *gtk.Image, cmd *config.Command) {
	w.total++
	w.PackStart(img, true, true, 0)

	img.SetName("image")
	w.children = append(w.children, img)
	w.commands = append(w.commands, cmd)

	if w.total == 1 {
		w.markActive()
	}
}

// ExecuteCommand executes active command.
func (w *Widget) ExecuteCommand() {
	w.onEscape()

	cmd := w.commands[w.active]

	exec.Command(cmd.Name, cmd.Args...).Run()
}

// OnKeyPress is a callback for key press events. It is intended to be used as
// follows:
//
//	win.Connect("key-press-event", widget.OnKeyPress)
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
