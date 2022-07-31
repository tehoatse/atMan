package main

import (
	"fmt"

	"github.com/muesli/reflow/wrap"
	"github.com/nathan-fiscaletti/consolesize-go"
)

func screenPainter(panels []Panel, m model) string {
	var s string
	width, height := consolesize.GetConsoleSize()
	length := width * height

	

	for i := 0; i < length; i++ {
		s += paintCell(panels, i, width)		
	}

	s = wrap.String(s, width)
	
	fmt.Fprint(m.logfile, s, "|")

	return s
}

func paintCell(panels []Panel, i, width int) string {
	var s string
	
	if width <= 0 {
		return s
	}

	for _, panel := range panels {
		cellCol := i % width
		cellRow := i / width

		if cellCol >= panel.anchorColumn &&
		cellRow >= panel.anchorRow &&
		cellRow < (panel.anchorRow + panel.height) &&
		cellCol < (panel.anchorColumn + panel.width) {
			s = string(panel.fillRune)
		}
	}
	
	return s
}