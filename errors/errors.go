package errors

import (
	"fmt"

	"github.com/goccy/go-yaml/token"
)

type SyntaxError struct {
	Underlying error
	Token      *token.Token
	Msg        string
}

func (se SyntaxError) Error() string {
	return fmt.Sprintf("%v", se.Underlying)
}
