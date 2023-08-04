package main

import (
	"fmt"
	"goexpressionevaluator/goexpression"
)

func main() {
	parser, err := goexpression.Parse("12*4+2*4+(5+6*(2+2))")
	if err != nil {
		fmt.Println(err)
	}
	parser.Print()
}
