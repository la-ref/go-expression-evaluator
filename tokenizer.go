package goexpression

import (
	"fmt"
	. "goexpression/utils"
	"strings"
)

type Tokenizer struct {
	tokens []*Token
}

func NewTokenizer(input string) (*Tokenizer, error) {
	tokenizer := &Tokenizer{
		make([]*Token, 0, len(input)),
	}
	err := tokenizer.Tokenize(input)
	if err != nil {
		return nil, err
	}
	return tokenizer, nil
}

func (p *Tokenizer) Tokenize(input string) error {
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
				if err != nil {
					return err
				}
				if t != NUMBER {
					break
				}
				val += string(char)
				i++
			}
			err := p.AppendToken(t, val)
			if err != nil {
				return err
			}
			i--
		} else {
			err := p.AppendToken(t, string(char))
			if err != nil {
				return err
			}
		}

		i++
	}
	return nil
}

func (p *Tokenizer) TokenizeNumber() {

}

func (p *Tokenizer) Print() {
	for i, v := range p.tokens {
		fmt.Printf("%d : [%d,%d]\n", i, v.t, v.value)
	}
}

// Add a new token at the end of the slice
func (p *Tokenizer) AppendToken(t TokenType, value string) error {
	token, err := NewToken(value)
	if err != nil {
		return err
	}
	p.tokens = append(p.tokens, token)
	return nil
}

// Pop the first token of the slice and returns it
func (p *Tokenizer) PopToken() *Token {
	var token *Token
	if len(p.tokens) > 0 {
		token = p.tokens[0]
		p.tokens = p.tokens[1:]
	}
	return token
}
