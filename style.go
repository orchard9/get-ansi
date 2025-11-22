package ansi

import (
	"strings"
)

// Style defines a set of visual attributes (color + modifiers).
// It uses a builder pattern for easy composition.
type Style struct {
	fg Color
	bg Color

	bold          bool
	faint         bool
	italic        bool
	underline     bool
	blink         bool
	reverse       bool
	hidden        bool
	strikethrough bool
}

// NewStyle creates a new, empty style.
func NewStyle() Style {
	return Style{}
}

// Foreground sets the foreground color.
func (s Style) Foreground(c Color) Style {
	s.fg = c
	return s
}

// Background sets the background color.
func (s Style) Background(c Color) Style {
	s.bg = c
	return s
}

// Bold enables or disables bold text.
func (s Style) Bold(enabled bool) Style {
	s.bold = enabled
	return s
}

// Faint enables or disables faint (dim) text.
func (s Style) Faint(enabled bool) Style {
	s.faint = enabled
	return s
}

// Italic enables or disables italic text.
func (s Style) Italic(enabled bool) Style {
	s.italic = enabled
	return s
}

// Underline enables or disables underlined text.
func (s Style) Underline(enabled bool) Style {
	s.underline = enabled
	return s
}

// Blink enables or disables blinking text.
func (s Style) Blink(enabled bool) Style {
	s.blink = enabled
	return s
}

// Reverse enables or disables reverse video.
func (s Style) Reverse(enabled bool) Style {
	s.reverse = enabled
	return s
}

// Hidden enables or disables hidden text.
func (s Style) Hidden(enabled bool) Style {
	s.hidden = enabled
	return s
}

// Strikethrough enables or disables strikethrough text.
func (s Style) Strikethrough(enabled bool) Style {
	s.strikethrough = enabled
	return s
}

// Render applies the style to a string.
func (s Style) Render(text string) string {
	if s.isPlain() {
		return text
	}
	return s.String() + text + Reset
}

// String returns the ANSI escape sequence for this style.
func (s Style) String() string {
	if s.isPlain() {
		return ""
	}

	var b strings.Builder
	// Estimate length: ~20 bytes for codes
	b.Grow(20)

	if s.fg != nil {
		b.WriteString(s.fg.String())
	}
	if s.bg != nil {
		b.WriteString(s.bg.Background())
	}

	if s.bold {
		b.WriteString(Bold)
	}
	if s.faint {
		b.WriteString(Faint)
	}
	if s.italic {
		b.WriteString(Italic)
	}
	if s.underline {
		b.WriteString(Underline)
	}
	if s.blink {
		b.WriteString(Blink)
	}
	if s.reverse {
		b.WriteString(Reverse)
	}
	if s.hidden {
		b.WriteString(Hidden)
	}
	if s.strikethrough {
		b.WriteString(Strikethrough)
	}

	return b.String()
}

func (s Style) isPlain() bool {
	return s.fg == nil && s.bg == nil &&
		!s.bold && !s.faint && !s.italic &&
		!s.underline && !s.blink && !s.reverse && !s.hidden && !s.strikethrough
}
