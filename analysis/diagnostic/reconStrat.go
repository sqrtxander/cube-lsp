package diagnostic

import (
	"cubelsp/lsp"
)

type ReconStrat struct {
	diagnosticGetter *DiagnosticGetter
}

func (s *ReconStrat) getDiagnostics(row int, line string) ([]lsp.Diagnostic, error) {
	return validateNotation(row, line, 0), nil
}
