package controllers

type Error struct {
	Message string
	forDeveloperMessage string
}

func NewError(message string, forDeveloperMessage string) *Error {
	// NOTE: &では任意の型からそのポインタ型を生成できる
	return &Error{
		Message: message,
		forDeveloperMessage: forDeveloperMessage,
	}
}
