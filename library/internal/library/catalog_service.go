package library

import (
	"library/internal/book"
	"library/internal/errors_alias"
)

type BookRepository interface {
	AddBook(b *book.Book) error
	DeleteBookById(id int) error
	GetBookById(id int) (book.Book, bool)
	GetListOfBooks() []book.Book
}

type BookIsLended interface {
	IsLended(id int) bool
}

type CatalogService struct {
	bookRepo       BookRepository
	lendingChecker BookIsLended
}

func NewCatalogService(repo BookRepository, l *LendingService) *CatalogService {
	return &CatalogService{
		bookRepo:       repo,
		lendingChecker: l,
	}
}

func (c *CatalogService) GetBooks() []book.Book {
	return c.bookRepo.GetListOfBooks()
}

func (c *CatalogService) AddBookToCatalog(b *book.Book) (int, error) {
	err := c.bookRepo.AddBook(b)
	if err != nil {
		return -1, err
	}
	return b.ID, nil
}

func (c *CatalogService) DeleteBookFromCatalog(id int) error {
	if c.lendingChecker.IsLended(id) {
		return errors_alias.ErrBookAlreadyLoaned
	}

	res := c.bookRepo.DeleteBookById(id)
	return res
}

func (c *CatalogService) AmountOfBooks() int {
	return len(c.bookRepo.GetListOfBooks())
}
