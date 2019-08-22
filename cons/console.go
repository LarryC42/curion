package cons

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

var (
	kernel32DLL                 = syscall.NewLazyDLL("kernel32.dll")
	wGetConsoleTitle            = kernel32DLL.NewProc("GetConsoleTitleW")
	wSetConsoleTitle            = kernel32DLL.NewProc("SetConsoleTitleW")
	wGetConsoleDisplayMode      = kernel32DLL.NewProc("GetConsoleDisplayMode")
	wSetConsoleDisplayMode      = kernel32DLL.NewProc("SetConsoleDisplayMode")
	wSetConsoleTextAttribute    = kernel32DLL.NewProc("SetConsoleTextAttribute")
	wGetConsoleScreenBufferInfo = kernel32DLL.NewProc("GetConsoleScreenBufferInfo")
	wSetConsoleCursorPosition   = kernel32DLL.NewProc("SetConsoleCursorPosition")
	wSetConsoleWindowInfo       = kernel32DLL.NewProc("SetConsoleWindowInfo")
	wSetConsoleScreenBufferSize = kernel32DLL.NewProc("SetConsoleScreenBufferSize")
	wFillConsoleOutputCharacter = kernel32DLL.NewProc("FillConsoleOutputCharacterA")
	wFillConsoleOutputAttribute = kernel32DLL.NewProc("FillConsoleOutputAttribute")
	wSetConsoleCP               = kernel32DLL.NewProc("SetConsoleCP")
	wSetConsoleOutputCP         = kernel32DLL.NewProc("SetConsoleOutputCP")
	wSetConsoleMode             = kernel32DLL.NewProc("SetConsoleMode")
	wReadConsoleInput           = kernel32DLL.NewProc("ReadConsoleInputA")
	wPeekConsoleInput           = kernel32DLL.NewProc("PeekConsoleInput")
	wGetStdHandle               = kernel32DLL.NewProc("GetStdHandle")
)

const (
	wVkShift   = 0x10
	wVkControl = 0x11
	wVkMenu    = 0x12
	wVkCapital = 0x14
	wVkLWin    = 0x5b
	wVkRWin    = 0x5c

	wCapsLockOn       = 0x80
	wEnhancedKey      = 0x100
	wLeftAltPressed   = 0x2
	wLeftCtrlPressed  = 0x8
	wNumlockOn        = 0x20
	wRightAltPressed  = 0x1
	wRightCtrlPressed = 0x4
	wScrollLockOn     = 0x40
	wShiftPressed     = 0x10
)
const (
	// KeyControlC is Control+C
	KeyControlC = 3
	// KeyEnter is the Enter key
	KeyEnter = 13
	// KeyControlEnter is the Control+Enter keys
	KeyControlEnter = 10
	// KeyTab is the Tab key
	KeyTab = 9
	// KeyBackspace is the backspace key
	KeyBackspace = 8
	// KeyEscape is the escape key
	KeyEscape = 27
	// KeyPageUp is the page up key
	KeyPageUp = 129
	// KeyPageDown is the page down key
	KeyPageDown = 130
	// KeyEnd is the end key
	KeyEnd = 131
	// KeyHome is the home key
	KeyHome = 132
	// KeyLeft is the left key
	KeyLeft = 133
	// KeyUp is the up key
	KeyUp = 134
	// KeyRight is the right key
	KeyRight = 135
	// KeyDown is the down key
	KeyDown = 136
	// KeyIns is the insert key
	KeyIns = 141
	// KeyDel is the delete key
	KeyDel = 142
	// KeyF1 is the F1 key
	KeyF1 = 143
	// KeyF2 is the F2 key
	KeyF2 = 144
	// KeyF3 is the F3 key
	KeyF3 = 145
	// KeyF4 is the F4 key
	KeyF4 = 146
	// KeyF5 is the F5 key
	KeyF5 = 147
	// KeyF6 is the F6 key
	KeyF6 = 148
	// KeyF7 is the F7 key
	KeyF7 = 149
	// KeyF8 is the F8 key
	KeyF8 = 150
	// KeyF9 is the F9 key
	KeyF9 = 151
	// KeyF10 is the F10 key
	KeyF10 = 152
	// KeyF11 is the F11 key
	KeyF11 = 153
	// KeyF12 is the F12 key
	KeyF12 = 154
	// KeyF13 is the F13 key
	KeyF13 = 155
	// KeyF14 is the F14 key
	KeyF14 = 156
	// KeyF15 is the F15 key
	KeyF15 = 157
	// KeyF16 is the F16 key
	KeyF16 = 158
	// KeyF17 is the F17 key
	KeyF17 = 159
	// KeyF18 is the F18 key
	KeyF18 = 160
	// KeyF19 is the F19 key
	KeyF19 = 161
	// KeyF20 is the F20 key
	KeyF20 = 162
	// KeyF21 is the F21 key
	KeyF21 = 163
	// KeyF22 is the F22 key
	KeyF22 = 164
	// KeyF23 is the F23 key
	KeyF23 = 165
	// KeyF24 is the F24 key
	KeyF24 = 166
)

