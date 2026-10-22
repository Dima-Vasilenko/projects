package main

import (
	"fmt"
	"library/internal/book"
	"library/internal/library"
	"library/internal/user"
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

type Session struct {
	clientId int
}

type Library struct {
	session    Session
	catService *library.CatalogService
	usService  *library.UserService
	lenService *library.LendingService
}

func (l *Library) AddBook() {
	title := GetString("Enter book's title:")
	descr := GetString("Enter book's description:")
	authorLast := GetString("Enter book's author lastname:")
	authorFirst := GetString("Enter book's author firstname:")
	pages := GetInt("Enter book's amount of pages:")

	b := book.NewBook(title, descr, authorLast, authorFirst, pages)
	l.catService.AddBookToCatalog(&b)
}

func (l *Library) DeleteBook() {
	id := GetInt("Enter book's id, that you want to delete:")
	if lended := l.lenService.IsLended(id); lended {
		fmt.Println("you can't delte book from catalog while it lending for someone")
		return
	}
	err := l.catService.DeleteBookFromCatalog(id)
	if err != nil {
		fmt.Println(err)
	}
}

func (l *Library) ShowBooks() {
	l.catService.ShowBooks()
}

func (l *Library) RegisterClient() {
	for {
		var err error
		us := GetString("Enter you username:")
		l.session.clientId, err = l.usService.RegisterUser(us)
		if err == nil {
			fmt.Println("You have been succesfully registered")
			fmt.Println("You user id: ", l.session.clientId)
			break
		}
		fmt.Println(err)
	}
}

func (l *Library) Login() bool {
	var err error
	us := GetString("Enter you username:")
	l.session.clientId, err = l.usService.LogIn(us)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (l *Library) GetBookToClient() {
	id := GetInt("Enter book's id that you want to take:")
	err := l.lenService.GetBookToClient(l.session.clientId, id)
	if err != nil {
		fmt.Println(err)
	}
}

func (l *Library) ReturnBookFromClient() {
	id := GetInt("Enter book's id that you want to return:")
	err := l.lenService.ReturnBookFromClient(l.session.clientId, id)
	if err != nil {
		fmt.Println(err)
	}
}

func (l *Library) ShowClientBooks() {
	err := l.lenService.ShowClientsBooks(l.session.clientId)
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
	for {
		var choice int
		if l.session.clientId == -1 {
			choice = ShowMainMenu()
		} else {
			choice = ShowAuthMenu()
		}
		switch choice {
		case 1:
			l.ShowBooks()
		case 2:
			if l.session.clientId == -1 {
				l.AddBook()
				continue
			}
			l.GetBookToClient()
		case 3:
			if l.session.clientId == -1 {
				l.DeleteBook()
				continue
			}
			l.ReturnBookFromClient()
		case 4:
			if l.session.clientId == -1 {
				l.RegisterClient()
				continue
			}
			l.ShowClientBooks()
		case 5:
			if l.session.clientId == -1 {
				l.Login()
				continue
			}
			l.session.clientId = -1
		case 6:
			if l.session.clientId != -1 {
				fmt.Println("uncknown command")
				continue
			}
			l.usService.ShowListOfUsers()
		case 0:
			return
		default:
			fmt.Println("uncknown command")
		}
	}
}

func main() {

	ses := Session{clientId: -1}

	usMem := user.NewUserMemory()
	bookMem := book.NewMemory()

	catServ := library.NewCatalogService(bookMem)
	usServ := library.NewUserService(usMem)
	lenServ := library.NewLendingService(bookMem)

	lib := Library{
		session:    ses,
		catService: catServ,
		usService:  usServ,
		lenService: lenServ,
	}
	lib.Start()
}
