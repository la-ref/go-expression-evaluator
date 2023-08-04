package goexpression

import (
	"fmt"
	"strings"
)

type TokenType int

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

type Parser struct {
	tokens []Token
}

func (p *Parser) Parse(input string) {
	input = strings.TrimSpace(input)
	i := 0
	for i != len(input)-1 {
		char := input[i]
		fmt.Println(char)
		i++
	}
}
