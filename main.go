package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nathan-fiscaletti/consolesize-go"
)

// todo
// we need a panel interface
// panel and state interfaces
// create state interface
// put in 'help' to show position
// put current state into state interface
// okay so I have it, the string that gets returned by 'view'
// can be drawn into, I need to figure out how to find the correct spot in the string
// and insert a string into it, essentially 'drawing' over the top of a string.
// yeeeeas
// this is broken right now...

const (
	topLeft int = iota
	topCentre
	topRight
	left
	centre
	right
	bottomLeft
	bottomCentre
	bottomRight
)

type model struct {
	logfile      *os.File
	xPos         int
	yPos         int
	screenWidth  int
	screenHeight int
	panels []Panel
}

func main() {

	// terminal getsize, does it start at zero?! let's see!
	// these are the absolute number of cells, it doesn't start at cell zero
	var man model
	var err error

	man.logfile, err = tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}

	defer man.logfile.Close()

	width, height := consolesize.GetConsoleSize()

	man.xPos = width / 2
	man.yPos = height / 2
	man.screenWidth = width
	man.screenHeight = height

	man.panels = man.createPanels()

	p := tea.NewProgram(man, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}

}

func (m model) Init() tea.Cmd {
	return tea.HideCursor
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	width, height := consolesize.GetConsoleSize()
	m.screenWidth = width
	m.screenHeight = height

	m.panels = m.createPanels()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "up":
			m.yPos--
		case "down":
			m.yPos++
		case "left":
			m.xPos--
		case "right":
			m.xPos++
		}
	}

	m = checkBoundaries(m)

	return m, nil
}

func (m model) View() string {
	s := ""

	// for y := 0; y < m.screenHeight; y++ {
	// 	for x := 0; x < m.screenWidth; x++ {
	// 		s += drawCell(m, x, y)
	// 	}
	// }
	// s = wrap.String(s, m.screenWidth)

	// s = insertByIndex(&s, "PISSSSSSSS", 222)
	// s = m.insertByCoords(&s, "X", 10, 10)
	// s = m.insertByAbsolute(&s, fmt.Sprintf("atman is at %dx, %dy", m.xPos, m.yPos), bottomLeft)
	// s = m.insertByAbsolute(&s, "atman ats", topLeft)
	// s += setStyles().Render("Farts")
	// fmt.Fprint(m.logfile, s)S

	//fmt.Fprint(m.logfile, s)

	s = screenPainter(m.panels, m)
	return s
}

func checkBoundaries(m model) model {
	if m.xPos < 0 {
		m.xPos = 0
	} else if m.yPos < 0 {
		m.yPos = 0
	} else if m.xPos >= m.screenWidth {
		m.xPos = m.screenWidth - 1
	} else if m.yPos >= m.screenHeight {
		m.yPos = m.screenHeight - 1
	}
	return m
}

// not sure what this is, maybe I'll keep
func setStyles() lipgloss.Style {
	s := lipgloss.NewStyle().
		Bold(true).
		Background(lipgloss.Color("#505050")).
		Foreground(lipgloss.Color("#000000"))
	return s
}

// needs to be refactored
func drawCell(m model, x, y int) string {
	s := ""
	if x == m.xPos && y == m.yPos {
		s = "@"
	} else {
		s = "."
	}

	// if x == m.screenWidth-1 {
	// 	s += "\n"
	// }
	return s
}

func insertByIndex(original *string, addition string, index int) string {
	s := []rune(*original)
	character := []rune(addition)

	for i, char := range character {
		s[index+i] = char
	}

	return string(s)
}

func (m *model) insertByCoords(original *string, addition string, x, y int) string {
	index := y*m.screenWidth + y + x
	return insertByIndex(original, addition, index)
}

func (m *model) insertByAbsolute(original *string, addition string, position int) string {

	var x, y int

	switch position {
	case topLeft:
		x, y = 0, 0
	case topCentre:
		x, y = m.screenWidth/2-len(addition)/2, 0
	case topRight:
		x, y = m.screenWidth-len(addition), 0
	case left:
		x, y = 0, m.screenHeight/2
	case centre:
		x, y = (m.screenWidth-len(addition))/2, m.screenHeight/2
	case right:
		x, y = m.screenWidth-len(addition), m.screenHeight/2
	case bottomLeft:
		x, y = 0, m.screenHeight-1
	case bottomCentre:
		x, y = (m.screenWidth-len(addition))/2, m.screenHeight-1
	case bottomRight:
		x, y = m.screenWidth-len(addition), m.screenHeight-1
	}

	return m.insertByCoords(original, addition, x, y)
}

func (m *model) createPanels() []Panel{
	var panels []Panel
	var p Panel

	p.anchorRow = 0
	p.anchorColumn = 0
	p.height = 1
	p.width = m.screenWidth
	p.fillRune = '#'
	panels = append(panels, p)

	p.anchorRow = 1
	p.anchorColumn = 0
	p.height = m.screenHeight - 2
	p.width = m.screenWidth
	p.fillRune = '.'
	panels = append(panels, p)

	p.anchorRow = m.screenHeight -1
	p.anchorColumn = 0
	p.height = 1
	p.width = m.screenWidth
	p.fillRune = '#'
	panels = append(panels, p)

	return panels
}

// func (m *model) createPanels() []Panel{
// 	var panels []Panel
// 	var p Panel
// 	p.anchorRow, p.anchorColumn = 0, 0
// 	p.width = m.screenWidth/2
// 	p.height = m.screenHeight/2
// 	p.fillRune = '*'
// 	panels = append(panels, p)
// 	p.anchorRow, p.anchorColumn = 0, m.screenWidth/2 + 1
// 	p.width = m.screenWidth/2
// 	p.height = m.screenHeight/2
// 	p.fillRune = '#'
// 	panels = append(panels, p)
// 	p.anchorRow, p.anchorColumn = m.screenHeight/2 + 1, 0
// 	p.width = m.screenWidth/2
// 	p.height = m.screenHeight/2
// 	p.fillRune = 'O'
// 	panels = append(panels, p)
// 	p.anchorRow, p.anchorColumn = m.screenWidth/2+1, m.screenHeight/2+1
// 	p.width = m.screenWidth/2
// 	p.height = m.screenHeight/2
// 	p.fillRune = '&'
//  panels = append(panels, p)
// 	return panels
// }



// func (m *model) createPanels() []Panel{
// 	var panels []Panel
// 	var p Panel
// 	p.anchorRow = 0
// 	p.anchorColumn = 0
// 	p.fillRune = '#'
// 	p.width = m.screenWidth
// 	p.height = m.screenHeight
// 	panels = append(panels, p)
// 	return panels
// }