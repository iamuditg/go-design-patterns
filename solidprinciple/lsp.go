package solidprinciple

import "fmt"

/*
	Liskov Substitution Principle
*/

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r Rectangle) GetWidth() int {
	return r.width
}

func (r Rectangle) SetWidth(width int) {
	r.width = width
}

func (r Rectangle) GetHeight() int {
	return r.height
}

func (r Rectangle) SetHeight(height int) {
	r.height = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Println("Expected an area of ", expectedArea, ", but got", actualArea, "\n")
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(width int) {
	s.width = width
	s.height = width
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)
}
