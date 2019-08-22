package cons

import (
	"fmt"
	"testing"
)

func TestVisualForegroundBackground(t *testing.T) {
	var f, b int8
	for b = 0; b < 8; b++ {
		for f = 0; f < 16; f++ {
			SetColor(f, b)
			fmt.Print(" X ")
		}
		fmt.Println()
	}
	SetColor(ColorWhite, ColorBlack)
}
