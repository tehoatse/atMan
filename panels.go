package main

import "github.com/muesli/reflow/wrap"

// so what do we want a panel to do?
// we want it to have the dimensions, the position and the contents of the panel
// it may also have a parent panel?

type Panel struct {
	String    string
	width     int
	height    int
	parent    *Panel
	alignment int
	anchorRow int
	anchorColumn int
	fillRune rune
	paddingTop int
	paddingLeft int
	paddingRight int
	paddingBottom int
	paragraphs []string
}

func NewPanel() *Panel {
	var p Panel

	p.String = ""
	p.width = 0
	p.height = 0
	p.alignment = 0

	return &p
}

func draw(p []Panel) string {
	s := ""
	for _, panel := range p {
		s+= panel.String
	}	
	return s
}

// run the panel fill first?
func (p *Panel) fillPanel() string {

	s := ""
	for i := 0; i < p.height * p.width; i++ {
		s += string(p.fillRune)
	}
	s = wrap.String(s, p.width)
	return s
}


func (p *Panel) insertByCoords(original *string, addition string, x, y int) string {
	index := y*p.width + y + x
	return insertByIndex(original, addition, index)
}

func (p *Panel) insertByAbsolute(original *string, addition string, position int) string {

	var x, y int

	switch position {
	case topLeft:
		x, y = 0, 0
	case topCentre:
		x, y = p.width/2-len(addition)/2, 0
	case topRight:
		x, y = p.width-len(addition), 0
	case left:
		x, y = 0, p.height/2
	case centre:
		x, y = (p.width-len(addition))/2, p.height/2
	case right:
		x, y = p.width-len(addition), p.height/2
	case bottomLeft:
		x, y = 0, p.height-1
	case bottomCentre:
		x, y = (p.width-len(addition))/2, p.height-1
	case bottomRight:
		x, y = p.width-len(addition), p.height-1
	}

	return p.insertByCoords(original, addition, x, y)
}




