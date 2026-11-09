package library_test

import (
	"library/internal/book"
	"library/internal/errors_alias"
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
		bServ       = library.NewCatalogService(bMem, nil)
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
		uMem  = user.NewUserMemory()
		lServ = library.NewLendingService(bMem, uMem)
		bServ = library.NewCatalogService(bMem, lServ)
	)

	bId, _ := bServ.AddBookToCatalog(&b1)
	bServ.AddBookToCatalog(&b2)

	err := bServ.DeleteBookFromCatalog(bId)
	assert.Nil(t, err)
}

func TestDeleteHandledBookFromCatalog(t *testing.T) {
	var (
		b1    = book.NewBook("A", "A", "A", "A", 10)
		bMem  = book.NewMemory()
		uMem  = user.NewUserMemory()
		uServ = library.NewUserService(uMem)
		lServ = library.NewLendingService(bMem, uMem)
		bServ = library.NewCatalogService(bMem, lServ)
	)

	bId, _ := bServ.AddBookToCatalog(&b1)
	i, _ := uServ.RegisterUser("f")
	err := lServ.GetBookToClient(i, bId)
	assert.Nil(t, err)

	err2 := bServ.DeleteBookFromCatalog(bId)
	assert.ErrorIs(t, err2, errors_alias.ErrTookedBook)
}

func TestDeleteBookFromCatalog_NotFound(t *testing.T) {
	var (
		b1    = book.NewBook("A", "A", "A", "A", 10)
		bMem  = book.NewMemory()
		uMem  = user.NewUserMemory()
		lServ = library.NewLendingService(bMem, uMem)
		bServ = library.NewCatalogService(bMem, lServ)
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

func TestRegisterUsedUsername(t *testing.T) {
	var (
		// u1 = user.NewUser("f")
		uMem  = user.NewUserMemory()
		uServ = library.NewUserService(uMem)
	)

	uServ.RegisterUser("f")
	_, err := uServ.RegisterUser("f")
	assert.ErrorIs(t, err, errors_alias.ErrRepeatingUsername)
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
		bServ = library.NewCatalogService(bMem, nil)
		lServ = library.NewLendingService(bMem, uMem)
	)

	bId, _ := bServ.AddBookToCatalog(&b1)
	i, _ := uServ.RegisterUser("f")
	err := lServ.GetBookToClient(i, bId)
	assert.Nil(t, err)
}

func TestGetHandlingBookToClient(t *testing.T) {
	var (
		b1    = book.NewBook("A", "A", "A", "A", 10)
		uMem  = user.NewUserMemory()
		uServ = library.NewUserService(uMem)
		bMem  = book.NewMemory()
		bServ = library.NewCatalogService(bMem, nil)
		lServ = library.NewLendingService(bMem, uMem)
	)

	bId, _ := bServ.AddBookToCatalog(&b1)
	i, _ := uServ.RegisterUser("f")
	i2, _ := uServ.RegisterUser("f2")
	err := lServ.GetBookToClient(i, bId)
	err2 := lServ.GetBookToClient(i2, bId)
	assert.Nil(t, err)
	assert.ErrorIs(t, err2, errors_alias.ErrTookedBook)
}

func TestReturnBookFromClient(t *testing.T) {
	var (
		b1    = book.NewBook("A", "A", "A", "A", 10)
		uMem  = user.NewUserMemory()
		uServ = library.NewUserService(uMem)
		bMem  = book.NewMemory()
		bServ = library.NewCatalogService(bMem, nil)
		lServ = library.NewLendingService(bMem, uMem)
	)

	bId, _ := bServ.AddBookToCatalog(&b1)
	i, _ := uServ.RegisterUser("f")
	lServ.GetBookToClient(i, bId)

	err := lServ.ReturnBookFromClient(i, bId)
	assert.Nil(t, err)
}

func TestReturnNotYoursBookFromClient(t *testing.T) {
	var (
		b1    = book.NewBook("A", "A", "A", "A", 10)
		uMem  = user.NewUserMemory()
		uServ = library.NewUserService(uMem)
		bMem  = book.NewMemory()
		bServ = library.NewCatalogService(bMem, nil)
		lServ = library.NewLendingService(bMem, uMem)
	)

	bId, _ := bServ.AddBookToCatalog(&b1)
	i, _ := uServ.RegisterUser("f")
	i2, _ := uServ.RegisterUser("f2")
	lServ.GetBookToClient(i, bId)

	err := lServ.ReturnBookFromClient(i2, bId)
	assert.ErrorIs(t, err, errors_alias.ErrIncorrectUserAction)
}

func TestShowClientsBooks(t *testing.T) {
	var (
		b1    = book.NewBook("A", "A", "A", "A", 10)
		uMem  = user.NewUserMemory()
		uServ = library.NewUserService(uMem)
		bMem  = book.NewMemory()
		bServ = library.NewCatalogService(bMem, nil)
		lServ = library.NewLendingService(bMem, uMem)
	)

	bServ.AddBookToCatalog(&b1)
	uServ.RegisterUser("f")
	lServ.GetBookToClient(0, 0)

	_, err := lServ.GetClientsBooks(0)
	assert.Nil(t, err)
}

func TestIsLended(t *testing.T) {
	var (
		b1    = book.NewBook("A", "A", "A", "A", 10)
		uMem  = user.NewUserMemory()
		uServ = library.NewUserService(uMem)
		bMem  = book.NewMemory()
		bServ = library.NewCatalogService(bMem, nil)
		lServ = library.NewLendingService(bMem, uMem)
	)

	bId, _ := bServ.AddBookToCatalog(&b1)
	i, _ := uServ.RegisterUser("f")
	lServ.GetBookToClient(i, bId)

	is := lServ.IsLended(bId)
	assert.True(t, is)
}

func TestReturningByIncorrectClient(t *testing.T) {
	var (
		b1    = book.NewBook("A", "A", "A", "A", 10)
		uMem  = user.NewUserMemory()
		uServ = library.NewUserService(uMem)
		bMem  = book.NewMemory()
		lServ = library.NewLendingService(bMem, uMem)
		bServ = library.NewCatalogService(bMem, lServ)
	)

	bId, _ := bServ.AddBookToCatalog(&b1)
	i, _ := uServ.RegisterUser("f")
	i2, _ := uServ.RegisterUser("f2")
	lServ.GetBookToClient(i, bId)
	err := lServ.ReturnBookFromClient(i2, bId)

	assert.ErrorIs(t, err, errors_alias.ErrIncorrectUserAction)
}
func TestReturnBookThatNotLended(t *testing.T) {
	var (
		b1    = book.NewBook("A", "A", "A", "A", 10)
		uMem  = user.NewUserMemory()
		uServ = library.NewUserService(uMem)
		bMem  = book.NewMemory()
		bServ = library.NewCatalogService(bMem, nil)
		lServ = library.NewLendingService(bMem, uMem)
	)

	bServ.AddBookToCatalog(&b1)
	uServ.RegisterUser("f")
	err := lServ.ReturnBookFromClient(0, 0)

	assert.ErrorIs(t, err, errors_alias.ErrFreeBook)
}
