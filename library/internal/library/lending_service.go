package library

import (
	"fmt"
	"slices"
)

type Loan struct {
	UserId int
	BookId int
}

type LendingService struct {
	bookRepo BookRepository
	loan     []*Loan
}

func NewLendingService(bookMem BookRepository) *LendingService {
	return &LendingService{
		bookRepo: bookMem,
		loan:     make([]*Loan, 0),
	}
}

func (l *LendingService) GetBookToClient(clientId, id int) error {
	if clientId == -1 {
		return ErrUserNotAuthorised
	}

	_, ok := l.bookRepo.GetBookById(id)
	if ok != true {
		return ErrUncknownBook
	}

	if index := slices.IndexFunc(l.loan, func(l *Loan) bool {
		return l.BookId == id
	}); index != -1 {
		return ErrTookedBook
	}

	l.loan = append(l.loan, &Loan{clientId, id})
	return nil
}

func (l *LendingService) ReturnBookFromClient(clientId, id int) error {
	if clientId == -1 {
		return ErrUserNotAuthorised
	}

	_, ok := l.bookRepo.GetBookById(id)
	if ok != true {
		return ErrUncknownBook
	}

	index := slices.IndexFunc(l.loan, func(l *Loan) bool {
		return l.BookId == id
	})
	if index == -1 {
		return ErrFreeBook
	}
	if l.loan[index].UserId != clientId {
		return ErrIncorrectUserAction
	}

	l.loan = append(l.loan[:index], l.loan[index+1:]...)
	return nil
}

func (l *LendingService) ShowClientsBooks(clientId int) error {
	if clientId == -1 {
		return ErrUserNotAuthorised
	}
	for _, v := range l.loan {
		if v.UserId == clientId {
			b, _ := l.bookRepo.GetBookById(v.BookId)
			fmt.Println(b)
		}
	}
	return nil
}

func (l *LendingService) IsLended(id int) bool {
	ok := slices.ContainsFunc(l.loan, func(l *Loan) bool {
		return l.BookId == id
	})
	return ok
}
