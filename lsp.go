package main

import "fmt"

// Liskov Substitution Principle

type Sized interface {
    GetWidth() int
    SetWidth(int)
    GetHeight() int
    SetHeight(int)
}

type Rectangle struct {
    width, height int
}

func (r *Rectangle) GetWidth() int {
   return r.width
}

func (r *Rectangle) SetWidth(i int) {
   r.width = i
}

func (r *Rectangle) GetHeight() int {
    return r.height
}

func (r *Rectangle) SetHeight(i int) {
    r.height = i
}

// Square breaks UseIt function and therefore violates LSP
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

func (s *Square) SetHeight(height int) {
    s.height = height
    s.width = height
}

// LSP states that interface functions should continue to work regardless of what has been done to the implementations
// objects that inherit should not break assumptions made before hand

func UseIt(sized Sized) {
    width := sized.GetWidth()
    sized.SetHeight(10)
    expectedArea := 10 * width
    actualArea := sized.GetWidth() * sized.GetHeight()
    fmt.Printf("Expected an area of %d, but got %d\n", expectedArea, actualArea )
}

func  lsp() {
    r := &Rectangle{2, 3}
    UseIt(r)
}