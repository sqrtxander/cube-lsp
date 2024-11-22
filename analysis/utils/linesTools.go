package utils

import (
	"cubelsp/lsp"
	"strings"
	"unicode"

	"github.com/google/go-cmp/cmp"
)

// lines is the lines in the selected range containing all of
// the first and last line including parts not in the range
func SelectionContainsNonCommentedText(lines []string, ran lsp.Range) bool {
	// "012 //abc"
	first := lines[0]
	last := lines[len(lines)-1]
	// first line check
	first = RemoveComment(first)
	if len(first) > ran.Start.Character {
		first = first[ran.Start.Character:]
		if strings.TrimSpace(first) != "" {
			return true
		}
	}
	if len(lines) == 1 {
		return false // haven't found in first line, and first line is only line
	}

	// middle line checks
	lines = lines[1 : len(lines)-1]
	for _, line := range lines {
		last = RemoveComment(last)
		if strings.TrimSpace(line) != "" {
			return true
		}
	}
	// last line check
	last = RemoveComment(last)
	last = last[:min(ran.End.Character, len(last))]
	if strings.TrimSpace(last) != "" {
		return true
	}

	return false
}

func selectionContainsAllValidMoves(lines []string, ran lsp.Range) bool {
	// TODO
	return false
}

func GetRangeFromSelection(lines []string, ran lsp.Range) (lsp.Range, bool) {
	return getRangeFromSelection(lines, ran, false)
}

func getRangeFromSelection(lines []string, ran lsp.Range, wasEqual bool) (lsp.Range, bool) {
	b := wasEqual
	if cmp.Equal(ran.Start, ran.End) {
		ran = lsp.Range{
			Start: lsp.Position{Line: 0, Character: 0},
			End:   lsp.Position{Line: len(lines) - 1, Character: len(lines[len(lines)-1])},
		}
		b = true
	}

	// extend selection backwards till previous whitespace
	first := lines[ran.Start.Line]
	charIdx := ran.Start.Character
	for charIdx > 0 && !unicode.IsSpace(rune(first[charIdx])) {
		charIdx--
	}
	// shrink selection forwards until trailing whitespace gone
	for charIdx < len(first) && unicode.IsSpace(rune(first[charIdx])) {
		charIdx++
	}
	if charIdx == len(first) { // && unicode.IsSpace(rune(first[charIdx])){
		ran.Start = lsp.Position{Line: ran.Start.Line + 1, Character: 0}
		return getRangeFromSelection(lines, ran, b)
	}
	ran.Start.Character = charIdx

	// extend selection forwards until next whitespace
	last := lines[ran.End.Line]
	charIdx = ran.End.Character - 1
	for charIdx < len(last) && !unicode.IsSpace(rune(last[charIdx])) {
		charIdx++
	}
	// shrink selection backwards until trailing whitespace gone
	for charIdx > 0 && unicode.IsSpace(rune(last[charIdx-1])) {
		charIdx--
	}
	if charIdx == 0 {
		ran.End = lsp.Position{Line: ran.End.Line - 1, Character: len(lines[ran.End.Line-1])}
		return getRangeFromSelection(lines, ran, b)
	}
	ran.End.Character = charIdx
	return ran, b
}