const (
	// LineStyleNone draws no line
	LineStyleNone = 0
	// LineStyleSingle draws a single line
	LineStyleSingle = 1
	// LineStyleDouble draws a double line3
	LineStyleDouble = 2
)
const (
	// KeyCapsLock is the caps lock modifier
	KeyCapsLock = 1
	// KeyAlt is the alt modifier
	KeyAlt = 2
	// KeyControl is the control modifier
	KeyControl = 4
	// KeyShift is the shift modifier
	KeyShift = 8
)

// KeyEvent represents a key press event
type KeyEvent struct {
	// Key is the key that was pressed
	Key uint8
	// Modifier is the bit mask of modifiers that were down when the key was pressed
	Modifier int8
}

const (
	// ColorBlack is the black color
	ColorBlack = 0
	// ColorBlue is the blue color
	ColorBlue = 1
	// ColorGreen is the green color
	ColorGreen = 2
	// ColorCyan is the cyan color
	ColorCyan = 3
	// ColorRed is the red color
	ColorRed = 4
	// ColorMagenta is the magenta color
	ColorMagenta = 5
	// ColorYellow is the yellow color
	ColorYellow = 6
	// ColorWhite is the white color
	ColorWhite = 7
	// ColorBright is a modifier for bright colors
	ColorBright = 8
	// ColorGray is the gray color
	ColorGray = ColorBlack | ColorBright
	// ColorBrightBlue is the bright blue color
	ColorBrightBlue = ColorBlue | ColorBright
	// ColorBrightGreen is the bright green color
	ColorBrightGreen = ColorGreen | ColorBright
	// ColorBrightCyan is the bright cyan color
	ColorBrightCyan = ColorCyan | ColorBright
	// ColorBrightRed is the bright red color
	ColorBrightRed = ColorRed | ColorBright
	// ColorBrightMagenta is the bright magenta color
	ColorBrightMagenta = ColorMagenta | ColorBright
	// ColorBrightYellow is the bright yellow color
	ColorBrightYellow = ColorYellow | ColorBright
	// ColorBrightWhite is the bright white color
	ColorBrightWhite = ColorWhite | ColorBright
)

