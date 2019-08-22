package cons

import (
	"fmt"
	"runtime"
	"testing"
)

func TestUnitClear(t *testing.T) {
	Cls()
	if Row() != 0 || Col() != 0 {
		t.Error("Row or Col wasn't 0.  Got (", Row(), ",", Col(), ")")
	}

}
func TestUnitPos(t *testing.T) {
	Cls()
	Locate(10, 11)
	if Row() != 10 || Col() != 11 {
		t.Error("Expected (10,11) got (", Row(), ",", Col(), ")")
	}
	Locate(10, 0)
}

func TestUnitTitle(t *testing.T) {
	SetTitle("My Title")
	if GetTitle() != "My Title" {
		t.Error(fmt.Sprintf("%s%s%s", "Title wasn't 'My Title', got '", GetTitle(), "'"))
	}
}

func TestUnitFullScreen(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		SetFullScreen()
		if IsFullScreen() {
			t.Error("64bit doesn't support full screen")
		}
		SetWindowed()
	} else {
		SetFullScreen()
		if !IsFullScreen() {
			t.Error("32bit not full screen")
		}
		SetWindowed()
		if IsFullScreen() {
			t.Error("32bit didn't go windowed")
		}
	}
}

func TestUnitWindowSize(t *testing.T) {
	SetWindowSize(10, 40)
	if Rows() != 10 || Cols() != 40 {
		t.Error("Expected Window size (40,10), got (", Rows(), ",", Cols(), ")")
	}
	SetWindowSize(25, 80)
}
func TestUnitWindowSizeAndBuffer(t *testing.T) {
	SetWindowAndBufferSize(10, 40, 25, 80)
	if Rows() != 10 || Cols() != 40 {
		t.Error("Expected Window size (10,40), got (", Rows(), ",", Cols(), ")")
	}
	if BufferRows() != 25 || BufferCols() != 80 {
		t.Error("Expected Window Buffer size (25,80), got (", BufferRows(), ",", BufferCols(), ")")
	}
	SetWindowSize(25, 80)
}
