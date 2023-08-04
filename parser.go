package goexpressionevaluatior

const (
	NUMBER = iota
	OPERATION_PLUS
	OPERATION_MINUS
	OPERATION_MULTIPLY
	OPERATION_DIVIDE
)

type Token struct {
}

type Parser struct {
	tokens []Token
}
