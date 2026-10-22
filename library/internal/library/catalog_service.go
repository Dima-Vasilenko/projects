package library

import (
	"fmt"
	"library/internal/book"
)

type BookRepository interface {
	AddBook(b *book.Book)
	DeleteBookByIndex(index int)
	GetBookById(id int) (*book.Book, bool)
	GetListOfBooks() []*book.Book
	DeleteBookById(id int) error
}

type CatalogService struct {
	bookRepo BookRepository
}

func NewCatalogService(bookMem *book.Memory) CatalogService {
	return CatalogService{
		bookRepo: bookMem,
	}
}

func (c *CatalogService) ShowBooks() {
	for _, v := range c.bookRepo.GetListOfBooks() {
		fmt.Println(v)
	}
}

func (c *CatalogService) AddBookToCatalog(b *book.Book) {
	c.bookRepo.AddBook(b)
}

func (c *CatalogService) DeleteBookFromCatalog(id int) error {
	result := c.bookRepo.DeleteBookById(id)
	return result
}
