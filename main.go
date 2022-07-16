package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

// todo
// code in boundaries
// create state interface
// put in 'help' to show position
// put current state into state interface

type model struct {
	logfile *os.File
	xPos    int
	yPos    int
}

func main() {

	width, height, err := term.GetSize(0)
	if err != nil {
		return
	}

	var man model
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

	for i := 0; i < m.yPos; i++ {
		s += "\n"
	}
	for i := 0; i < m.xPos; i++ {
		s += " "
	}

	s += "@"

	//	fmt.Fprint(m.logfile, s)
	return s
}

func checkBoundaries(m model) model {
	width, height, err := term.GetSize(0)

	if err != nil {
		return m
	} else if m.xPos < 0 {
		m.xPos = 0
	} else if m.yPos < 0 {
		m.yPos = 0
	} else if m.xPos >= width {
		m.xPos = width - 1
	} else if m.yPos >= height {
		m.yPos = height - 1
	}
	return m

}

func setStyles() lipgloss.Style {
	s := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA"))
	return s
}
