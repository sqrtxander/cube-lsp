package utils

import (
	"cubelsp/lsp"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetRangeFromSelection(t *testing.T) {
	lines := []string{
		"this is line 0000",
		"this is line 1 that is longer than line 2",
		"this is line 2 len > 0, len < 1",
	}

    // Nothing selected, expect entire file
	ran := lsp.Range{
		Start: lsp.Position{Line: 2, Character: 3},
		End:   lsp.Position{Line: 2, Character: 3},
	}

	expectedRan := lsp.Range{
		Start: lsp.Position{Line: 0, Character: 0},
		End:   lsp.Position{Line: 2, Character: 31},
	}
	expectedBool := true

	actualRan, actualBool := GetRangeFromSelection(lines, ran)
	if !cmp.Equal(expectedRan, actualRan) {
		t.Fatalf("TestGetRangeFromSelection: expected\n%v got\n%v", expectedRan, actualRan)
	}
	if expectedBool != actualBool {
		t.Fatalf("TestGerRangeFromSelection: expected %t got %t", expectedBool, actualBool)
	}

    // going back until a space
    // going forwards until a space
	ran = lsp.Range{
		Start: lsp.Position{Line: 0, Character: 10},
		End:   lsp.Position{Line: 1, Character: 16},
	}

	expectedRan = lsp.Range{
		Start: lsp.Position{Line: 0, Character: 8},
		End:   lsp.Position{Line: 1, Character: 19},
	}
	expectedBool = false

	actualRan, actualBool = GetRangeFromSelection(lines, ran)
	if !cmp.Equal(expectedRan, actualRan) {
		t.Fatalf("TestGetRangeFromSelection: expected\n%v got\n%v", expectedRan, actualRan)
	}
	if expectedBool != actualBool {
		t.Fatalf("TestGerRangeFromSelection: expected %t got %t", expectedBool, actualBool)
	}

    // going back to the start of the line
    // going forwards until a space
	ran = lsp.Range{
		Start: lsp.Position{Line: 0, Character: 2},
		End:   lsp.Position{Line: 1, Character: 16},
	}

	expectedRan = lsp.Range{
		Start: lsp.Position{Line: 0, Character: 0},
		End:   lsp.Position{Line: 1, Character: 19},
	}
	expectedBool = false

	actualRan, actualBool = GetRangeFromSelection(lines, ran)
	if !cmp.Equal(expectedRan, actualRan) {
		t.Fatalf("TestGetRangeFromSelection: expected\n%v got\n%v", expectedRan, actualRan)
	}
	if expectedBool != actualBool {
		t.Fatalf("TestGerRangeFromSelection: expected %t got %t", expectedBool, actualBool)
	}

    // going back to the start of the line
    // going forwards to the end of the line
	ran = lsp.Range{
		Start: lsp.Position{Line: 0, Character: 2},
		End:   lsp.Position{Line: 0, Character: 14},
	}

	expectedRan = lsp.Range{
		Start: lsp.Position{Line: 0, Character: 0},
		End:   lsp.Position{Line: 0, Character: 17},
	}
	expectedBool = false

	actualRan, actualBool = GetRangeFromSelection(lines, ran)
	if !cmp.Equal(expectedRan, actualRan) {
		t.Fatalf("TestGetRangeFromSelection: expected\n%v got\n%v", expectedRan, actualRan)
	}
	if expectedBool != actualBool {
		t.Fatalf("TestGerRangeFromSelection: expected %t got %t", expectedBool, actualBool)
	}

    // an entire line selected including trailing \n
	ran = lsp.Range{
		Start: lsp.Position{Line: 2, Character: 0},
		End:   lsp.Position{Line: 2, Character: 31},
	}

	expectedRan = lsp.Range{
		Start: lsp.Position{Line: 2, Character: 0},
		End:   lsp.Position{Line: 2, Character: 31},
	}
	expectedBool = false

	actualRan, actualBool = GetRangeFromSelection(lines, ran)
	if !cmp.Equal(expectedRan, actualRan) {
		t.Fatalf("TestGetRangeFromSelection: expected\n%v got\n%v", expectedRan, actualRan)
	}
	if expectedBool != actualBool {
		t.Fatalf("TestGerRangeFromSelection: expected %t got %t", expectedBool, actualBool)
}

    // starting on a space either end
    // shrinking one char to hit non whitespace
	ran = lsp.Range{
		Start: lsp.Position{Line: 0, Character: 4},
		End:   lsp.Position{Line: 0, Character: 8},
	}

	expectedRan = lsp.Range{
		Start: lsp.Position{Line: 0, Character: 5},
		End:   lsp.Position{Line: 0, Character: 7},
	}
	expectedBool = false

	actualRan, actualBool = GetRangeFromSelection(lines, ran)
	if !cmp.Equal(expectedRan, actualRan) {
		t.Fatalf("TestGetRangeFromSelection: expected\n%v got\n%v", expectedRan, actualRan)
	}
	if expectedBool != actualBool {
		t.Fatalf("TestGerRangeFromSelection: expected %t got %t", expectedBool, actualBool)
	}

    // Shrinking multiple spaces either end
    lines = []string{
        "   this line    has many spaces    ",
        "this one too     ",
    }
	ran = lsp.Range{
		Start: lsp.Position{Line: 0, Character: 14},
		End:   lsp.Position{Line: 1, Character: 15},
	}

	expectedRan = lsp.Range{
		Start: lsp.Position{Line: 0, Character: 16},
		End:   lsp.Position{Line: 1, Character: 12},
	}
	expectedBool = false

	actualRan, actualBool = GetRangeFromSelection(lines, ran)
	if !cmp.Equal(expectedRan, actualRan) {
		t.Fatalf("TestGetRangeFromSelection: expected\n%v got\n%v", expectedRan, actualRan)
	}
	if expectedBool != actualBool {
		t.Fatalf("TestGerRangeFromSelection: expected %t got %t", expectedBool, actualBool)
	}

    // Shrinking with whole file selected
	ran = lsp.Range{
		Start: lsp.Position{Line: 0, Character: 0},
		End:   lsp.Position{Line: 1, Character: 17},
	}

	expectedRan = lsp.Range{
		Start: lsp.Position{Line: 0, Character: 3},
		End:   lsp.Position{Line: 1, Character: 12},
	}
	expectedBool = false

	actualRan, actualBool = GetRangeFromSelection(lines, ran)
	if !cmp.Equal(expectedRan, actualRan) {
		t.Fatalf("TestGetRangeFromSelection: expected\n%v got\n%v", expectedRan, actualRan)
	}
	if expectedBool != actualBool {
		t.Fatalf("TestGetRangeFromSelection: expected %t got %t", expectedBool, actualBool)
	}

    // Shrinking full file with nothing selected
	ran = lsp.Range{
		Start: lsp.Position{Line: 1, Character: 1},
		End:   lsp.Position{Line: 1, Character: 1},
	}

	expectedRan = lsp.Range{
		Start: lsp.Position{Line: 0, Character: 3},
		End:   lsp.Position{Line: 1, Character: 12},
	}
	expectedBool = true

	actualRan, actualBool = GetRangeFromSelection(lines, ran)
	if !cmp.Equal(expectedRan, actualRan) {
		t.Fatalf("TestGetRangeFromSelection: expected\n%v got\n%v", expectedRan, actualRan)
	}
	if expectedBool != actualBool {
		t.Fatalf("TestGetRangeFromSelection: expected %t got %t", expectedBool, actualBool)
	}

    // Start on a line with only whitespace
    lines = []string{
        "        ",
        " haha bro, ",
        "    ",
    }
	ran = lsp.Range{
		Start: lsp.Position{Line: 0, Character: 2},
		End:   lsp.Position{Line: 1, Character: 2},
	}

	expectedRan = lsp.Range{
		Start: lsp.Position{Line: 1, Character: 1},
		End:   lsp.Position{Line: 1, Character: 5},
	}
	expectedBool = false

	actualRan, actualBool = GetRangeFromSelection(lines, ran)
	if !cmp.Equal(expectedRan, actualRan) {
		t.Fatalf("TestGetRangeFromSelection: expected\n%v got\n%v", expectedRan, actualRan)
	}
	if expectedBool != actualBool {
		t.Fatalf("TestGetRangeFromSelection: expected %t got %t", expectedBool, actualBool)
	}

    // End on a line with only whitespace
	ran = lsp.Range{
		Start: lsp.Position{Line: 1, Character: 2},
		End:   lsp.Position{Line: 2, Character: 2},
	}

	expectedRan = lsp.Range{
		Start: lsp.Position{Line: 1, Character: 1},
		End:   lsp.Position{Line: 1, Character: 10},
	}
	expectedBool = false

	actualRan, actualBool = GetRangeFromSelection(lines, ran)
	if !cmp.Equal(expectedRan, actualRan) {
		t.Fatalf("TestGetRangeFromSelection: expected\n%v got\n%v", expectedRan, actualRan)
	}
	if expectedBool != actualBool {
		t.Fatalf("TestGetRangeFromSelection: expected %t got %t", expectedBool, actualBool)
	}

    // End and start on a line with only whitespace
	ran = lsp.Range{
		Start: lsp.Position{Line: 0, Character: 1},
		End:   lsp.Position{Line: 2, Character: 2},
	}

	expectedRan = lsp.Range{
		Start: lsp.Position{Line: 1, Character: 1},
		End:   lsp.Position{Line: 1, Character: 10},
	}
	expectedBool = false

	actualRan, actualBool = GetRangeFromSelection(lines, ran)
	if !cmp.Equal(expectedRan, actualRan) {
		t.Fatalf("TestGetRangeFromSelection: expected\n%v got\n%v", expectedRan, actualRan)
	}
	if expectedBool != actualBool {
		t.Fatalf("TestGetRangeFromSelection: expected %t got %t", expectedBool, actualBool)
	}

    // multiple lines with only whitespace adjacent
    lines = []string{
        "        ",
        "           ",
        "      ",
        "        ",
        " haha bro, ",
        "       ",
        "        ",
        "      ",
        "        ",
        "   ",
    }

    // End and start on a line with only whitespace
	ran = lsp.Range{
		Start: lsp.Position{Line: 0, Character: 1},
		End:   lsp.Position{Line: 9, Character: 2},
	}

	expectedRan = lsp.Range{
		Start: lsp.Position{Line: 4, Character: 1},
		End:   lsp.Position{Line: 4, Character: 10},
	}
	expectedBool = false

	actualRan, actualBool = GetRangeFromSelection(lines, ran)
	if !cmp.Equal(expectedRan, actualRan) {
		t.Fatalf("TestGetRangeFromSelection: expected\n%v got\n%v", expectedRan, actualRan)
	}
	if expectedBool != actualBool {
		t.Fatalf("TestGetRangeFromSelection: expected %t got %t", expectedBool, actualBool)
	}

    // pesky characters at the borders of the lines
    lines = []string{
        "        ",
        "           ",
        "     a",
        "        ",
        " haha bro, ",
        "       ",
        "        ",
        "a     ",
        "        ",
        "   ",
    }

    // End and start on a line with only whitespace
	ran = lsp.Range{
		Start: lsp.Position{Line: 0, Character: 1},
		End:   lsp.Position{Line: 9, Character: 2},
	}

	expectedRan = lsp.Range{
		Start: lsp.Position{Line: 2, Character: 5},
		End:   lsp.Position{Line: 7, Character: 1},
	}
	expectedBool = false

	actualRan, actualBool = GetRangeFromSelection(lines, ran)
	if !cmp.Equal(expectedRan, actualRan) {
		t.Fatalf("TestGetRangeFromSelection: expected\n%v got\n%v", expectedRan, actualRan)
	}
	if expectedBool != actualBool {
		t.Fatalf("TestGetRangeFromSelection: expected %t got %t", expectedBool, actualBool)
	}
}
