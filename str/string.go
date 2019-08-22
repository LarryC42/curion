package str

import (
	"strings"
)

// LeftPad pads value to length using pad string
func LeftPad(value string, length int, pad string) string {
	vl := len([]rune(value))
	pl := len([]rune(pad))
	result := value + strings.Repeat(pad, (length-vl)/pl)
	if len([]rune(result)) < pl {
		result += pad
	}
	return result[:length]
}

// Text combines multiple strings
func Text(args ...string) string {
	var sb strings.Builder
	for _, arg := range args {
		sb.WriteString(arg)
	}
	return sb.String()
}

// Left returns up to the left length characters
func Left(value string, length int) string {
	e := min(length, len([]rune(value)))
	return value[:e]
}

// Right returns up to the right length characters
func Right(value string, length int) string {
	s := max(0, len([]rune(value))-length)
	return value[s:]
}

// Mid returns characters from start to end
func Mid(value string, start int) string {
	if start > len([]rune(value)) {
		return ""
	}
	return value[start:]
}

// MidLen returns characters from start up to length
func MidLen(value string, start int, length int) string {
	if start > len([]rune(value)) {
		return ""
	}
	e := min(len([]rune(value)), start+length)
	return value[start:e]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
