package ansi

import (
	"fmt"
	"strconv"
)

// Color represents a color that can be rendered to the terminal.
type Color interface {
	String() string     // Returns the ANSI sequence to set this color as foreground
	Background() string // Returns the ANSI sequence to set this color as background
}

// NoColor represents no color (reset/default).
type NoColor struct{}

func (n NoColor) String() string     { return "" }
func (n NoColor) Background() string { return "" }

// BasicColor represents a standard 4-bit ANSI color.
type BasicColor struct {
	fg string
	bg string
}

func (c BasicColor) String() string     { return c.fg }
func (c BasicColor) Background() string { return c.bg }

// Color256 represents an 8-bit ANSI color.
type Color256Index uint8

func (c Color256Index) String() string {
	return fmt.Sprintf("\x1b[38;5;%dm", c)
}

func (c Color256Index) Background() string {
	return fmt.Sprintf("\x1b[48;5;%dm", c)
}

// RGBColor represents a 24-bit True Color.
type RGBColor struct {
	R, G, B uint8
}

func (c RGBColor) String() string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", c.R, c.G, c.B)
}

func (c RGBColor) Background() string {
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm", c.R, c.G, c.B)
}

// Constructors

func Black() Color         { return BasicColor{fg: FgBlack, bg: BgBlack} }
func Red() Color           { return BasicColor{fg: FgRed, bg: BgRed} }
func Green() Color         { return BasicColor{fg: FgGreen, bg: BgGreen} }
func Yellow() Color        { return BasicColor{fg: FgYellow, bg: BgYellow} }
func Blue() Color          { return BasicColor{fg: FgBlue, bg: BgBlue} }
func Magenta() Color       { return BasicColor{fg: FgMagenta, bg: BgMagenta} }
func Cyan() Color          { return BasicColor{fg: FgCyan, bg: BgCyan} }
func White() Color         { return BasicColor{fg: FgWhite, bg: BgWhite} }
func BrightBlack() Color   { return BasicColor{fg: FgBrightBlack, bg: BgBrightBlack} }
func BrightRed() Color     { return BasicColor{fg: FgBrightRed, bg: BgBrightRed} }
func BrightGreen() Color   { return BasicColor{fg: FgBrightGreen, bg: BgBrightGreen} }
func BrightYellow() Color  { return BasicColor{fg: FgBrightYellow, bg: BgBrightYellow} }
func BrightBlue() Color    { return BasicColor{fg: FgBrightBlue, bg: BgBrightBlue} }
func BrightMagenta() Color { return BasicColor{fg: FgBrightMagenta, bg: BgBrightMagenta} }
func BrightCyan() Color    { return BasicColor{fg: FgBrightCyan, bg: BgBrightCyan} }
func BrightWhite() Color   { return BasicColor{fg: FgBrightWhite, bg: BgBrightWhite} }

func Color256(index uint8) Color {
	return Color256Index(index)
}

func RGB(r, g, b uint8) Color {
	return RGBColor{R: r, G: g, B: b}
}

func Hex(hex string) Color {
	if len(hex) > 0 && hex[0] == '#' {
		hex = hex[1:]
	}

	if len(hex) == 3 {
		// Expand short hex: "F0A" -> "FF00AA"
		r, err1 := strconv.ParseUint(string(hex[0])+string(hex[0]), 16, 8)
		g, err2 := strconv.ParseUint(string(hex[1])+string(hex[1]), 16, 8)
		b, err3 := strconv.ParseUint(string(hex[2])+string(hex[2]), 16, 8)
		if err1 == nil && err2 == nil && err3 == nil {
			return RGB(uint8(r), uint8(g), uint8(b))
		}
	}

	if len(hex) == 6 {
		r, err1 := strconv.ParseUint(hex[0:2], 16, 8)
		g, err2 := strconv.ParseUint(hex[2:4], 16, 8)
		b, err3 := strconv.ParseUint(hex[4:6], 16, 8)
		if err1 == nil && err2 == nil && err3 == nil {
			return RGB(uint8(r), uint8(g), uint8(b))
		}
	}

	// Fallback to no color if invalid
	return NoColor{}
}
