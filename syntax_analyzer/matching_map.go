package syntax_analyzer

import (
	"fmt"
	"lab2_merged/structures"
)

func initMap() map[string][]string {
	return map[string][]string{
		key("A", structures.If_statement): {structures.If_statement.Str(), structures.Par_open.Str(), "E",
			structures.Par_close.Str(), structures.Curly_open.Str(), "F", structures.Curly_close.Str(),
			structures.Else_statement.Str(), structures.Curly_open.Str(), "F", structures.Curly_close.Str()},
		key("C", structures.Simple_var): {structures.Simple_var.Str(), "D"},
		key("E", structures.Simple_var): {"C", structures.Comparison.Str(), "J"},
		key("G", structures.Simple_var): {structures.Simple_var.Str(), "I"},
		key("J", structures.Simple_var): {structures.Simple_var.Str(), "K"},
		key("D", structures.Comparison): {"e"},
		key("D", structures.Operation):  {structures.Operation.Str(), "C"},
		key("I", structures.Operation):  {structures.Operation.Str(), "G"},
		key("K", structures.Operation):  {structures.Operation.Str(), "J"},
		key("K", structures.Par_close):  {"e"},
		key("F", structures.Var_type): {structures.Var_type.Str(), structures.Var_statement.Str(),
			structures.Is.Str(), "G", structures.Semicolon.Str(), "H"},
		key("H", structures.Var_type):    {"F"},
		key("I", structures.Semicolon):   {"e"},
		key("H", structures.Curly_close): {"e"},
	}
}

func key(nterminal string, token structures.Token) string {
	return fmt.Sprintf("%s_%d", nterminal, token)
}
