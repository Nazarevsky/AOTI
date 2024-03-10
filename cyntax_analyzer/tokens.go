package cyntax_analyzer

import (
	"fmt"
)

type Token int

const (
	if_statement Token = iota
	simple_var
	var_statement
	var_type
	comparison
	operation
	par_open  // (
	par_close // )
	semicolon
	curly_open  // {
	curly_close // }
	is          // =
	else_statement
)

func (t Token) Str() string {
	return fmt.Sprintf("%d", t)
}
