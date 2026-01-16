package validators

import "errors"

var (
	errEmptyFields       error = errors.New("error: empty value for fied")
	errTooLong           error = errors.New("error: too long value for fied")
	errTooShort          error = errors.New("error: too short value for fied")
	errInvalidId         error = errors.New("error: provided id is invalid")
	errInvalidUnitPrice  error = errors.New("error: invalid unit price")
	errInvalidCurrency   error = errors.New("error: invalid currency for product")
	errInvalidProductCat error = errors.New("error: invalid product category")
	errInvalidStockField error = errors.New("error: invalid value for field")
	errInvalidStockQty   error = errors.New("error: invalid value for stock quantity")
)
