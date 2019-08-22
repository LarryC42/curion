package cons

import "fmt"

// LineInputLen allows entry of up to max characters
func LineInputLen(max int) string {
	var ch KeyEvent
	var entry string

	for {
		ch = GetKey()
		if ch.Key == KeyEnter || ch.Key == KeyEscape {
			break
		}
		switch ch.Key {
		case KeyControlC: // ^c
			SetColor(ColorGray, ColorBlack) // Ensure we're at normal background
			fmt.Print("^c")
			panic("Terminating")

		case KeyBackspace:
			if len([]rune(entry)) > 0 {
				fmt.Print("\b \b")
				n := len(entry) - 1
				entry = (entry)[:n]
			} else {
				Beep()
			}
			break
		default:
			if len([]rune(entry)) >= max {
				Beep()
				break // ignore character
			}
			fmt.Print(string(ch.Key))
			entry += string(ch.Key)
			break
		}
	}
	if ch.Key == 27 {
		return ""
	}
	return entry
}

// LineInput allows entry of a line
func LineInput() string {
	var ch KeyEvent
	var entry string

	for {
		ch = GetKey()
		if ch.Key == KeyEnter || ch.Key == KeyEscape {
			break
		}
		switch ch.Key {
		case KeyControlC: // ^c
			SetColor(ColorGray, ColorBlack) // Ensure we're at normal background
			fmt.Print("^c")
			panic("Terminating")

		case KeyBackspace:
			if len([]rune(entry)) > 0 {
				fmt.Print("\b \b")
				n := len(entry) - 1
				entry = (entry)[:n]
			} else {
				Beep()
			}
			break
		default:
			fmt.Print(string(ch.Key))
			entry += string(ch.Key)
			break
		}
	}
	if ch.Key == 27 {
		return ""
	}
	return entry
}
