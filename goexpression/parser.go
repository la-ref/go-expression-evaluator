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
		make([]*Token, 0, len(input)),
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
		if t == NUMBER {
			val := ""
			for t == NUMBER && i < len(input) {
				char = Char(input[i])
				t, err = CheckType(char)
				fmt.Println(t, err, char)
				if err != nil {
					continue
				}
				if t != NUMBER {
					break
				}
				val += string(char)
				i++
			}
			err := p.AppendToken(t, val)
			if err != nil {
				continue
			}
			i--
		} else {
			err := p.AppendToken(t, string(char))
			if err != nil {
				continue
			}
		}

		i++
	}
	p.Print()
}

func (p *Parser) Print() {
	for i, v := range p.tokens {
		fmt.Printf("%d : [%d,%d]\n", i, v.t, v.value)
	}
}

// Add a new token at the end of the slice
func (p *Parser) AppendToken(t TokenType, value string) error {
	token, err := NewToken(value)
	if err != nil {
		return err
	}
	p.tokens = append(p.tokens, token)
	return nil
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
