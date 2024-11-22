package diagnostic

import "cubelsp/lsp"

type Strategy interface {
	getDiagnostics(int, string) ([]lsp.Diagnostic, error)
}
