package book

import (
	"errors"
	"slices"
)

type Memory struct {
	books  []*Book
	nextID int
}

func NewMemory() *Memory {
	return &(Memory{
		books:  make([]*Book, 0),
		nextID: 0,
	})
}

func (m *Memory) AddBook(b *Book) {
	b.ID = m.nextID
	m.nextID++
	m.books = append(m.books, b)
}

func (m *Memory) DeleteBookByIndex(index int) {

	copy(m.books[index:], m.books[index+1:])

	m.books[len(m.books)-1] = nil
	m.books = m.books[:len(m.books)-1]
}

func (m *Memory) DeleteBookById(id int) error {
	index := slices.IndexFunc(m.books, func(b *Book) bool {
		return b.ID == id
	})

	if index == -1 {
		return errors.New("undefind book's id")
	}

	m.DeleteBookByIndex(index)
	return nil
}

func (m *Memory) GetBookById(id int) (*Book, bool) {
	for _, val := range m.books {
		if val.ID == id {
			return val, true
		}
	}
	return nil, false
}

func (m *Memory) GetListOfBooks() []*Book {
	return m.books
}
