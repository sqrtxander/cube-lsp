package diagnostic

import (
	"cubelsp/analysis/utils"
	"cubelsp/lsp"
	"strings"
)

func validateNotation(row int, text string, padding int) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}
	for _, move := range strings.Fields(strings.TrimSpace(text)) {
		valid, message := utils.IsValidMove(move)
		if !valid {
			idx := strings.Index(text, move) + padding
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    utils.LineRange(row, idx, idx+len(move)),
				Severity: 1,
				Source:   "cubelsp",
				Message:  message,
			})
		}
	}
	return diagnostics
}
