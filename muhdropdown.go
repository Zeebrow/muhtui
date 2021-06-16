package main

import (
	"log"

	"github.com/rivo/tview"
)

func Muhdropdown() *tview.DropDown {
	Muhdropdown := tview.NewDropDown().SetLabel("muhdropdown:").
		SetOptions([]string{"muhfirstoption", "muhnotheroption", "mop3", "mop4", "mop5", "mop6"}, func(o string, i int) {
			log.Printf("muhbox selection now (%v)%v ", i, o)
		})
	return Muhdropdown
}
