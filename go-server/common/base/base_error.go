package base

type BaseError struct {
	ErrNo  int64
	ErrMsg string
}

func (e BaseError) Error() string {
	return e.ErrMsg
}

var ErrUnmarshalJsonParam *BaseError = &BaseError{1001, "参数错误"}
