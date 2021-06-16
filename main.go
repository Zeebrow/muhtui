package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	// "github.com/zeebrow/muhtui/muhdropdown"
	"github.com/zeebrow/muhtui/popup"
)

// Plan:
//  ______________________
// |----------------+-----|
// |                |     |
// |                |     |
// |                |     |
// +------+---------+     |
// |      |         |     |
// +------+---------+     |
// |      |         |     |
// |      |         |     |
// |      |         |     |
// |      |         |     |
// |      |         |     |
// |------+---------+-----|
// |______|_________|_____|

func main() {
	logfile := "logs/muhtui.log"

	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("Start")
	defer log.Println("Exit")

	var app = tview.NewApplication().EnableMouse(true)

	example_yaml_file := "example.yaml"

	read_yaml := func(filename string) string {
		// file, err := os.Open(filename)
		// if err != nil {
		// 	log.Fatalf("Could not open file %v", filename)
		// }
		// defer func() {
		// 	if err = file.Close(); err != nil {
		// 		log.Fatal(err)
		// 	}
		// }()

		b, _ := ioutil.ReadFile(filename)
		return string(b)
	}

	// load_yaml := func(filename string) (*mydata, error) {
	// 	buf, err := ioutil.ReadFile(filename)
	// 	if err != nil {
	// 		log.Fatalf("Could not open yaml file %v. %v", filename, err)
	// 		return nil, err
	// 	}
	// 	c := &mydata{}
	// 	err = yaml.Unmarshal(buf, c)
	// 	if err != nil {
	// 		log.Fatalf("Error reading yaml file %v: %v", filename, err)
	// 		return nil, err
	// 	}
	// 	return c, nil
	// }

	muhdropdown := tview.NewDropDown().SetLabel("click to select:").
		SetOptions([]string{"muhfirstoption", example_yaml_file, "mop3", "mop4", "mop5", "mop6"}, func(o string, i int) {
			log.Printf("muhbox selection now (%v)%v ", i, o)
			d := read_yaml(o)
			log.Printf("%s", d)
			// data, err := load_yaml(o)
			// if err != nil {
			// 	log.Printf("error retrieving yaml data from option (%v)", o)
			// } else {
			// 	log.Printf("%v", data)
			// }
		}).SetCurrentOption(0)

	logframe := tview.NewFlex().SetDirection(tview.FlexRow).
		//AddItem(tview.NewBox().SetBorder(true).SetTitle("muhactivities").SetBackgroundColor(tcell.ColorDarkGoldenrod), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("muhlogs").SetBackgroundColor(tcell.ColorDarkGoldenrod), 0, 1, false)

	editor := tview.NewFlex().SetDirection(tview.FlexColumn).
		//AddItem(tview.NewBox().SetBorder(true).SetTitle("mmmmmmmm").SetBackgroundColor(tcell.ColorDarkRed), 0, 1, false).
		AddItem(muhdropdown, 0, 5, false).
		//AddItem(tview.NewBox().SetBorder(true).SetTitle("uuuuuuuu").SetBackgroundColor(tcell.ColorDarkKhaki), 0, 2, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("preview").SetBackgroundColor(tcell.ColorDarkOrchid), 0, 3, false)

	muhtopbox := tview.NewBox().
		SetBorder(true).
		SetTitle("muhtopbox").
		SetBackgroundColor(tcell.ColorDarkGray)

	muhbasebox := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(muhtopbox, 0, 2, false).
		AddItem(editor, 0, 5, false).
		AddItem(logframe, 0, 3, false)

	//SetBorderAttributes(tcell.AttrBold).
	//SetTitleColor(tcell.ColorDarkOrange.TrueColor())

	// muhbox := tview.NewBox().
	// 	SetBorder(true).
	// 	SetTitle("muhbox").
	// 	SetBackgroundColor(tcell.ColorGray)

	// muhRightpane := tview.NewBox().

	mlayout := tview.NewFlex().
		//SetDirection(tview.FlexColumn).
		AddItem(muhbasebox, 0, 3, false)
		//AddItem(muhbox, 0, 1, false)

	var muhpopup = popup.NewPopup("muhPOPUP")

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		log.Printf("Key event: %v\t %v", event.Key(), event.Name())
		if event.Key() == tcell.KeyCtrlN {
			//app.SetFocus(muhbox)
			return nil
		} else if event.Key() == tcell.KeyCtrlP {
			app.SetFocus(muhbasebox)
			return nil
		} else if event.Key() == tcell.KeyEnter {
			popup.Toggle(muhpopup)
		}
		return event
	})

	if err := app.SetRoot(mlayout, true).Run(); err != nil {
		panic(err)
	}

}
