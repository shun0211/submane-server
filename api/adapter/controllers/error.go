package controllers

type Error struct {
	Message string
}

func NewError(err error) *Error {
	// NOTE: &では任意の型からそのポインタ型を生成できる
	return &Error{
		Message: err.Error(),
	}
}
