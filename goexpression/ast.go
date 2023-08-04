package goexpression

import (
	"encoding/json"
	"fmt"
)

type ASTType int

func (t ASTType) MarshalJSON() ([]byte, error) {
	names := [2]string{"OPERATOR", "VALUE"}

	if t < 0 || t >= ASTType(len(names)) {
		return nil, fmt.Errorf("Invalid ASTType value: %d", t)
	}

	return json.Marshal(names[t])
}

const (
	OPERATOR ASTType = iota
	VALUE
)

type AST struct {
	T     ASTType `json:"type"`
	Value any     `json:"value"`
	Left  *AST    `json:"left"`
	Right *AST    `json:"right"`
}

func (a *AST) AddRight(r *AST) {
	a.Right = r
}

func (a *AST) AddLeft(l *AST) {
	a.Left = l
}

func (a *AST) Print() {
	b, err := json.MarshalIndent(a, "", " ")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(string(b))
}
