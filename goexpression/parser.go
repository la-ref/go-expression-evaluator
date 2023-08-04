package goexpression

import (
	"fmt"
	"goexpressionevaluator/goexpression/utils"
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

func Evaluate(input string) {
	parser := &Parser{
		make([]Token, len(input)),
	}
	parser.Parse(input)
}

func (p *Parser) Parse(input string) {
	input = strings.TrimSpace(input)
	i := 0
	for i != len(input) {
		char := input[i]
		fmt.Println(utils.IsCharANumberASCII(char))
		i++
	}
}
