package library

import (
	"errors"
	"fmt"
	"library/internal/book"
	"library/internal/user"
	"slices"
)

type Loan struct {
	UserId int
	BookId int
}

type LendingService struct {
	usRepo   UserRepository
	bookRepo BookRepository
	loan     []*Loan
}

func NewLendingService(usMem *user.UserMemory, bookMem *book.Memory) LendingService {
	return LendingService{
		usRepo:   usMem,
		bookRepo: bookMem,
		loan:     make([]*Loan, 0),
	}
}

func (l *LendingService) GetBookToClient(clientId, id int) error {
	if clientId == -1 {
		return errors.New("unathorised getting books")
	}

	_, ok := l.bookRepo.GetBookById(id)
	if ok != true {
		return errors.New("not such book in lib with this id")
	}

	if index := slices.IndexFunc(l.loan, func(l *Loan) bool {
		return l.BookId == id
	}); index != -1 {
		return errors.New("this book has been toked by another client, choise something else")
	}

	l.loan = append(l.loan, &Loan{clientId, id})
	return nil
}

func (l *LendingService) ReturnBookFromClient(clientId, id int) error {
	if clientId == -1 {
		return errors.New("unathorised returning books")
	}

	_, ok := l.bookRepo.GetBookById(id)
	if ok != true {
		return errors.New("not such book in lib with this id")
	}

	index := slices.IndexFunc(l.loan, func(l *Loan) bool {
		return l.BookId == id
	})
	if index == -1 {
		return errors.New("this book hasn't been tooked for someone")
	}

	l.loan = append(l.loan[:index], l.loan[index+1:]...)
	return nil
}

func (l *LendingService) ShowClientsBooks(clientId int) error {
	if clientId == -1 {
		return errors.New("unathorised getting books")
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