type (
	// wCharInfo is the windows CHAR_INFO structure
	wCharInfo struct {
		UnicodeChar uint16
		Attributes  uint16
	}

	// wConsoleCursorInfo is the windows CONSOLE_CURSOR_INFO
	wConsoleCursorInfo struct {
		Size    uint32
		Visible int32
	}

	// wConsoleScreenBufferInfo is the windows CONSOLE_SCREEN_BUFFER_INFO
	wConsoleScreenBufferInfo struct {
		Size              wCoord
		CursorPosition    wCoord
		Attributes        uint16
		Window            wSmallRect
		MaximumWindowSize wCoord
	}

	// wCoord is the windows COORD structure
	wCoord struct {
		X int16
		Y int16
	}

	// wSmallRect is the windows SMALL_RECT structure
	wSmallRect struct {
		Left   int16
		Top    int16
		Right  int16
		Bottom int16
	}

	// wInputRecord is the windows INPUT_RECORD structure
	wInputRecord struct {
		EventType uint16
		KeyEvent  wKeyEventRecord
	}

	// wKeyEventRecord is the windows KEY_EVENT_RECORD structure
	wKeyEventRecord struct {
		KeyDown         int32
		RepeatCount     uint16
		VirtualKeyCode  uint16
		VirtualScanCode uint16
		ASCIIChar       [2]uint8
		ControlKeyState uint32
	}

	// wWindowBufferSize is the windows WINDOW_BUFFER_SIZEW structure
	wWindowBufferSize struct {
		Size wCoord
	}
)

var doBeep = false

// SetTitle sets the console window title
func SetTitle(title string) {
	if len(title) > 240 {
		title = title[0:240]
	}
	wSetConsoleTitle.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))))
}

// GetTitle gets the console window title
func GetTitle() string {
	arr := make([]uint16, 257)
	var arrSize = uint32(len(arr) - 1)

	wGetConsoleTitle.Call(uintptr(unsafe.Pointer(&arr[0])), uintptr(arrSize))
	return strings.TrimRight(string(utf16.Decode(arr)), " \000")
}

// GetStdOut gets the stdout windows handle
func GetStdOut() uintptr {
	return uintptr(os.Stdout.Fd())
}

// GetStdIn gets the stdin windows handle
func GetStdIn() uintptr {
	return uintptr(os.Stdin.Fd())
}

// IsFullScreen returns true if full screen or false if windowed
func IsFullScreen() bool {
	var mode uint32
	wGetConsoleDisplayMode.Call(uintptr(unsafe.Pointer(&mode)))
	return (mode == 1) // 1=full screen, 2=windowed
}

// SetFullScreen changes to full screen for 32-bit applications.  Does not work for 64bit applications
func SetFullScreen() {
	var c wCoord
	if runtime.GOARCH == "amd64" {
		return
	}
	stdout := GetStdOut()
	fmt.Println(stdout)
	a, b, d := wSetConsoleDisplayMode.Call(stdout, 1, uintptr(unsafe.Pointer(&c)))
	fmt.Println("Set FullScreen", a, b, d, c.X, c.Y)
}

// SetWindowed changes to windowed mode for 32-bit applications.  Does not work for 64bit applications
func SetWindowed() {
	var c wCoord
	if runtime.GOARCH == "amd64" {
		return
	}

	stdout := GetStdOut()
	fmt.Println(stdout)
	a, b, d := wSetConsoleDisplayMode.Call(stdout, 2, uintptr(unsafe.Pointer(&c)))
	fmt.Println("Set Windowed", a, b, d, c.X, c.Y)
}

// SetColor sets the foreground and background colors for subsequent output
func SetColor(foreground int8, background int8) {
	stdout := GetStdOut()
	wSetConsoleTextAttribute.Call(stdout, uintptr(foreground|(background<<4)))
}

// Rows returns the number of rows on the console window
func Rows() int {
	var info wConsoleScreenBufferInfo
	stdout := GetStdOut()
	wGetConsoleScreenBufferInfo.Call(stdout, uintptr(unsafe.Pointer(&info)))
	return int(info.Window.Bottom - info.Window.Top + 1)
}

// BufferRows returns the number of buffer rows on the console window
func BufferRows() int {
	var info wConsoleScreenBufferInfo
	stdout := GetStdOut()
	wGetConsoleScreenBufferInfo.Call(stdout, uintptr(unsafe.Pointer(&info)))
	return int(info.Size.Y)
}

// Cols returns the number of columns on the console screen
func Cols() int {
	var info wConsoleScreenBufferInfo
	stdout := GetStdOut()
	wGetConsoleScreenBufferInfo.Call(stdout, uintptr(unsafe.Pointer(&info)))
	return int(info.Window.Right - info.Window.Left + 1)
}

