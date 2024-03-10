package structures

import (
	"fmt"
)

type Token int

const (
	If_statement Token = iota
	Simple_var
	Var_statement
	Var_type
	Comparison
	Operation
	Par_open  // (
	Par_close // )
	Semicolon
	Curly_open  // {
	Curly_close // }
	Is          // =
	Else_statement
)

func (t Token) Str() string {
	return fmt.Sprintf("%d", t)
}

func (t Token) ToWord() string {
	switch t {
	case If_statement:
		return "if"
	case Simple_var:
		return "simple_var"
	case Var_statement:
		return "var"
	case Var_type:
		return "var_type"
	case Comparison:
		return "comparison"
	case Operation:
		return "operation"
	case Par_open:
		return "("
	case Par_close:
		return ")"
	case Semicolon:
		return ";"
	case Curly_open:
		return "{"
	case Curly_close:
		return "}"
	case Is:
		return "="
	case Else_statement:
		return "else"
	}
	return "UNKNOWN"
}

func PrintTokens(tokens []Token) {
	for i := 0; i < len(tokens); i++ {
		fmt.Printf("%s ", tokens[i].ToWord())
	}
	println()
}
