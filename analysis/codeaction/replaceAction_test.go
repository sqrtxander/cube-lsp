package codeaction

import (
	"cubelsp/lsp"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetReplaceActions(t *testing.T) {
	text := `L' D' L D// comment :)
L' U' L U // comment 2
M E' M' E // final comment
`
	uri := "foo"

	ran := lsp.Range{
		Start: lsp.Position{Line: 1, Character: 3},
		End:   lsp.Position{Line: 1, Character: 3},
	}

	expectedEditChangesUri := []lsp.TextEdit{
		{
			Range:   lsp.Range{Start: lsp.Position{Line: 0, Character: 0}, End: lsp.Position{Line: 0, Character: 9}},
			NewText: "R D R' D'",
		},
		{
			Range:   lsp.Range{Start: lsp.Position{Line: 1, Character: 0}, End: lsp.Position{Line: 1, Character: 10}},
			NewText: "R U R' U' ",
		},
		{
			Range:   lsp.Range{Start: lsp.Position{Line: 2, Character: 0}, End: lsp.Position{Line: 2, Character: 10}},
			NewText: "M E M' E' ",
		},
	}
	actualEditChangesUri := GetReplaceActions(text, uri, ran)[0].Edit.Changes[uri]

    if !cmp.Equal(expectedEditChangesUri, actualEditChangesUri) {
        t.Fatalf("TestGetReplaceActions: Got\n%v but expected\n%v", actualEditChangesUri, expectedEditChangesUri)
    }

	ran = lsp.Range{
		Start: lsp.Position{Line: 1, Character: 3},
		End:   lsp.Position{Line: 1, Character: 4},
	}

	expectedEditChangesUri = []lsp.TextEdit{
		{
			Range:   lsp.Range{Start: lsp.Position{Line: 1, Character: 3}, End: lsp.Position{Line: 1, Character: 5}},
			NewText: "U",
		},
	}
	actualEditChangesUri = GetReplaceActions(text, uri, ran)[0].Edit.Changes[uri]

    if !cmp.Equal(expectedEditChangesUri, actualEditChangesUri) {
        t.Fatalf("TestGetReplaceActions: Got\n%v but expected\n%v", actualEditChangesUri, expectedEditChangesUri)
    }

	ran = lsp.Range{
		Start: lsp.Position{Line: 0, Character: 4},
		End:   lsp.Position{Line: 2, Character: 5},
	}

	expectedEditChangesUri = []lsp.TextEdit{
		{
			Range:   lsp.Range{Start: lsp.Position{Line: 0, Character: 3}, End: lsp.Position{Line: 0, Character: 9}},
			NewText: "D R' D'",
		},
		{
			Range:   lsp.Range{Start: lsp.Position{Line: 1, Character: 0}, End: lsp.Position{Line: 1, Character: 10}},
			NewText: "R U R' U' ",
		},
		{
			Range:   lsp.Range{Start: lsp.Position{Line: 2, Character: 0}, End: lsp.Position{Line: 2, Character: 4}},
			NewText: "M E",
		},
	}
	actualEditChangesUri = GetReplaceActions(text, uri, ran)[0].Edit.Changes[uri]

    if !cmp.Equal(expectedEditChangesUri, actualEditChangesUri) {
        t.Fatalf("TestGetReplaceActions: Got\n%v but expected\n%v", actualEditChangesUri, expectedEditChangesUri)
    }
}
