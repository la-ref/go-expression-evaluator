<h1 align="center"><b>GO Expression Evaluator</b></h1>
<div align="center">
  <a href="https://golang.org/dl" target="_blank">
    <img alt="GO" src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" />
  </a>
  <a href="./LICENSE" target="./LICENSE">
    <img alt="MIT" src="https://camo.githubusercontent.com/3dbcfa4997505c80ef928681b291d33ecfac2dabf563eb742bb3e269a5af909c/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f6c6963656e73652f496c65726961796f2f6d61726b646f776e2d6261646765733f7374796c653d666f722d7468652d6261646765" />
  </a>
    <p></p>A Simple Expression Evaluator for Arithmetic Operations written in Go</p>
</div>

## **Introduction**

This project is a simple expression evaluator written in Go. The evaluator can handle basic arithmetic operations (+, -, *, /), parentheses (()), and unary operators. The main components of the project are:

1. **Tokenizer:** The tokenizer is responsible for converting the input expression into a stream of tokens. Each token represents a specific element, such as a number, operator, or parenthesis.

2. **AST (Abstract Syntax Tree):** After tokenization, the tokens are used to construct an Abstract Syntax Tree (AST). The AST is a hierarchical representation of the expression that makes it easier to evaluate the expression later.

3. **DFS Evaluation:** The evaluator uses a Depth-First Search (DFS) algorithm to traverse the AST and perform the expression evaluation. This process considers operator precedence and handles nested expressions efficiently.

I made this project to discover the world of expression evalution.

## **Usage**

### How to Use

You can utilize the expression evaluator by calling the `Evaluate` function with your expression as a string argument. The function will return the result of the expression evaluation as a `float64`.

```go
package main

import (
	"fmt"
	evaluator "github.com/la-ref/go-expression-evaluator/goexpression"
)

func main() {
	expression := "(2 + 3) * 4 - 6 / 2"
	result, err := evaluator.Evaluate(expression)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result) // Result: 17
}
```

### Examples
Here are some example expressions that the evaluator can handle:

- Basic arithmetic: **`2 + 3 * 4 - 6 / 2`**
- Parentheses: **`(2 + 3) * 4`**
- Unary operators: **`-3 + 5`**

## **Getting Started**
To use the expression evaluator in your project, follow these steps:

1. Clone this repository to your local machine:
    ```bash
    git clone https://github.com/la-ref/go-expression-evaluator
    cd go-expression-evaluator
    ```
    OR
    ```
    go get -u github.com/la-ref/goexpression
    ```
2. Implement the evaluator package in your Go code. Use the provided Evaluate function to evaluate expressions.
3. Customize the evaluator to fit your project requirements. You can add more functionalities or optimize the code as needed.

## **Contribution**
Contributions are welcome! If you find any issues or want to add new features, feel free to submit a pull request. Please make sure to follow the Go coding conventions.

## **Support Me**
Give me a ⭐ if this project was helpful in any way!

## **Licence**
This project is licensed under the MIT License. You are free to use, modify, and distribute this code following the terms of the license

