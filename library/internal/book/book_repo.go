package book

import (
	"errors"
	"slices"
)

type Memory struct {
	books []*Book
}

func NewMemory() *Memory {
	return new(Memory{
		books: make([]*Book, 0),
	})
}

func (m *Memory) AddBook(b *Book) {
	m.books = append(m.books, b)
}

func (m *Memory) DeleteBookByIndex(index int) {

	copy(m.books[index:], m.books[index+1:])

	m.books[len(m.books)-1] = nil
	m.books = m.books[:len(m.books)-1]
}

func (m *Memory) DeleteBookById(id int) error {
	index := slices.IndexFunc(m.books, func(b *Book) bool {
		return b.Id == id
	})

	if index == -1 {
		return errors.New("undefind book's id")
	}

	m.DeleteBookByIndex(index)
	return nil
}

func (m *Memory) GetBookById(id int) (*Book, bool) {
	for _, val := range m.books {
		if val.Id == id {
			return val, true
		}
	}
	return nil, false
}

func (m *Memory) GetListOfBooks() []*Book {
	return m.books
}
