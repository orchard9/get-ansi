package ansi

import (
	"fmt"
)

// CursorUp moves the cursor up n lines.
func CursorUp(n int) string {
	return fmt.Sprintf("\x1b[%dA", n)
}

// CursorDown moves the cursor down n lines.
func CursorDown(n int) string {
	return fmt.Sprintf("\x1b[%dB", n)
}

// CursorForward moves the cursor forward n columns.
func CursorForward(n int) string {
	return fmt.Sprintf("\x1b[%dC", n)
}

// CursorBack moves the cursor back n columns.
func CursorBack(n int) string {
	return fmt.Sprintf("\x1b[%dD", n)
}

// CursorPosition moves the cursor to the specified row and column.
// Note: Coordinates are 1-based.
func CursorPosition(row, col int) string {
	return fmt.Sprintf("\x1b[%d;%dH", row, col)
}

// HideCursor hides the cursor.
func HideCursor() string {
	return "\x1b[?25l"
}

// ShowCursor shows the cursor.
func ShowCursor() string {
	return "\x1b[?25h"
}
