package main

var (
	ErrExpressionBlank           = CustomError("ErrExpressionBlank", "expression cannot blank")
	ErrExpressionContainAlphabet = CustomError("ErrExpressionContainAlphabet", "expression cannot contain alphabet")
)

type customError struct {
	k string
	v string
}

func (ce *customError) Error() string {
	return ce.v
}

func (ce *customError) Key() string {
	return ce.k
}

func CustomError(k, v string) *customError {
	return &customError{k, v}
}
