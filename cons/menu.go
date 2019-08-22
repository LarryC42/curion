package cons

import (
	"fmt"
	"lib/dt"
	"lib/str"
	"strconv"
)

// Choose draws the menu and prompts for input.  Title and subTitle are not drawn if empty.
func Choose(title string, subTitle string, items []string, borderStyle int, foreground int8, background int8, inputForeground int8, inputBackground int8, errorForeground int8) int {
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

	maxLength := 0
	for _, v := range items {
		maxLength = max(maxLength, len([]rune(v)))
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
	hl := lineChar[horizontalLine]
	li := lineChar[leftIntersection]
	ri := lineChar[rightIntersection]
	ti := lineChar[topIntersection]
	bi := lineChar[bottomIntersection]
	bl := lineChar[bottomLeft]
	br := lineChar[bottomRight]
	vl := lineChar[verticalLine]

	var line string
	SetWindowSize(24, 80)
	SetColor(foreground, background)
	SetWindowSize(24, 80)
	if len([]rune(title)) > 0 {
		Center(title)
	}
	if len([]rune(subTitle)) > 0 {
		Center(subTitle)
	}
	Center(dt.Dtols(dt.Today()))
	fmt.Println()
	var row, col, count int
	prompt := fmt.Sprint("Choice? (1-", strconv.Itoa(len(items)), ")  ")
	maxLength = max(maxLength, len([]rune(prompt))+1)
	count = len(items)

	if len(items) > 7 {
		half := (count + 1) / 2
		// 2 column menu
		halfLine := str.LeftPad("", maxLength+1, hl)
		line = fmt.Sprint(tl, hl, halfLine, ti, hl, halfLine, tr)
		Center(line)
		fmt1 := fmt.Sprint(vl, " %2d. %-", strconv.Itoa(maxLength), "s ", vl, " %2d. %-", strconv.Itoa(maxLength), "s ", vl)
		fmt2 := fmt.Sprint(vl, " %2d. %-", strconv.Itoa(maxLength), "s ", vl, " %2s  %-", strconv.Itoa(maxLength), "s ", vl)
		for i := 0; i < half; i++ {
			a := items[i]
			b := ""
			if i+half < len(items) {
				b = items[i+half]
			}
			if i+half < len(items) {
				line = fmt.Sprintf(fmt1, i+1, a, i+half+1, b)
			} else {
				line = fmt.Sprintf(fmt2, i+1, a, "", "")
			}
			Center(line)
		}
		line = fmt.Sprint(li, hl, halfLine, bi, hl, halfLine, ri)
		Center(line)
		row = Row()
		col = 40 + (len(prompt)-2)/2
		line = fmt.Sprint(vl, " ", str.LeftPad(prompt, maxLength*2+11, " "), " ", vl)
		Center(line)
		line = fmt.Sprint(bl, str.LeftPad("", maxLength*2+13, hl), br)
		Center(line)
	} else {
		halfLine := str.LeftPad("", maxLength+1, hl)
		line = fmt.Sprint(tl, hl, halfLine, tr)
		Center(line)
		fmt1 := fmt.Sprint(vl, " %2d. %-", strconv.Itoa(maxLength), "s ", vl)
		for i := 0; i < count; i++ {
			a := items[i]
			line = fmt.Sprintf(fmt1, i+1, a)
			Center(line)
		}
		line = fmt.Sprint(li, hl, halfLine, ri)
		Center(line)
		row = Row()
		col = 40 + (len([]rune(prompt))-2)/2
		line = fmt.Sprint(vl, " ", str.LeftPad(prompt, maxLength+4, " "), " ", vl)
		Center(line)
		line = fmt.Sprint(bl, str.LeftPad("", maxLength+6, hl), br)
		Center(line)
	}
	c := 0
	wl := len([]rune(line))
	msgCol := ((80 - wl) / 2)
	SetColor(inputForeground, inputBackground)
	for c < 1 || c > count {
		Locate(row, col)
		fmt.Printf("  \b\b")
		input := LineInputLen(2)
		c, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			c = 0
		}
		if c < 1 || int(c) > count {
			SetColor(errorForeground, background)
			Locate(row+2, msgCol)
			fmt.Printf("'%s' is not valid.  ", input)
			SetColor(inputForeground, inputBackground)
		}
	}

	SetColor(foreground, background)
	Locate(row+2, msgCol)
	fmt.Printf("                  ")
	Locate(row+2, msgCol)
	return c
}
