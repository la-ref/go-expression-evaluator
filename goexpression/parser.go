package goexpression

import (
	"fmt"
	. "goexpressionevaluator/goexpression/utils"
	"strings"
)

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
	for i < len(input) {
		char := Char(input[i])
		t, err := CheckType(char)
		if err != nil {
			continue
		}
		val := ""
		for t == NUMBER && i < len(input) {
			char = Char(input[i])
			t, err = CheckType(char)
			if err != nil {
				continue
			}
			val += string(char)
			i++
		}
		p.AppendToken(t, string(char))
		fmt.Println(val)

		i++
	}
}

func (p *Parser) Print() {
	for i, v := range p.tokens {
		fmt.Printf("%d : [%d,%d]\n", i, v.t, v.value)
	}
}

// Add a new token at the end of the slice
func (p *Parser) AppendToken(t TokenType, value string) {
	token, _ := NewToken(value)
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
