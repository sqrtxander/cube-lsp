package diagnostic

import (
	"cubelsp/analysis/globs"
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
	globs.SCRAMBLE,
	globs.TIME,
}

func (s *PreambleStrat) getDiagnostics(row int, line string) ([]lsp.Diagnostic, error) {
	if !utils.IsKey(line) {
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
			Message:  fmt.Sprintf("Invalid key `%s`", key),
		})
	} else if key == globs.SCRAMBLE {
		diagnostics = append(diagnostics, validateNotation(row, value, len(globs.SCRAMBLE)+1)...)
	}
	return diagnostics, nil
}
