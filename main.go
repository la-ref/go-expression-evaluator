package main

import (
	"fmt"
	"goexpressionevaluator/goexpression"
)

func main() {
	tok, err := goexpression.NewTokenizer("15+14-12")
	if err != nil {
		fmt.Println(err)
	}
	tok.Print()
	parser, err := goexpression.Parse("12*12+14-12*12*45")
	if err != nil {
		fmt.Println(err)
	}
	parser.Print()
}
