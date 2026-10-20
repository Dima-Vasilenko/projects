package book

import (
	"fmt"
)

type Author struct {
	LastName  string
	FirstName string
}

type Book struct {
	Title       string
	Description string
	Pages       int
	Id          int
	WasTaken    bool
	Author
}

func (b *Book) String() string {
	return fmt.Sprintf("%v, %vpages. %v; ID: %v", b.Title, b.Pages, b.LastName, b.Id)
}

func (b *Book) HasTaken() bool {
	return b.WasTaken
}

func (b *Book) SetTaken() {
	b.WasTaken = true
}

func (b *Book) SetFree() {
	b.WasTaken = false
}
