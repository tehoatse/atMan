package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

type model struct {
	xPos int
	yPos int
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
	man.xPos = width/2
	man.yPos = height/2

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
		switch msg.String(){
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

	for i := 0; i <= m.yPos; i++ {
		s += "\n"
	}
	for i := 0; i <= m.xPos; i++ {
		s += " "
	}	

	s += "@"
	return s
}