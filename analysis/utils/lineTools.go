package utils

import (
	"cubelsp/analysis/globs"
	"strings"
)

func RemoveComment(line string) string {
	line, _, _ = strings.Cut(line, globs.COMMENTSTRING)
	return line
}

func RemoveKey(line string) string {
	first, second, hasKey := strings.Cut(line, globs.PREAMBLESEP)
	if hasKey {
		return second
	}
	return first
}

func GetCommentIndexOrEnd(line string) int {
	idx := strings.Index(line, globs.COMMENTSTRING)
	if idx == -1 {
		return len(line)
	}
	return idx
}

func GetKeySepNextIndexOrStart(line string) int {
	line = RemoveComment(line)
	return strings.Index(line, globs.PREAMBLESEP) + 1
}

func IsKeyOf(line string, key string) bool {
	return IsKey(line) && strings.HasPrefix(line, key+globs.PREAMBLESEP)
}

func IsKey(line string) bool {
	line = RemoveComment(line)                       // ensure not checking for separators in comments
	return strings.Contains(line, globs.PREAMBLESEP) // if the line contains a key separator
}

func IsKeyOtherThan(line string, key string) bool {
	return IsKey(line) && !strings.HasPrefix(line, key+globs.PREAMBLESEP)
}
