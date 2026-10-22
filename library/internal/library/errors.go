package library

import "errors"

var (
	ErrUserNotAuthorised   = errors.New("user not authorised")
	ErrUncknownBook        = errors.New("unknown book")
	ErrTookedBook          = errors.New("this book has been toked by another client")
	ErrFreeBook            = errors.New("this book hasn't been tooked for someone")
	ErrRepeatingUsername   = errors.New("user with such username has been added eirleir")
	ErrUncknownUser        = errors.New("unknown user")
	ErrIncorrectUserAction = errors.New("incorrect user return book")
)
