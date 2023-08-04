package goexpression

import (
	"goexpressionevaluator/goexpression/utils"
	"strings"
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
func NewToken(t TokenType, value int) *Token {
	token := new(Token)
	if t == NUMBER {
		token.value = value
	} else {
		token.value = 0
	}
	token.t = t
	return token
}

type Parser struct {
	tokens []*Token
}

func Evaluate(input string) {
	parser := &Parser{
		make([]*Token, len(input)),
	}
	parser.Parse(input)
}

func (p *Parser) Parse(input string) {
	input = strings.TrimSpace(input)
	i := 0
	for i != len(input) {
		char := input[i]
		for utils.IsCharANumberASCII(char) {

		}

		i++
	}
}

// Add a new token at the end of the slice
func (p *Parser) AppendToken(t TokenType, value int) {
	token := NewToken(t, value)
	p.tokens = append(p.tokens, token)
}

// Pop the first token of the slice and returns it
func (p *Parser) PopToken() *Token {
	var token *Token
	if len(p.tokens) > 0 {
		token = p.tokens[0]
		p.tokens = p.tokens[1:]
	}
	return token
}
