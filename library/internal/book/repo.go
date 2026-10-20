package book

type Memory struct {
	books []Book
}

func NewMemory() *Memory {
	return new(Memory{
		books: make([]Book, 0),
	})
}

func (m *Memory) AddBook(b Book) {
	m.books = append(m.books, b)
}

func (m *Memory) DeleteBookByIndex(index int) {

	copy(m.books[index:], m.books[index+1:])

	m.books[len(m.books)-1] = Book{}
	m.books = m.books[:len(m.books)-1]
}

func (m *Memory) GetBookById(id int) (Book, bool) {
	for _, val := range m.books {
		if val.Id == id {
			return val, true
		}
	}
	return Book{}, false
}

func (m *Memory) GetListOfBooks() []Book {
	return m.books
}
