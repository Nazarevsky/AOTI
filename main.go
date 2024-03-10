package main

import "lab2/cyntax_analyzer"

const (
	if_statement = iota
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

func main() {
	/*
		if (x + 1 == 1) {
			int b = 1 + 2 / 23;
			b = 1;
		} else {
			int z = 2;
			int z = 2;
		}
	*/
	input := []cyntax_analyzer.Token{
		if_statement, par_open, simple_var, comparison, simple_var, operation, simple_var, par_close, curly_open,
		var_type, var_statement, is, simple_var, simple_var, semicolon,
		curly_close, else_statement, curly_open,
		var_type, var_statement, is, simple_var, semicolon,
		var_type, var_statement, is, simple_var, semicolon,
		curly_close,
	}

	analyzer := cyntax_analyzer.New(input)
	println(analyzer.IsInputCorrect())
}
