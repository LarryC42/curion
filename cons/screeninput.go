package cons

import (
	"fmt"
	"lib/dt"
	"lib/str"
	"strings"
)

//==============================================================================
// Section: Screen entry
//==============================================================================

// /**
// definition of the location, size, prompt, and value to use for a entry field.
// **/

// InputField describes a field used for full screen data entry
type InputField struct {
	row           int
	col           int
	size          int
	prompt        string
	value         string
	validateKey   func(field *InputField, key KeyEvent) bool
	validateField func(field *InputField) bool
}

// CreateInputField creates a fully populated input field
func CreateInputField(row int, col int, prompt string, value string, size int, field *InputField, validateKey func(field *InputField, key KeyEvent) bool, validateField func(field *InputField) bool) InputField {
	return InputField{size: size,
		prompt:        prompt,
		value:         value,
		row:           row,
		col:           col,
		validateKey:   validateKey,
		validateField: validateField,
	}
}

// NewInputField creates an input field with prompt, initialial value, and max size
func NewInputField(prompt string, value string, size int) InputField {
	return InputField{size: size, prompt: prompt, value: value}
}

// ValidatedInputField adds key and field validation
func ValidatedInputField(field *InputField, validateKey func(field *InputField, key KeyEvent) bool, validateField func(field *InputField) bool) {
	field.validateKey = validateKey
	field.validateField = validateField
}

// PositionInputField sets the row/col for the input field
func PositionInputField(field *InputField, row int, col int) {
	field.row = row
	field.col = col
}

// FieldValue returns the entered field value
func FieldValue(field *InputField) string {
	return field.value
}

