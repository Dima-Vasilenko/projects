package book

import (
	"library/internal/errors_alias"
	"slices"
)

type Memory struct {
	books  []Book
	nextID int
}

func NewMemory() *Memory {
	return &(Memory{
		books:  make([]Book, 0),
		nextID: 0,
	})
}

func (m *Memory) AddBook(b *Book) error {
	if b == nil {
		return errors_alias.ErrUnknownBook
	}
	b.ID = m.nextID
	m.nextID++
	m.books = append(m.books, *b)
	return nil
}

func (m *Memory) deleteBookByIndex(index int) {

	copy(m.books[index:], m.books[index+1:])

	m.books[len(m.books)-1] = Book{}
	m.books = m.books[:len(m.books)-1]
}

func (m *Memory) DeleteBookById(id int) error {
	index := slices.IndexFunc(m.books, func(b Book) bool {
		return b.ID == id
	})

	if index == -1 {
		return errors_alias.ErrUnknownBook
	}

	m.deleteBookByIndex(index)
	return nil
}

func (m *Memory) GetBookById(id int) (Book, bool) {
	for _, val := range m.books {
		if val.ID == id {
			return val, true
		}
	}
	return Book{}, false
}

func (m *Memory) GetListOfBooks() []Book {
	return append([]Book{}, m.books...)
}
