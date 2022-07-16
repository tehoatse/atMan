package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

type model struct {
	logfile *os.File
	xPos    int
	yPos    int
	screenX int
	screenY int
}

func main() {

	width, height, err := term.GetSize(0)
	if err != nil {
		return
	}

	var man model
	man.screenX = width
	man.screenY = height
	man.xPos = width / 2
	man.yPos = height / 2

	man.logfile, err = tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer man.logfile.Close()

	p := tea.NewProgram(man, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}

}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	width, height, err := term.GetSize(0)
	if err != nil {
		return m, nil
	}

	m.screenX = width
	m.screenY = height

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

	return m, nil
}

func (m model) View() string {
	s := ""

	for y := 0; y < m.screenY; y++ {
		for x := 0; x < m.screenX; x++ {
			s += draw(m, x, y)
		}
	}

	fmt.Fprint(m.logfile, s)
	return s
}

func draw(m model, x, y int) string {
	s := ""
	if x == m.xPos && y == m.yPos {
		s = "@"
	} else {
		s = "."
	}

	if x == m.screenX-1 {
		s += "\n"
	}
	return s
}