// Entry simplifies full screen data entry.  Calculates row,col positions, draws screen, does input
func Entry(title string, subTitle string, fields []InputField, foreground int8,
	background int8, fieldForeground int8, fieldBackground int8, borderStyle int) bool {
	const (
		topLeft            = 0
		topRight           = 1
		bottomLeft         = 2
		bottomRight        = 3
		topIntersection    = 4
		bottomIntersection = 5
		horizontalLine     = 6
		verticalLine       = 7
		leftIntersection   = 8
		rightIntersection  = 9
	)

	SetColor(foreground, background)
	Cls()
	if len(title) > 0 {
		Center(title)
	}
	if len(subTitle) > 0 {
		Center(subTitle)
	}
	Center(dt.Dtols(dt.Today()))
	fmt.Println()
	promptLength := 0
	valueLength := 0
	for index := range fields {
		promptLength = max(promptLength, len([]rune(fields[index].prompt)))
		valueLength = max(valueLength, fields[index].size)
	}
	for index := range fields {
		fields[index].size = min(fields[index].size, Cols()-(promptLength+2))
	}

	singleLine := []string{"┌", "┐", "└", "┘", "┬", "┴", "─", "│", "├", "┤"}
	doubleLine := []string{"╔", "╗", "╚", "╝", "╦", "╩", "═", "║", "╠", "╣"}
	noLine := []string{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "}
	var lineChar []string
	switch borderStyle {
	case LineStyleSingle:
		lineChar = singleLine
	case LineStyleDouble:
		lineChar = doubleLine
	case LineStyleNone:
		lineChar = noLine
	}

	tl := lineChar[topLeft]
	tr := lineChar[topRight]
	bl := lineChar[bottomLeft]
	br := lineChar[bottomRight]
	vl := lineChar[verticalLine]

	width := 2 + promptLength + 2 + valueLength + 2
	var line string
	divider := strings.Repeat(lineChar[horizontalLine], width-2)

	line = tl + divider + tr
	Center(line)

	row := Row()
	col := ((80 - width) / 2) + 2 + promptLength + 2 // '| ', prompt, ': '
	//int pcol = (80 - width + 1) / 2;
	for index := range fields {
		fields[index].row = row
		fields[index].col = col
		p := str.LeftPad(fields[index].prompt+":", promptLength+1, ".") + " "
		Center(fmt.Sprint(vl, " ", str.LeftPad(p, promptLength+valueLength+3, " "), vl))
		row++
	}
	Center(fmt.Sprint(bl, divider, br))
	bottom := Row()

	PaintFields(fields, fieldForeground, fieldBackground)
	ret := StartEntry(fields)
	Locate(bottom, 0)
	return ret
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

// PaintFields draws all field contents using foreground/background colors
func PaintFields(fields []InputField, foreground int8, background int8) {
	SetColor(foreground, background)
	for i := range fields {
		Locate(fields[i].row, fields[i].col)
		fmt.Print(str.LeftPad(fields[i].value, fields[i].size, " "))
	}
}

// StartEntry performs full screen entry.  It is assumed you have already updated the screen in preparation.
// Params:
// 	fields = The field locations, sizes, and values to capture input for.
// Enter, down arrow, and tab moves to the next field.  On the last field, Enter will exit while down arrow and tab will move to the first field.
// up arrow and shift tab will move to the prior field.  When on the first field, they will move to the last field.
// home moves to the begining of the line.
// end moves to the end of the entered line.
// left arrow moves left a character
// right arrow moves right a character up to the end of line
// control+left arrow moves left a word
// control+right arrow moves right a word.
// insert inserts a space
// delete deletes the current character
// control+delete deletes to the end of line
// backspace delets the character to the left of the cursor and moves left one character.
// f10 or control+Enter will exit entry with success.
// escape will exit entry with failure.
// typing a character will change the current character and advance the cursor.
func StartEntry(fields []InputField) bool {
	currentField := 0
	offset := 0
	var ch KeyEvent
	if len(fields) < 1 {
		return false
	}

	for {
		field := &(fields[currentField])

		Locate(field.row, field.col+offset)
		ch = GetKey()
		switch ch.Key {
		case KeyEnter, KeyDown:
			if !field.validateField(field) {
				Beep()
				break
			}
			currentField++
			if currentField >= len(fields) {
				if ch.Key == KeyEnter {
					for index := range fields {
						fields[index].value = strings.TrimRight(fields[index].value, " \t\r\n")
					}
					return true
				}
				currentField = 0
			}
			offset = len([]rune(fields[currentField].value))
		case KeyTab:
			if !field.validateField(field) {
				Beep()
				break
			}
			var dir int
			if ch.Modifier&KeyShift != 0 {
				dir = -1
			} else {
				dir = 1
			}
			currentField += dir
			if currentField < 0 {
				currentField = len(fields) - 1
			} else if currentField >= len(fields) {
				currentField = 0
			}
			offset = len([]rune(fields[currentField].value))
		case KeyLeft:
			if ch.Modifier&KeyControl != 0 {
				for offset > 0 {
					if field.value[offset-1] != ' ' {
						break
					}
					offset--
				}
				for offset > 0 {
					if field.value[offset-1] == ' ' {
						break
					}
					offset--
				}
			} else {
				if offset > 0 {
					offset--
				}
			}
		case KeyRight:
			if ch.Modifier&KeyControl != 0 {
				for offset < len([]rune(field.value)) {
					if field.value[offset] == ' ' {
						for field.value[offset] == ' ' {
							offset++
						}
						break
					}
					offset++
				}
			} else {
				if offset < len([]rune(field.value)) {
					offset++
				}
			}
			offset = min(offset, len([]rune(field.value)))
		case KeyUp:
			if !field.validateField(field) {
				Beep()
				break
			}
			currentField--
			if currentField < 0 {
				currentField = len(fields) - 1
			}
			offset = len([]rune(fields[currentField].value))
		case KeyBackspace:
			if offset > 0 {
				offset--
				field.value = field.value[:len([]rune(field.value))-1]
				fmt.Println("\b \b") // backup
			}
		case KeyIns:
			field.value = fmt.Sprint(field.value[:offset], " ", field.value[offset:])
			Locate(field.row, field.col)
			fmt.Println(field.value)
			break
		case KeyDel:
			v := field.value
			if ch.Modifier&KeyControl != 0 {
				v = field.value[:offset]
			}
			field.value = fmt.Sprint(v[:offset], v[offset+1:])
			Locate(field.row, field.col+offset)
			if offset < len([]rune(v)) {
				fmt.Print(field.value[offset:])
			}
			extra := 0
			if offset > len([]rune(field.value)) {
				extra = 1
			}
			fmt.Print(str.LeftPad("", field.size-len([]rune(field.value))-extra, " "))
		case KeyHome:
			if ch.Modifier&KeyControl != 0 {
				if field.validateField(field) {
					Beep()
					break
				}
				currentField = 0
				offset = len([]rune(fields[currentField].value))
			} else {
				offset = 0
			}
		case KeyEnd:
			if ch.Modifier&KeyControl != 0 {
				if !field.validateField(field) {
					Beep()
					break
				}
				currentField = len(fields) - 1
			}
			fields[currentField].value = strings.TrimRight(fields[currentField].value, " \t\r\n")
			offset = len([]rune(fields[currentField].value))
		case KeyF10, KeyControlEnter:
			for index := range fields {
				fields[index].value = strings.TrimRight(fields[index].value, " \t\r\n")
			}
			return true
		case KeyEscape:
			return false
		default:
			if offset < field.size {
				if field.validateKey(field, ch) {
					fmt.Print(string(ch.Key))
					field.value = fmt.Sprint(field.value[:offset], string(ch.Key), field.value[offset+1:])
					offset++
				} else {
					Beep()
				}
			}
		}
	}
}
