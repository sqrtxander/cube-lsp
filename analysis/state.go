package analysis

import (
	"cubelsp/analysis/codeaction"
	"cubelsp/analysis/diagnostic"
	"cubelsp/analysis/hover"
	"cubelsp/analysis/utils"
	"cubelsp/lsp"
	"strings"
)

type State struct {
	// Maps of file names to contents, scrambles
	Documents map[string]string
	Scrambles map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func getDiagnosticsForFile(text string) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}
	diagnosticGetter := diagnostic.NewDiagnosticGetter()
	for row, line := range strings.Split(text, "\n") {
		line = utils.RemoveComment(line)
		newDiagnostics, _ := diagnosticGetter.GetDiagnostics(row, line)
		diagnostics = append(diagnostics, newDiagnostics...)
	}
	return diagnostics
}

func (s *State) OpenDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return getDiagnosticsForFile(text)
}

func (s *State) UpdateDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return getDiagnosticsForFile(text)
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	text := s.Documents[uri]
	result := hover.GetHoverResult(text, position)

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: result,
	}
}

func (s *State) TextDocumentCodeAction(id int, params lsp.TextDocumentCodeActionParams) lsp.TextDocumentCodeActionResponse {
	uri := params.TextDocument.URI
	ran := params.Range
	text := s.Documents[uri]

	actions := []lsp.CodeAction{}
	actions = append(actions, codeaction.GetReplaceActions(text, uri, ran)...)

	response := lsp.TextDocumentCodeActionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: actions,
	}

	return response
}

func (s *State) TextDocumentCompletion(id int, uri string) lsp.CompletionResponse {
	items := []lsp.CompletionItem{}

	response := lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}

	return response
}
