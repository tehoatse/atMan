package main

import (
	"github.com/muesli/reflow/wrap"
	"github.com/nathan-fiscaletti/consolesize-go"
)

func screenPainter(panels []Panel) string {
	var s string
	width, height := consolesize.GetConsoleSize()
	length := width * height

	for i := 0; i < length; i ++ {
		s += paintCell(panels, i, width)		
	}
	
	return wrap.String(s, width)
}

func paintCell(panels []Panel, i, width int) string {
	var s string
	
	if width == 0 {
		return s
	}

	for _, panel := range panels {
		col := i % width
		row := i / width

		if panel.anchorColumn >= col 
		&& panel.anchorRow >= row 
		&& col < panel.anchorColumn + panel.width
		&& row < panel.anchorRow + panel.height
	}
	return s
}