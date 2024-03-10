package cyntax_analyzer

import (
	"fmt"
)

func initMap() map[string][]string {
	return map[string][]string{
		key("A", if_statement): {if_statement.Str(), par_open.Str(), "E", par_close.Str(), curly_open.Str(), "F", curly_close.Str(),
			else_statement.Str(), curly_open.Str(), "F", curly_close.Str()},
		key("C", simple_var):  {simple_var.Str(), "D"},
		key("E", simple_var):  {"C", comparison.Str(), "J"},
		key("G", simple_var):  {simple_var.Str(), "I"},
		key("J", simple_var):  {simple_var.Str(), "K"},
		key("D", comparison):  {"e"},
		key("D", operation):   {operation.Str(), "C"},
		key("I", operation):   {operation.Str(), "G"},
		key("K", operation):   {operation.Str(), "J"},
		key("K", par_close):   {"e"},
		key("F", var_type):    {var_type.Str(), var_statement.Str(), is.Str(), "G", semicolon.Str(), "H"},
		key("H", var_type):    {"F"},
		key("I", semicolon):   {"e"},
		key("H", curly_close): {"e"},
	}
}

func key(nterminal string, token Token) string {
	return fmt.Sprintf("%s_%d", nterminal, token)
}