// BufferCols returns the number of buffer columns on the console screen
func BufferCols() int {
	var info wConsoleScreenBufferInfo
	stdout := GetStdOut()
	wGetConsoleScreenBufferInfo.Call(stdout, uintptr(unsafe.Pointer(&info)))
	return int(info.Size.X)
}

// Row returns the current screen row the cursor is on from 0 to Rows-1
func Row() int {
	var info wConsoleScreenBufferInfo
	stdout := GetStdOut()
	wGetConsoleScreenBufferInfo.Call(stdout, uintptr(unsafe.Pointer(&info)))
	return int(info.CursorPosition.Y)
}

// Col returns the current screen column the cursor is on from 0 to Cols-1
func Col() int {
	var info wConsoleScreenBufferInfo
	stdout := GetStdOut()
	wGetConsoleScreenBufferInfo.Call(stdout, uintptr(unsafe.Pointer(&info)))
	return int(info.CursorPosition.X)
}

// Locate positions the cursor to a row and column.  valid values are 0 to Rows-1 and 0 to Cols-1
func Locate(row int, col int) {
	var coord wCoord
	coord.X = int16(col)
	coord.Y = int16(row)

	stdout := GetStdOut()
	wSetConsoleCursorPosition.Call(stdout, coordToUintptr(coord))
}

// SetWindowSize sets the console window rows and columns and matches the buffer to it.
func SetWindowSize(rows int, cols int) {
	var rect wSmallRect
	rect.Top = 0
	rect.Left = 0
	rect.Bottom = int16(rows - 1)
	rect.Right = int16(cols - 1)
	stdout := GetStdOut()
	wSetConsoleWindowInfo.Call(stdout, 1, uintptr(unsafe.Pointer(&rect)))

	var coord wCoord
	coord.X = int16(cols)
	coord.Y = int16(rows)
	wSetConsoleScreenBufferSize.Call(stdout, coordToUintptr(coord))
}

func coordToUintptr(coord wCoord) uintptr {
	return uintptr(*((*uint32)(unsafe.Pointer(&coord))))
}

// SetWindowAndBufferSize sets the window rows/cols and buffer rows/cols.  buffer must be >= window size
func SetWindowAndBufferSize(rows int, cols int, bufRows int, bufCols int) {
	var rect wSmallRect
	rect.Top = 0
	rect.Left = 0
	rect.Bottom = int16(rows - 1)
	rect.Right = int16(cols - 1)
	stdout := GetStdOut()
	wSetConsoleWindowInfo.Call(stdout, 1, uintptr(unsafe.Pointer(&rect)))

	var coord wCoord
	coord.X = int16(bufCols)
	coord.Y = int16(bufRows)
	wSetConsoleScreenBufferSize.Call(stdout, coordToUintptr(coord))
}

// Cls clears the screen using the current foreground/background
func Cls() {
	var coordScreen wCoord
	var cCharsWritten uint32
	var csbi wConsoleScreenBufferInfo
	var dwConSize uint32

	// Get the number of character cells in the current buffer.
	stdout := GetStdOut()
	if err, _, _ := wGetConsoleScreenBufferInfo.Call(stdout, uintptr(unsafe.Pointer(&csbi))); err == 0 {
		return
	}

	dwConSize = uint32(csbi.Size.X * csbi.Size.Y)

	var c = uintptr(' ')
	// Fill the entire screen with blanks.

	if err, _, _ := wFillConsoleOutputCharacter.Call(stdout, c, uintptr(dwConSize), coordToUintptr(coordScreen), uintptr(unsafe.Pointer(&cCharsWritten))); err == 0 {
		return
	}

	// Get the current text attribute.
	if err, _, _ := wGetConsoleScreenBufferInfo.Call(stdout, uintptr(unsafe.Pointer(&csbi))); err == 0 {
		return
	}

	// Set the buffer's attributes accordingly.
	if err, _, _ := wFillConsoleOutputAttribute.Call(stdout,
		uintptr(csbi.Attributes),
		uintptr(dwConSize),
		coordToUintptr(coordScreen),
		uintptr(unsafe.Pointer(&cCharsWritten))); err == 0 {
		return
	}

	// Put the cursor at its home coordinates.
	wSetConsoleCursorPosition.Call(stdout, coordToUintptr(coordScreen))
	wSetConsoleCP.Call(uintptr(65001))
	wSetConsoleOutputCP.Call(uintptr(65001))
}

