package diagnostic

import "cubelsp/lsp"

type DiagnosticGetter struct {
	currentState Strategy
}

func NewDiagnosticGetter() *DiagnosticGetter {
	dg := &DiagnosticGetter{}
	current := &PreambleStrat{diagnosticGetter: dg}
    dg.currentState = current
	return dg
}

func (dg *DiagnosticGetter) GetDiagnostics(row int, line string) ([]lsp.Diagnostic, error) {
	return dg.currentState.getDiagnostics(row, line)
}

func (dg *DiagnosticGetter) setStrategy(s Strategy) {
	dg.currentState = s
}
