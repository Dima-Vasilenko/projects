package main

import (
	"fmt"
	"library/internal/book"
	"library/internal/library"
)

func GetInt(mess string) (res int) {
	fmt.Println(mess)
	fmt.Scan(&res)
	return
}

func GetString(mess string) (res string) {
	fmt.Println(mess)
	fmt.Scan(&res)
	return
}

type Library struct {
	system   library.LibrarySystem
	loggedIn bool
}

var booksId int = 0

func (l *Library) AddBook() {
	title := GetString("Enter book's title:")
	descr := GetString("Enter book's description:")
	authorLast := GetString("Enter book's author lastname:")
	authorFirst := GetString("Enter book's author firstname:")
	pages := GetInt("Enter book's amount of pages:")

	var book book.Book = book.Book{
		Title:       title,
		Description: descr,
		Pages:       pages,
		Id:          booksId,
		WasTaken:    false,
		Author: book.Author{
			LastName:  authorLast,
			FirstName: authorFirst,
		},
	}
	l.system.AddBookToCatalog(book)
	booksId++
}

func (l *Library) DeleteBook() {
	id := GetInt("Enter book's id, that you want to delete:")
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
		us := GetString("Enter you username:")
		err := l.system.RegisterUser(us)
		if err == nil {
			fmt.Println("You have been succesfully registered")
			break
		}
		fmt.Println(err)
	}
}

func (l *Library) Login() bool {
	us := GetString("Enter you username:")
	err := l.system.LogIn(us)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (l *Library) GetBookToClient() {
	id := GetInt("Enter book's id that you want to take:")
	err := l.system.GetBookToClient(id)
	if err != nil {
		fmt.Println(err)
	}
}

func (l *Library) ReturnBookFromClient() {
	id := GetInt("Enter book's id that you want to return:")
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
	fmt.Println("1 - show all book from catalog")
	fmt.Println("2 - add book to library")
	fmt.Println("3 - delete book from library")
	fmt.Println("4 - register user")
	fmt.Println("5 - login in system")
	fmt.Println("6 - show list of users")
	fmt.Println("0 - exit")
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
	fmt.Println("0 - exit")
	fmt.Scan(&choice)
	return
}

func (l *Library) Start() {
	fmt.Println("Welocome to our Library!")
	for true {
		var choice int
		if !l.loggedIn {
			choice = ShowMainMenu()
		} else {
			choice = ShowAuthMenu()
		}
		switch choice {
		case 1:
			l.ShowBooks()
		case 2:
			if !l.loggedIn {
				l.AddBook()
				continue
			}
			l.GetBookToClient()
		case 3:
			if !l.loggedIn {
				l.DeleteBook()
				continue
			}
			l.ReturnBookFromClient()
		case 4:
			if !l.loggedIn {
				l.RegisterClient()
				l.loggedIn = true
				continue
			}
			l.ShowClientBooks()
		case 5:
			if !l.loggedIn {
				ok := l.Login()
				if ok {
					l.loggedIn = true
				}
				continue
			}
			l.system.LogOut()
			l.loggedIn = false
		case 6:
			l.system.ShowListOfUsers()
			if l.loggedIn {
				fmt.Println("uncknown command")
			}
		case 0:
			return
		default:
			fmt.Println("uncknown command")
		}
	}
}

func main() {
	lib := Library{
		system: library.LibrarySystem{
			Users:   make([]*library.User, 0),
			Current: -1,
			Catalog: book.NewMemory(),
		},
		loggedIn: false,
	}

	lib.Start()
}
