package main

// so what do we want a panel to do?
// we want it to have the dimensions, the position and the contents of the panel
// it may also have a parent panel?

type Panel struct {
	String    string
	Width     int
	Height    int
	Parent    *Panel
	Alignment int
	anchorRow int
	anchorColumn int
}

func NewPanel() *Panel {
	var p Panel

	p.String = ""
	p.Width = 0
	p.Height = 0
	p.Alignment = 0

	return &p
}

func draw(p []Panel) string {
	s := ""
	for _, panel := range p {
		s+= panel.String
	}	
	return s
}



