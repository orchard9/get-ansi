package ansi

// ClearScreen clears the entire screen and moves the cursor to top-left.
func ClearScreen() string {
	return "\x1b[2J\x1b[H"
}

// ClearLine clears the entire current line.
func ClearLine() string {
	return "\x1b[2K"
}

// ClearLineLeft clears from the cursor to the beginning of the line.
func ClearLineLeft() string {
	return "\x1b[1K"
}

// ClearLineRight clears from the cursor to the end of the line.
func ClearLineRight() string {
	return "\x1b[0K"
}

// AltScreenBuffer switches to the alternate screen buffer.
func AltScreenBuffer() string {
	return "\x1b[?1049h"
}

// NormalScreenBuffer switches back to the main screen buffer.
func NormalScreenBuffer() string {
	return "\x1b[?1049l"
}
