package errors_alias

import "errors"

var (
	ErrIncorrectUserData   = errors.New("incorrect data to create user")
	ErrUserNotAuthorized   = errors.New("user not authorised")
	ErrUnknownBook         = errors.New("unknown book")
	ErrBookAlreadyLoaned   = errors.New("this book has been toked by another client")
	ErrFreeBook            = errors.New("this book hasn't been tooked for someone")
	ErrRepeatingUsername   = errors.New("user with such username has been added eirleir")
	ErrUnknownUser         = errors.New("unknown user")
	ErrIncorrectUserAction = errors.New("incorrect user return book")
	ErrEmptyUserLibrary    = errors.New("user doesn't lended any books")
)
