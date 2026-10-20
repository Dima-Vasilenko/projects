package main

import (
	"fmt"
	"library/internal/book"
	"library/internal/library"
)

type Library struct {
	system   library.LibrarySystem
	loggedIn bool
}

var booksId int = 0

func (l *Library) AddBook() {
	var t, descr, authorLast, authorFirst string
	var pages int
	fmt.Println("Enter book's title: ")
	fmt.Scan(&t)
	fmt.Println("Enter book's description: ")
	fmt.Scan(&descr)
	fmt.Println("Enter book's author lastname: ")
	fmt.Scan(&authorLast)
	fmt.Println("Enter book's author firstname: ")
	fmt.Scan(&authorFirst)
	fmt.Println("Enter book's amount of pages: ")
	fmt.Scan(&pages)
	var b book.Book = book.Book{
		Title:       t,
		Description: descr,
		Pages:       pages,
		Id:          booksId,
		WasTaken:    false,
		Author: book.Author{
			LastName:  authorLast,
			FirstName: authorFirst,
		},
	}
	l.system.AddBookToCatalog(b)
	booksId++
}

func (l *Library) DeleteBook() {
	var id int
	fmt.Println("Enter book's id, that you want to delete: ")
	fmt.Scan(&id)
	err := l.system.DeleteBookFromCatalog(id)
	if err != nil {
		fmt.Println(err)
	}
}

func (l *Library) ShowBooks() {
	l.system.ShowBooks()
}

func (l *Library) RegisterClient() {
	for true {
		var us string
		fmt.Println("Enter you username: ")
		fmt.Scan(&us)
		err := l.system.RegisterUser(us)
		if err != nil {
			fmt.Println("You have been succesfully registered")
		}
		fmt.Println("Such username is using in library enter something else")
	}
}

func (l *Library) Login() bool {
	var us string
	fmt.Println("Enter you username: ")
	fmt.Scan(&us)
	err := l.system.LogIn(us)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (l *Library) GetBookToClient() {
	var id int
	fmt.Println("Enter book's id that you want to take: ")
	fmt.Scan(&id)
	err := l.system.GetBookToClient(id)
	if err != nil {
		fmt.Println(err)
	}
}

func (l *Library) ReturnBookFromClient() {
	var id int
	fmt.Println("Enter book's id that you want to return: ")
	fmt.Scan(&id)
	err := l.system.ReturnBookFromClient(id)
	if err != nil {
		fmt.Println(err)
	}
}

func (l *Library) ShowClientBooks() {
	err := l.system.ShowClientsBooks()
	if err != nil {
		fmt.Println(err)
	}
}

func ShowMainMenu() (choice int) {
	fmt.Println("1 - add book to library")
	fmt.Println("2 - delete book from library")
	fmt.Println("3 - show books from library")
	fmt.Println("4 - register user")
	fmt.Println("5 - login in system")
	fmt.Println("6 - exit")
	fmt.Println("Choice what you want to do:")
	fmt.Scan(&choice)
	return
}

func ShowAuthMenu() (choice int) {
	fmt.Println("1 - show all book from catalog")
	fmt.Println("2 - get book")
	fmt.Println("3 - return book")
	fmt.Println("4 - show tooked books")
	fmt.Println("5 - log out")
	fmt.Println("6 - exit")
	fmt.Scan(&choice)
	return
}

func (l *Library) Start() {
	fmt.Println("Welocome to our Library!")
	for true {
		if !l.loggedIn {
			choice := ShowMainMenu()
			switch choice {
			case 1:
				l.AddBook()
			case 2:
				l.DeleteBook()
			case 3:
				l.ShowBooks()
			case 4:
				l.RegisterClient()
				l.loggedIn = true
			case 5:
				ok := l.Login()
				if ok {
					l.loggedIn = true
				}
			case 6:
				break
			default:
				fmt.Println("uncknown command")
			}
		}
		choice := ShowAuthMenu()
		switch choice {
		case 1:
			l.ShowBooks()
		case 2:
			l.GetBookToClient()
		case 3:
			l.ReturnBookFromClient()
		case 4:
			l.ShowClientBooks()
		case 5:
			l.system.LogOut()
			l.loggedIn = false
		case 6:
			break
		default:
			fmt.Println("uncknown command")
		}
	}
}

func main() {
	lib := Library{
		system: library.LibrarySystem{
			Users:   make([]library.User, 1),
			Current: -1,
			Catalog: book.NewMemory(),
		},
		loggedIn: false,
	}

	lib.Start()
}
