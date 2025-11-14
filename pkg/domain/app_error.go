package domain

import (
	"errors"
	"fmt"
	"log"
)

const (
	ErrCodeDuplicateKey        = "duplicate_key"
	ErrCodeInternalServerError = "internal_server_error"
	ErrCodeInvalidParams       = "invalid_params"
	ErrCodeNotFound            = "not_found"
	ErrCodeTimeout             = "timeout"
)

var (
	ErrDuplicateKey = errors.New("duplicate key error")
	ErrIncorrectID  = errors.New("incorrect id error")
	ErrNotFound     = errors.New("record not found error")
	ErrTimeout      = errors.New("timeout error")
)

// AppError es un tipo de error personalizado que implementa la interfaz error.
type AppError struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

// NewAppError crea un nuevo AppError con el código y mensaje dados.
func NewAppError(code string, msg string) AppError {
	return AppError{
		Code: code,
		Msg:  msg,
	}
}

// Error retorna una representación en cadena del error. Es parte de la interfaz error.
func (e AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Msg)
}

func ManageError(err error, msg string) error {
	var appErr AppError

	switch {
	case errors.Is(err, ErrDuplicateKey):
		log.Println("duplicate key")
		appErr = AppError{
			Code: ErrCodeDuplicateKey,
			Msg:  "Duplicate key",
		}
	case errors.Is(err, ErrIncorrectID):
		log.Println("incorrect id error")
		appErr = AppError{
			Code: ErrCodeInvalidParams,
			Msg:  "Incorrect id",
		}
	case errors.Is(err, ErrNotFound):
		log.Println("not found error")
		appErr = AppError{
			Code: ErrCodeNotFound,
			Msg:  "Not found",
		}
	case errors.Is(err, ErrTimeout):
		log.Println("timeout error")
		appErr = AppError{
			Code: ErrCodeTimeout,
			Msg:  "Timeout",
		}
	default:
		log.Println(err.Error())
		appErr = AppError{
			Code: ErrCodeInternalServerError,
			Msg:  "Server Error",
		}
	}

	// Solo agregamos el mensaje personalizado si el error no es un error interno del servidor
	if msg != "" && appErr.Code != ErrCodeInternalServerError {
		appErr.Msg = fmt.Sprintf("%s: %s", appErr.Msg, msg)
	}
	return appErr
}
