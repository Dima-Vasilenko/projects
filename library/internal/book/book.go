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
	ID          int
	Author
}

func NewBook(title, descr, authorLast, authorFirst string, pages int) Book {
	book := Book{
		Title:       title,
		Description: descr,
		Pages:       pages,
		Author: Author{
			LastName:  authorLast,
			FirstName: authorFirst,
		},
	}
	return book
}

func (b *Book) String() string {
	return fmt.Sprintf("Title: %v, %v pages. Author: %v; ID: %v", b.Title, b.Pages, b.LastName, b.ID)
}
