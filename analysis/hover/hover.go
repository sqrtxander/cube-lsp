package hover

import (
	"cubelsp/analysis/utils"
	"cubelsp/cube"
	"cubelsp/lsp"
	"fmt"
	"strings"
	"unicode"
)

func getModifiedPosition(lines []string, position lsp.Position) lsp.Position {
	last := lines[position.Line]
	lastIdx := position.Character
	for lastIdx < len(last) && !unicode.IsSpace(rune(last[lastIdx])) {
		lastIdx++
	}
	return lsp.Position{Line: position.Line, Character: lastIdx}
}

func GetHoverResult(text string, position lsp.Position) lsp.HoverResult {
	lines := strings.Split(text, "\n")
	position = getModifiedPosition(lines, position)
	lines = lines[:position.Line+1]
	lines[position.Line] = lines[position.Line][:position.Character]

	moves := ""
	for _, line := range lines {
		line = utils.RemoveComment(line)
		if utils.IsKeyOtherThan(line, "scramble") {
			continue
		}
		line = utils.RemoveKey(line)
		moves += line + " "
	}

	cubeState, err := cube.DoMoves(*cube.GetSolvedCube(), moves)
	contents := ""
	if err == nil {
		contents = fmt.Sprintf("State at cursor:\n%s",
			cube.ToFatString(*cubeState, cube.ToNetString))
	}

	return lsp.HoverResult{
		Contents: contents,
	}
}
