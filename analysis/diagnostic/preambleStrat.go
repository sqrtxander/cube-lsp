package diagnostic

import (
	"cubelsp/analysis/utils"
	"cubelsp/lsp"
	"fmt"
	"slices"
	"strings"
)

type PreambleStrat struct {
	diagnosticGetter *DiagnosticGetter
}

var validKeys = []string{
	"scramble",
	"time",
}


func (s *PreambleStrat) getDiagnostics(row int, line string) ([]lsp.Diagnostic, error) {
	if !strings.Contains(line, ":") {
		s.diagnosticGetter.setStrategy(
			&ReconStrat{diagnosticGetter: s.diagnosticGetter},
		)
        return s.diagnosticGetter.GetDiagnostics(row, line)
	}
	diagnostics := []lsp.Diagnostic{}
	key, value, _ := strings.Cut(line, ":")
	if !slices.Contains(validKeys, key) {
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Range:    utils.LineRange(row, 0, len(key)),
			Severity: 1,
			Source:   "cubelsp",
			Message:  fmt.Sprintf("Unknown key `%s`", key),
		})
	} else if key == "scramble" {
        diagnostics = append(diagnostics, validateNotation(row, value, len("scramble:"))...);
    }
	return diagnostics, nil
}
