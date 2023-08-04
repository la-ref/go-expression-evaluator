package goexpression

func Evaluate(value string) (float64, error) {
	ast, err := Parse(value)
	if err != nil {
		return 0, err
	}
	res := eval(ast)
	return res, nil
}

func eval(ast *AST) float64 {
	switch ast.T {
	case VALUE:
		return float64(ast.Value.(int))
	case OPERATOR:
		switch ast.Value {
		case "+":
			return eval(ast.Left) + eval(ast.Right)
		case "-":
			return eval(ast.Left) - eval(ast.Right)
		case "*":
			return eval(ast.Left) * eval(ast.Right)
		case "/":
			return eval(ast.Left) / eval(ast.Right)
		}
	}
	return 0
}
