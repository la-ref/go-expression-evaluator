package goexpression

import (
	"fmt"
)

type Parser struct {
	t            *Tokenizer // type
	currentToken *Token
}

func Parse(value string) (*AST, error) {
	tok, err := NewTokenizer(value)
	if err != nil {
		return nil, err
	}
	parser := &Parser{
		t:            tok,
		currentToken: tok.PopToken(),
	}
	ast, err := parser.ParseExpression()
	if err != nil {
		return nil, err
	}
	return ast, nil
}

func (p *Parser) Pop() {
	p.currentToken = p.t.PopToken()
}

// Parse addition and substraction
func (p *Parser) ParseExpression() (*AST, error) {
	firstValue, err := p.ParseTerms()
	if err != nil {
		return nil, err
	}
	for p.currentToken != nil && (p.currentToken.t == OPERATION_PLUS || p.currentToken.t == OPERATION_MINUS) {
		operator, err := GetTypeValue(p.currentToken.t)
		if err != nil {
			return nil, err
		}
		p.Pop()
		secondValue, err := p.ParseTerms()
		if err != nil {
			return nil, err
		}
		firstValue = &AST{T: OPERATOR, Value: operator, Left: firstValue, Right: secondValue}
	}
	return firstValue, nil

}

// Parse Multiplication and division
func (p *Parser) ParseTerms() (*AST, error) {
	firstValue, err := p.ParseFactor()
	if err != nil {
		return nil, err
	}
	for p.currentToken != nil && (p.currentToken.t == OPERATION_MULTIPLY || p.currentToken.t == OPERATION_DIVIDE) {
		operator, err := GetTypeValue(p.currentToken.t)
		if err != nil {
			return nil, err
		}
		p.Pop()
		secondValue, err := p.ParseFactor()
		if err != nil {
			return nil, err
		}
		firstValue = &AST{T: OPERATOR, Value: operator, Left: firstValue, Right: secondValue}
	}
	return firstValue, nil
}

func (p *Parser) ParseFactor() (*AST, error) {
	if p.currentToken != nil && p.currentToken.t == NUMBER {
		val := p.currentToken.value
		ast := &AST{
			T:     VALUE,
			Value: val,
		}
		p.Pop()
		return ast, nil
	}

	return nil, fmt.Errorf("Unknown Factor")
}
