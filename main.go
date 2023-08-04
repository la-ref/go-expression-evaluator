package main

import (
	"fmt"
	"goexpressionevaluator/goexpression"
)

func main() {
	parser, err := goexpression.Parse("12/3")
	if err != nil {
		fmt.Println(err)
	}
	parser.Print()
	fmt.Println(goexpression.Evaluate("-13/-3"))
}
