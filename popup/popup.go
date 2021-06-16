package popup

import (
	"github.com/rivo/tview"
)

//A popup is a box you can get rid of.
type Popup struct {
	Name string

	Message string

	Box tview.Box

	CloseButton tview.Button

	x, y, w, h int

	Hidden bool
}

func NewPopup(name string) *Popup {
	var window = tview.NewBox().SetBorder(true)
	var button = tview.NewButton("Ok")
	p := &Popup{
		Name:        name,
		Message:     "BOO!",
		Box:         *window,
		CloseButton: *button,
		x:           15,
		y:           10,
		w:           15,
		h:           10,
		Hidden:      true,
	}
	p.Box.SetRect(-1, -1, 0, 0)
	return p
}

func Toggle(p *Popup) bool {
	if p.Hidden {
		p.Box.SetRect(p.x, p.y, p.w, p.h)
	} else {
		p.Box.SetRect(-1, -1, 0, 0)
	}
	p.Hidden = !p.Hidden
	return p.Hidden
}
