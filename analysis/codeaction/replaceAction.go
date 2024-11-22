package codeaction

import (
	"cubelsp/analysis/globs"
	"cubelsp/analysis/utils"
	"cubelsp/lsp"
	"strings"
)

func GetReplaceActions(text string, uri string, ran lsp.Range) []lsp.CodeAction {
	text = strings.TrimRight(text, "\n")
	lines := strings.Split(text, "\n")
	ran, isFull := utils.GetRangeFromSelection(lines, ran)
	titlePost := " selection"
	if isFull {
		titlePost = " file"
	}
	lines = lines[ran.Start.Line : ran.End.Line+1]
	if !utils.SelectionContainsNonCommentedText(lines, ran) {
		return []lsp.CodeAction{}
	}
	if len(lines) == 1 {
		lines[0] = lines[0][ran.Start.Character:ran.End.Character]
	} else {
		lines[0] = lines[0][ran.Start.Character:]
		lines[len(lines)-1] = lines[len(lines)-1][:ran.End.Character]
	}

	actions := []lsp.CodeAction{}
	actions = append(actions, replaceAction(lines, ran, uri, "M-Mirror"+titlePost, Mmirrmap))
	actions = append(actions, replaceAction(lines, ran, uri, "E-Mirror"+titlePost, Emirrmap))
	actions = append(actions, replaceAction(lines, ran, uri, "S-Mirror"+titlePost, Smirrmap))
	return actions
}

func replaceAction(lines []string, ran lsp.Range, uri string, title string, method map[string]string) lsp.CodeAction {
	replaceChange := map[string][]lsp.TextEdit{}
	replaceChange[uri] = []lsp.TextEdit{}
	for row, line := range lines {
		if utils.IsKeyOtherThan(line, globs.SCRAMBLE) {
			continue
		}
		keySepIdx := utils.GetKeySepNextIndexOrStart(line)
		commIdx := utils.GetCommentIndexOrEnd(line)

		newRan := utils.LineRange(row+ran.Start.Line, keySepIdx, commIdx)
		line = utils.RemoveComment(line)
		line = utils.RemoveKey(line)
		if row == 0 {
			newRan.Start.Character += ran.Start.Character
			newRan.End.Character += ran.Start.Character
		}

		newText := replaceLine(line, method)
		replaceChange[uri] = append(replaceChange[uri], lsp.TextEdit{
			Range: newRan, NewText: newText,
		})
	}
	return lsp.CodeAction{
		Title: title,
		Edit: &lsp.WorkspaceEdit{
			Changes: replaceChange,
		},
	}
}

func replaceLine(line string, method map[string]string) string {
	newLine := ""
	for _, move := range strings.Split(line, " ") {
		if move != " " {
			newLine += method[move]
		}
		newLine += " "
	}
	return newLine[:len(newLine)-1]
}

func allMovesValid(lines []string) bool {
	for _, line := range lines {
		line = utils.RemoveComment(line)
	}
	return true
}
