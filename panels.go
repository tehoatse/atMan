package main

// so what do we want a panel to do?
// we want it to have the dimensions, the position and the contents of the panel
// it may also have a parent panel?


type Panel interface {
	String() string
	Width() int
	Height() int
	Position() (int, int)
	Parent() Panel
	Alignment() int
}

