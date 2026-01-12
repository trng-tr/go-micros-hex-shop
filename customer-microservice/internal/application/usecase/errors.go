package usecase

import "errors"

var (
	errSendObject     error = errors.New("error of sending bs object")
	errRetrieveObject error = errors.New("error: impossible to retrieve demanded objet")
)
