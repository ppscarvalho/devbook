package respostas

import (
	"encoding/json"
	"reflect"
)

type Response struct {
	Mensagem string `json:"mensagem"`
}

func NewResponse(msg string) *Response {
	return &Response{Mensagem: msg}
}

func (r *Response) JSON(statusCode int) ([]byte, error) {
	return json.Marshal(r)
}

func (r *Response) Error(statusCode int) error {
	return NewError(statusCode, r.Mensagem)
}

type Error struct {
	StatusCode int    `json:"-"`
	Mensagem   string `json:"mensagem"`
}

// Error implements error.
func (e *Error) Error() string {
	panic("unimplemented")
}

func NewError(statusCode int, msg string) *Error {
	return &Error{StatusCode: statusCode, Mensagem: msg}
}

// Função para verificar se o struct está vazio
func IsEmptyStruct(v interface{}) bool {
	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}
