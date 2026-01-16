package usecase

import "errors"

var (
	errSavingObject      error = errors.New("error of saving object")
	errObjectNotFound    error = errors.New("error: object not found")
	errPrintingObjects   error = errors.New("error: impossible to print data")
	ErrInsufficientStock error = errors.New("error: quantity for the product in stock insufficient")
	errNoDataRegistered  error = errors.New("error: no data registered")
)
