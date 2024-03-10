package syntax_analyzer

import (
	"lab2_merged/structures"
)

type Analyzer struct {
	matchingMap map[string][]string
	stack       structures.Stack
	input       []structures.Token
}

func New(input []structures.Token) Analyzer {
	return Analyzer{
		matchingMap: initMap(),
		stack:       structures.NewStack(),
		input:       input,
	}
}

func (a *Analyzer) IsInputCorrect() bool {
	var i = 0
	for {
		if a.stack.Len() == 0 && len(a.input)-i == 0 {
			return true
		}

		if a.stack.Len() == 0 && len(a.input) != 0 || a.stack.Len() != 0 && len(a.input) == 0 {
			return false
		}

		lastStackElem, err := a.stack.Pop()
		if err != nil {
			return false
		}

		if lastStackElem == a.input[i].Str() {
			i++
		} else if lastStackElem == "e" {
			continue
		} else {
			stackElements := a.matchingMap[key(lastStackElem, a.input[i])]
			if len(stackElements) == 0 {
				return false
			}

			a.stack.AddArray(stackElements)
		}
	}
}
