package analysis

import (
	"cubelsp/analysis/codeaction"
	"cubelsp/analysis/diagnostic"
	"cubelsp/analysis/utils"
	"cubelsp/cube"
	"cubelsp/lsp"
	"fmt"
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
	document := s.Documents[uri]

	documentLines := strings.Split(document, "\n")[:position.Line+1]
	lastIdx := 0
	lastLine := documentLines[position.Line]
	if lastLine != "" {
		lastLine = lastLine[position.Character:] + " "
		lastIdx = position.Character + max(strings.Index(lastLine, " "), 0)
	}
	documentLines[position.Line] = documentLines[position.Line][:lastIdx]

	moves := ""
	for _, line := range documentLines {
        line = utils.RemoveComment(line)
        if utils.IsKeyOtherThan(line, "scramble") {
			continue
		}
        line = utils.RemoveKey(line)
		moves += line + " "
	}

	cubeState, err := cube.DoMoves(*cube.GetSolvedCube(), moves)
	contents := ""
	if err == nil {
		contents = fmt.Sprintf("State at cursor:\n%s", cube.ToFatString(*cubeState, cube.ToNetString))
	}

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: contents,
		},
	}
}

func (s *State) TextDocumentCodeAction(id int, params lsp.TextDocumentCodeActionParams) lsp.TextDocumentCodeActionResponse {
    uri := params.TextDocument.URI
    ran := params.Range
	text := s.Documents[uri]
	_ = text

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
