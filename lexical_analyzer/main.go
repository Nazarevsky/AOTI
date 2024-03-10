package lexical_analyzer

import (
	"fmt"
	"lab2_merged/structures"
	"regexp"
	"strings"
)

var tokenPatterns = []struct {
	pattern string
	token   structures.Token
}{
	{`^if`, structures.If_statement},
	{`^else`, structures.Else_statement},
	{`^(int|float)`, structures.Var_type},
	{`^[a-zA-Z_]\w*`, structures.Var_statement},
	{`^\d+(\.\d+)?([Еe][-]?\d+)?`, structures.Simple_var},
	{`^[-+*/%]`, structures.Operation},
	{`^(==|>=|<=|!=|>|<)`, structures.Comparison},
	{`^=`, structures.Is},
	{`^;`, structures.Semicolon},
	{`^\(`, structures.Par_open},
	{`^\)`, structures.Par_close},
	{`^{`, structures.Curly_open},
	{`^}`, structures.Curly_close},
}

func Analyze(input string) []structures.Token {
	tokens := []structures.Token{}
	inputLines := strings.Split(input, "\n")

	for i := 0; i < len(inputLines); i++ {
		line := strings.TrimSpace(inputLines[i])

		for len(line) > 0 {
			isMatch := false
			for _, tokenPattern := range tokenPatterns {
				regex := regexp.MustCompile(tokenPattern.pattern)
				match := regex.FindString(line)
				if match != "" {
					isMatch = true
					tokens = append(tokens, tokenPattern.token)
					line = strings.TrimSpace(line[len(match):])
					break
				}
			}
			if !isMatch {
				panic(fmt.Sprintf(`Unknown symbol "%s" at line %d`, line[:1], i))
			}
		}
	}

	return tokens
}
