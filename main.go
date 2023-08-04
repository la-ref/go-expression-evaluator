package main

import (
	"fmt"
	"goexpressionevaluator/goexpression"
	"goexpressionevaluator/goexpression/utils"
)

func main() {
	fmt.Println(utils.IsCharANumber("-"))
	goexpression.Evaluate("1547")
}
