package main

import (
	"fmt"
	"regexp"
	"strings"
)

var tokenPatterns = []struct {
	pattern string
	token   string
}{
	{`^if`, "if_statement"},
	{`^else`, "else_statement"},
	{`^(int|float)`, "var_type"},
	{`^[a-zA-Z_]\w*\s*=`, "var_statement"},
	{`^[a-zA-Z_]\w*|^\d+(\.\d+)?([Еe][-]?\d+)?`, "simple_var"},
	{`^[-+*/%]`, "operation"},
	{`^(==|>=|<=|!=|>|<)`, "comparison"},
	{`^=`, "is"},
	{`^;`, "semicolon"},
	{`^\(`, "par_open"},
	{`^\)`, "par_close"},
	{`^{`, "curly_open"},
	{`^}`, "curly_close"},
}

func tokenize(input string) []string {
	tokens := []string{}
	inputLines := strings.Split(input, "\n")

	for i := 0; i < len(inputLines); i++ {
		line := strings.TrimSpace(inputLines[i])
		lineTokens := []string{}

		for len(line) > 0 {
			isMatch := false
			for _, tokenPattern := range tokenPatterns {
				regex := regexp.MustCompile(tokenPattern.pattern)
				match := regex.FindString(line)
				if len(match) > 0 {
					isMatch = true
					lineTokens = append(lineTokens, tokenPattern.token)
					if lineTokens[len(lineTokens)-1] == "var_statement" {
						line = strings.TrimSpace(line[len(match)-1:])
					} else {
						line = strings.TrimSpace(line[len(match):])
					}
					break
				}
			}
			if !isMatch {
				panic(fmt.Sprintf(`Unknown symbol "%s" at line %d`, line[:1], i))
			}
		}

		if len(lineTokens) > 0 {
			tokens = append(tokens, strings.Join(lineTokens, ", "))
		}
	}

	return tokens
}

func printTokens(tokens []string) {
	for i := 0; i < len(tokens); i++ {
		fmt.Println(tokens[i])
	}
}

func main() {
	input := `int x = 10;
    int y = 5;
	float z = x / y;
    int result;

    if (z >= 7.6) {
        result = x + y;
    } else {
        result = x - y;
    }`

	tokens := tokenize(input)
	printTokens(tokens)
}
