package utils

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var ValidFaces = []byte{
	'U', 'D', 'L', 'R', 'F', 'B',
	'u', 'd', 'l', 'r', 'f', 'b',
	'M', 'E', 'S', 'x', 'y', 'z',
}

var ValidAmounts = []byte{'2', '\''}

// IsValidMove determines if a move is valid. The move passed in must
// be a string with no whitespace for the returned string to make sense.
//
// returns a boolean whether the move is valid and a string as to why
// the move is invalid.
func IsValidMove(move string) (bool, string) {
	// valid e.g., R2, M', U3, U2', U42
	// invalid e.g., R'2, P, F0

	if len(move) == 0 {
		return false, "If you are seeing this, an error occurred because this should be unreachable"
	}

	if !slices.Contains(ValidFaces, move[0]) {
		return false, fmt.Sprintf("Invalid face in move '%s'", move)
	}

	if len(move) == 1 {
		return true, ""
	}

	amountStr := move[1:]
	if strings.HasSuffix(amountStr, "'") {
		amountStr = amountStr[:len(amountStr)-1]
	}

	if amountStr == "" {
		return true, ""
	}

	amount, err := strconv.Atoi(amountStr)
	if err != nil || amount <= 0 {
		return false, fmt.Sprintf("Invalid amount in move '%s'", move)
	}

	return true, ""
}