// Center writes a string centered in the console and advances to the next line.
func Center(value string) {
	length := len([]rune(value))
	col := (Cols() - length) / 2
	Locate(Row(), col)
	fmt.Println(value)
}

// GetKey returns a keypress event
func GetKey() KeyEvent {
	var c uint8
	var read uint32
	var kc uint16
	var rec wInputRecord
	stdin := GetStdIn()
	wSetConsoleMode.Call(stdin, uintptr(0))

	for rec.EventType != 1 || (rec.EventType == 1 && rec.KeyEvent.KeyDown == 0) || c == 0 {
		wReadConsoleInput.Call(stdin, uintptr(unsafe.Pointer(&rec)), uintptr(1), uintptr(unsafe.Pointer(&read)))
		c = rec.KeyEvent.ASCIIChar[0]
		switch rec.KeyEvent.VirtualKeyCode {
		case wVkShift, wVkControl, wVkMenu, wVkCapital, wVkLWin, wVkRWin:
			c = 0
		}
		kc = rec.KeyEvent.VirtualKeyCode
		if kc >= 0x21 && kc <= 0x2f {
			c = 1
		}
		if kc >= 0x70 && kc <= 0x87 {
			c = 1
		}
	}
	// Special Arrow handling
	if kc >= 0x21 && kc <= 0x2f {
		c = uint8(129 - 0x21 + kc)
	}
	if kc >= 0x70 && kc <= 0x87 {
		c = uint8(143 - 0x70 + kc)
	}

	mods := 0
	if (rec.KeyEvent.ControlKeyState & wCapsLockOn) != 0 {
		mods |= KeyCapsLock
	}
	if (rec.KeyEvent.ControlKeyState & (wLeftAltPressed | wRightAltPressed)) != 0 {
		mods |= KeyAlt
	}
	if (rec.KeyEvent.ControlKeyState & (wLeftCtrlPressed | wRightCtrlPressed)) != 0 {
		mods |= KeyControl
	}
	if (rec.KeyEvent.ControlKeyState & wShiftPressed) != 0 {
		mods |= KeyShift
	}
	return KeyEvent{Key: uint8(c), Modifier: int8(mods)}
}

// Inkey returns 0 key if no input available or a keypress if available.  Does not block or wait.
func Inkey() KeyEvent {
	var rec wInputRecord
	var e = KeyEvent{0, 0}
	var reads uint
	stdin := GetStdIn()
	for {
		wPeekConsoleInput.Call(stdin, uintptr(unsafe.Pointer(&rec)), 1, uintptr(unsafe.Pointer(&reads)))
		if reads > 0 && rec.EventType == 1 {
			break
		}
		wReadConsoleInput.Call(stdin, uintptr(unsafe.Pointer(&rec)), 1, uintptr(unsafe.Pointer(&reads))) // Discard non Keyboard events
	}
	if reads == 0 {
		return e
	}
	return GetKey()
}

// GetBeep returns true if a beep will make a sound or false if it will be silent.
func GetBeep() bool {
	return doBeep
}

// SetBeep will allow beep to sound if value is true or be silent if value is false.
func SetBeep(value bool) {
	doBeep = value
}

// Beep plays a sound if SetBeep was passed true
func Beep() {
	if doBeep {
		fmt.Print('\a')
	}
}

// IsNumeric tests if string is an integer
func IsNumeric(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}
