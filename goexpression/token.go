package goexpression

import (
	"fmt"
	. "goexpressionevaluator/goexpression/utils"
	"strconv"
)

type TokenType int

// Token types
const (
	NUMBER TokenType = iota
	OPERATION_PLUS
	OPERATION_MINUS
	OPERATION_MULTIPLY
	OPERATION_DIVIDE
	PARENTHESIS_LEFT
	PARENTHESIS_RIGHT
)

type Token struct {
	t     TokenType
	value int
}

// Create a new token with its type and its value
func NewToken(value string) (*Token, error) {
	token := new(Token)
	t, err := CheckType(value)
	if err != nil {
		return nil, err
	}
	token.t = t
	token.value = 0
	if t == NUMBER {
		token.value, _ = strconv.Atoi(string(value))
	}
	return token, nil
}

func CheckType[T string | Char](value T) (TokenType, error) {
	if ISANumber[T](value) {
		return NUMBER, nil
	}
	switch value {
	case "+":
		return OPERATION_PLUS, nil
	case "-":
		return OPERATION_MINUS, nil
	case "*":
		return OPERATION_MULTIPLY, nil
	case "/":
		return OPERATION_DIVIDE, nil
	case "(":
		return PARENTHESIS_LEFT, nil
	case ")":
		return PARENTHESIS_RIGHT, nil
	}
	return 0, fmt.Errorf("Unknow Token type")
}
