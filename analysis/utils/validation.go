package utils

import (
	"fmt"
	"slices"
)

var ValidFaces = []byte{
	'U', 'D', 'L', 'R', 'F', 'B',
	'u', 'd', 'l', 'r', 'f', 'b',
	'M', 'E', 'S', 'x', 'y', 'z',
}

var ValidAmounts = []byte{'2', '\''}

// IsValidMove determines if a move is valid. The move passed in must
// be a string with no spaces for the returned string to make sense.
//
// returns a boolean whether the move is valid and a string as to why
// the move is invalid.
func IsValidMove(move string) (bool, string) {
    if len(move) == 0 {
        return false, "This is unexpected, should never occur with this lsp"
    }

    if len(move) > 2 {
        return false, fmt.Sprintf("Move '%s' too long", move)
    }

    if !slices.Contains(ValidFaces, move[0]) {
        return false, fmt.Sprintf("Invalid face in move '%s'", move)
    }

    if len(move) == 2 && !slices.Contains(ValidAmounts, move[1]) {
        return false, fmt.Sprintf("Invalid amount in move '%s'", move)
    }

    return true, ""
}
