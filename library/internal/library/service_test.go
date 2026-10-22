package library_test

import (
	"fmt"
	"library/internal/book"
	"library/internal/library"
	"library/internal/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowBooks(t *testing.T) {
	var (
		b1          = book.NewBook("A", "A", "A", "A", 10)
		b2          = book.NewBook("A2", "A2", "A2", "A2", 10)
		bMem        = book.NewMemory()
		bServ       = library.NewCatalogService(bMem)
		expextedLen = 2
	)

	bServ.AddBookToCatalog(&b1)
	bServ.AddBookToCatalog(&b2)

	assert.Equal(t, expextedLen, bServ.AmountOfBooks())
}

func TestDeleteBookFromCatalog(t *testing.T) {
	var (
		b1    = book.NewBook("A", "A", "A", "A", 10)
		b2    = book.NewBook("A2", "A2", "A2", "A2", 10)
		bMem  = book.NewMemory()
		bServ = library.NewCatalogService(bMem)
	)

	bId := bServ.AddBookToCatalog(&b1)
	bServ.AddBookToCatalog(&b2)

	err := bServ.DeleteBookFromCatalog(bId)
	fmt.Println("Err in deliting:", err)
	assert.Nil(t, err)
}

func TestDeleteBookFromCatalog2(t *testing.T) {
	var (
		b1    = book.NewBook("A", "A", "A", "A", 10)
		bMem  = book.NewMemory()
		bServ = library.NewCatalogService(bMem)
	)

	bServ.AddBookToCatalog(&b1)

	err := bServ.DeleteBookFromCatalog(1)

	assert.NotNil(t, err)
}

func TestRegisterUser(t *testing.T) {
	var (
		// u1 = user.NewUser("f")
		uMem   = user.NewUserMemory()
		uServ  = library.NewUserService(uMem)
		expect = 0
	)

	id, err := uServ.RegisterUser("f")
	assert.Nil(t, err)
	assert.Equal(t, expect, id)
}

func TestLoginUser(t *testing.T) {
	var (
		// u1 = user.NewUser("f")
		uMem   = user.NewUserMemory()
		uServ  = library.NewUserService(uMem)
		expect = 1
	)

	uServ.RegisterUser("f")
	uServ.RegisterUser("f2")

	id, err := uServ.LogIn("f2")

	assert.Nil(t, err)
	assert.Equal(t, expect, id)
}

func TestGetBookToClient(t *testing.T) {
	var (
		b1    = book.NewBook("A", "A", "A", "A", 10)
		uMem  = user.NewUserMemory()
		uServ = library.NewUserService(uMem)
		bMem  = book.NewMemory()
		bServ = library.NewCatalogService(bMem)
		lServ = library.NewLendingService(uMem, bMem)
	)

	bId := bServ.AddBookToCatalog(&b1)
	i, _ := uServ.RegisterUser("f")
	err := lServ.GetBookToClient(i, bId)
	assert.Nil(t, err)
}

func TestReturnBookFromClient(t *testing.T) {
	var (
		b1    = book.NewBook("A", "A", "A", "A", 10)
		uMem  = user.NewUserMemory()
		uServ = library.NewUserService(uMem)
		bMem  = book.NewMemory()
		bServ = library.NewCatalogService(bMem)
		lServ = library.NewLendingService(uMem, bMem)
	)

	bId := bServ.AddBookToCatalog(&b1)
	i, _ := uServ.RegisterUser("f")
	lServ.GetBookToClient(i, bId)

	err := lServ.ReturnBookFromClient(i, bId)
	assert.Nil(t, err)
}

func TestShowClientsBooks(t *testing.T) {
	var (
		b1    = book.NewBook("A", "A", "A", "A", 10)
		uMem  = user.NewUserMemory()
		uServ = library.NewUserService(uMem)
		bMem  = book.NewMemory()
		bServ = library.NewCatalogService(bMem)
		lServ = library.NewLendingService(uMem, bMem)
	)

	bServ.AddBookToCatalog(&b1)
	uServ.RegisterUser("f")
	lServ.GetBookToClient(0, 0)

	err := lServ.ShowClientsBooks(0)
	assert.Nil(t, err)
}

func TestIsLended(t *testing.T) {
	var (
		b1    = book.NewBook("A", "A", "A", "A", 10)
		uMem  = user.NewUserMemory()
		uServ = library.NewUserService(uMem)
		bMem  = book.NewMemory()
		bServ = library.NewCatalogService(bMem)
		lServ = library.NewLendingService(uMem, bMem)
	)

	bId := bServ.AddBookToCatalog(&b1)
	i, _ := uServ.RegisterUser("f")
	lServ.GetBookToClient(i, bId)

	is := lServ.IsLended(bId)
	assert.True(t, is)
}
