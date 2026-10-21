package library

import (
	"errors"
	"fmt"
	"library/internal/book"
	"slices"
)

type Repository interface {
	AddBook(b book.Book)
	DeleteBookByIndex(index int)
	GetBookById(id int) (*book.Book, bool)
	GetListOfBooks() []*book.Book
}

type User struct {
	Username string
	books    []int
}

type LibrarySystem struct {
	Users   []*User
	Current int
	Catalog Repository
}

func (l *LibrarySystem) RegisterUser(us string) error {
	exists := slices.ContainsFunc(l.Users, func(u *User) bool {
		return u.Username == us
	})
	if exists {
		return errors.New("user with such username has been added eirleir, try something another")
	}

	books := make([]int, 0)
	l.Users = append(l.Users, &User{us, books})
	l.Current = len(l.Users) - 1
	return nil
}

func (l *LibrarySystem) LogOut() {
	l.Current = -1
}

func (l *LibrarySystem) LogIn(us string) error {
	index := slices.IndexFunc(l.Users, func(u *User) bool {
		return u.Username == us
	})
	if index == -1 {
		return errors.New("no such user has been registed")
	}
	l.Current = index
	return nil
}

func (l *LibrarySystem) ShowListOfUsers() {
	for i, val := range l.Users {
		fmt.Printf("User %d: %s\n", i, val.Username)
	}
}

func (l *LibrarySystem) ShowBooks() {
	for _, v := range l.Catalog.GetListOfBooks() {
		fmt.Println(v)
	}
}

func (l *LibrarySystem) GetBookToClient(id int) error {
	if l.Current == -1 {
		return errors.New("unathorised getting books")
	}

	b, ok := l.Catalog.GetBookById(id)
	if ok != true {
		return errors.New("not such book in lib with this id")
	}

	if b.HasTaken() {
		return errors.New("this book has been toked by another client, choise something else")
	}
	l.Users[l.Current].books = append(l.Users[l.Current].books, id)
	b.SetTaken()

	return nil
}

func (l *LibrarySystem) ReturnBookFromClient(id int) error {
	if l.Current == -1 {
		return errors.New("unathorised returning books")
	}

	b, ok := l.Catalog.GetBookById(id)
	if ok != true {
		return errors.New("not such book in lib with this id")
	}

	if !b.HasTaken() {
		return errors.New("this book hasn't been tooked for someone")
	}

	index := slices.Index(l.Users[l.Current].books, id)
	if index == -1 {
		return errors.New("no such book in user library")
	}

	l.Users[l.Current].books = append(l.Users[l.Current].books[:index], l.Users[l.Current].books[index+1:]...)
	b.SetFree()
	return nil
}

func (l *LibrarySystem) ShowClientsBooks() error {
	if l.Current == -1 {
		return errors.New("unathorised getting books")
	}
	for _, v := range l.Users[l.Current].books {
		fmt.Println("Book's id to show: ", v)
		b, _ := l.Catalog.GetBookById(v)
		fmt.Println(b)
	}
	return nil
}

func (l *LibrarySystem) AddBookToCatalog(b book.Book) {
	l.Catalog.AddBook(b)
}

func (l *LibrarySystem) DeleteBookFromCatalog(id int) error {
	b, ok := l.Catalog.GetBookById(id)
	if !ok {
		return errors.New("no such book in catalog")
	}
	if b.WasTaken {
		return errors.New("cannot delete book that has tooked by user")
	}
	index := slices.Index(l.Catalog.GetListOfBooks(), b)
	l.Catalog.DeleteBookByIndex(index)
	return nil
}
