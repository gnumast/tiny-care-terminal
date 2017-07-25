package main

import ui "github.com/gizak/termui"

func main() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	today := dailyCommits()
	week := dailyCommits()
	weather := makeWeather()
	tweet1 := makeWeather()
	tweet2 := makeWeather()

	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(6, 0, today), ui.NewCol(6, 0, weather, tweet1, tweet2)),
		ui.NewRow(ui.NewCol(6, 0, week)),
	)

	// calculate layout
	ui.Body.Align()

	ui.Render(ui.Body)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Handle("/timer/1s", func(e ui.Event) {
		t := e.Data.(ui.EvtTimer)
		i := t.Count
		if i > 103 {
			ui.StopLoop()
			return
		}

		ui.Render(ui.Body)
	})

	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Clear()
		ui.Render(ui.Body)
	})

	ui.Loop()
}

func dailyCommits() (ls *ui.List) {
	strs := []string{
		"[/Users/alex/go/src/github.com/gnumast/tiny-care-terminal](fg-red)",
		"[e9da701b](fg-green) - Fixed a thing",
		"[d825f5e0](fg-green) - Added a feature",
		"[2a84fe96](fg-green) - Initial commit"}

	ls = ui.NewList()
	ls.Items = strs
	ls.ItemFgColor = ui.ColorWhite
	ls.BorderLabel = "Today"
	ls.BorderLabelFg = ui.ColorWhite
	ls.BorderFg = ui.ColorBlue
	ls.Height = 21

	return ls
}

func makeWeather() (w *ui.Par) {
	w = ui.NewPar("Simple colored text\nwith label. It [can be](fg-red) multilined with \\n or something!")
	w.Height = 7
	w.Y = 4
	w.BorderLabel = "☀️ Weather"
	w.BorderFg = ui.ColorBlue
	w.BorderLabelFg = ui.ColorWhite

	return
}