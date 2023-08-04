package goexpression

type ASTType int

const (
	OPERATOR = iota
	VALUE
)

type AST struct {
	left  *AST
	right *AST
	t     ASTType
	value int
}

func (a *AST) AddRight(r *AST) {
	a.right = r
}

func (a *AST) AddLeft(l *AST) {
	a.left = l
}

type Parser struct {
	t            Tokenizer
	currentToken *Token
	ast          *AST
}

func (p *Parser) Pop() {
	p.currentToken = p.t.PopToken()
}

// Parse addition and substraction
func (p *Parser) ParseExpression() *AST {
	p.ParseTerms()
	return nil

}

// Parse Multiplication and division
func (p *Parser) ParseTerms() *AST {
	p.ParseFactor()
	return nil
}

func (p *Parser) ParseFactor() *AST {
	if p.currentToken.t == NUMBER {
		val := p.currentToken.value
		ast := &AST{
			t:     VALUE,
			value: val,
		}
		p.Pop()
		return ast
	}

	return nil
}
