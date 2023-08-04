package main

import (
	"fmt"
	"goexpressionevaluator/goexpression"
	"goexpressionevaluator/goexpression/utils"
)

func main() {
	fmt.Println(utils.IsCharANumber("-"))
	goexpression.Evaluate("15+14*21+(41-2)")
}
