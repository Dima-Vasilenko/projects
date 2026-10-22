package library

import (
	"fmt"
	"library/internal/book"
	"slices"
)

type BookRepository interface {
	AddBook(b *book.Book)
	DeleteBookByIndex(index int)
	GetBookById(id int) (*book.Book, bool)
	GetListOfBooks() []*book.Book
}

type CatalogService struct {
	bookRepo BookRepository
}

func NewCatalogService(repo BookRepository) *CatalogService {
	return &CatalogService{
		bookRepo: repo,
	}
}

func (c *CatalogService) ShowBooks() {
	for _, v := range c.bookRepo.GetListOfBooks() {
		fmt.Println(v)
	}
}

func (c *CatalogService) AddBookToCatalog(b *book.Book) int {
	c.bookRepo.AddBook(b)
	return b.ID
}

func (c *CatalogService) DeleteBookFromCatalog(id int) error {
	index := slices.IndexFunc(c.bookRepo.GetListOfBooks(), func(b *book.Book) bool {
		return b.ID == id
	})

	if index == -1 {
		return ErrUncknownBook
	}
	c.bookRepo.DeleteBookByIndex(index)
	return nil
}

func (c *CatalogService) AmountOfBooks() int {
	return len(c.bookRepo.GetListOfBooks())
}
