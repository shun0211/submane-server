package controllers

import "fmt"

type Error struct {
	Message string `json:"message"`
	ForDeveloperMessage string `json:"forDeveloperMessage"`
}

func NewError(message string, forDeveloperMessage string) *Error {
	// NOTE: &では任意の型からそのポインタ型を生成できる
	return &Error{
		Message: message,
		ForDeveloperMessage: forDeveloperMessage,
	}
}

func NotFoundError(resource string) *Error {
	return NewError(
		fmt.Sprintf("%s が見つかりませんでした😱", resource),
		"Echo Server Not Found",
	)
}
