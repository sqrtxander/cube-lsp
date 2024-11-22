package utils

import (
	"cubelsp/lsp"
)

func LineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Character: start,
		},
		End: lsp.Position{
			Line:      line,
			Character: end,
		},
	}
}

func inStrBounds(str string, idx int) bool {
    return idx >= 0 && idx < len(str)
}
