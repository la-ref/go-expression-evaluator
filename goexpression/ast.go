package goexpression

import (
	"encoding/json"
	"fmt"
)

type ASTType int

const (
	OPERATOR = iota
	VALUE
)

type AST struct {
	Left  *AST    `json:"left"`
	Right *AST    `json:"right"`
	T     ASTType `json:"type"`
	Value any     `json:"value"`
}

func (a *AST) AddRight(r *AST) {
	a.Right = r
}

func (a *AST) AddLeft(l *AST) {
	a.Left = l
}

func (a *AST) Print() {
	b, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(string(b))
}
