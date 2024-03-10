package main

import (
	"fmt"
	"lab2_merged/lexical_analyzer"
	"lab2_merged/reader"
	"lab2_merged/syntax_analyzer"
)

func getCorrectnessWord(isCorrect bool) string {
	if isCorrect {
		return "correct"
	}
	return "not correct"
}

func main() {
	fileText, err := reader.ReadBinaryFile("test.fakec")
	if err != nil {
		fmt.Println(err)
		return
	}

	tokens := lexical_analyzer.Analyze(fileText)

	cynAnalyzer := syntax_analyzer.New(tokens)
	fmt.Printf("Syntax of the provided code is %s\n", getCorrectnessWord(cynAnalyzer.IsInputCorrect()))
}
