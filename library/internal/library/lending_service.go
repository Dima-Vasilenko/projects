package library

import (
	"library/internal/book"
	"library/internal/errors_alias"
	"slices"
)

type Loan struct {
	UserId int
	BookId int
}

type LendingService struct {
	bookRepo BookRepository
	userRepo UserRepository
	loan     []*Loan
}

func NewLendingService(bookMem BookRepository, usMem UserRepository) *LendingService {
	return &LendingService{
		bookRepo: bookMem,
		userRepo: usMem,
		loan:     make([]*Loan, 0),
	}
}

func (l *LendingService) GetBookToClient(clientId, id int) error {
	if !l.userRepo.IsIDExist(clientId) {
		return errors_alias.ErrUserNotAuthorised
	}

	_, ok := l.bookRepo.GetBookById(id)
	if ok != true {
		return errors_alias.ErrUncknownBook
	}

	if index := slices.IndexFunc(l.loan, func(l *Loan) bool {
		return l.BookId == id
	}); index != -1 {
		return errors_alias.ErrTookedBook
	}

	l.loan = append(l.loan, &Loan{clientId, id})
	return nil
}

func (l *LendingService) ReturnBookFromClient(clientId, id int) error {
	if clientId == -1 {
		return errors_alias.ErrUserNotAuthorised
	}

	_, ok := l.bookRepo.GetBookById(id)
	if ok != true {
		return errors_alias.ErrUncknownBook
	}

	index := slices.IndexFunc(l.loan, func(l *Loan) bool {
		return l.BookId == id
	})
	if index == -1 {
		return errors_alias.ErrFreeBook
	}
	if l.loan[index].UserId != clientId {
		return errors_alias.ErrIncorrectUserAction
	}

	l.loan = append(l.loan[:index], l.loan[index+1:]...)
	return nil
}

func (l *LendingService) GetClientsBooks(clientId int) ([]*book.Book, error) {
	if clientId == -1 {
		return nil, errors_alias.ErrUserNotAuthorised
	}
	var res []*book.Book
	for _, v := range l.loan {
		if v.UserId == clientId {
			b, _ := l.bookRepo.GetBookById(v.BookId)
			res = append(res, b)
		}
	}
	return res, nil
}

func (l *LendingService) IsLended(id int) bool {
	ok := slices.ContainsFunc(l.loan, func(l *Loan) bool {
		return l.BookId == id
	})
	return ok
}
