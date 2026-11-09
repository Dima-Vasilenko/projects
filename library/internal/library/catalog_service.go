package library

import (
	"library/internal/book"
	"library/internal/errors_alias"
	"slices"
)

type BookRepository interface {
	AddBook(b *book.Book) error
	DeleteBookById(id int) error
	GetBookById(id int) (*book.Book, bool)
	GetListOfBooks() []*book.Book
}

type CatalogService struct {
	bookRepo    BookRepository
	lendingServ *LendingService
}

func NewCatalogService(repo BookRepository, l *LendingService) *CatalogService {
	return &CatalogService{
		bookRepo:    repo,
		lendingServ: l,
	}
}

func (c *CatalogService) GetBooks() []*book.Book {
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
	if exist := slices.ContainsFunc(c.lendingServ.GetLoan(), func(l *Loan) bool {
		return id == l.BookId
	}); exist {
		return errors_alias.ErrTookedBook
	}
	res := c.bookRepo.DeleteBookById(id)
	return res
}

func (c *CatalogService) AmountOfBooks() int {
	return len(c.bookRepo.GetListOfBooks())
}
