package ansi

import (
	"testing"
)

func TestColors(t *testing.T) {
	tests := []struct {
		name   string
		color  Color
		wantFg string
		wantBg string
	}{
		{"Red", Red(), "\x1b[31m", "\x1b[41m"},
		{"BrightBlue", BrightBlue(), "\x1b[94m", "\x1b[104m"},
		{"Color256", Color256(208), "\x1b[38;5;208m", "\x1b[48;5;208m"},
		{"RGB", RGB(255, 0, 0), "\x1b[38;2;255;0;0m", "\x1b[48;2;255;0;0m"},
		{"Hex Short", Hex("#F00"), "\x1b[38;2;255;0;0m", "\x1b[48;2;255;0;0m"},
		{"Hex Long", Hex("00FF00"), "\x1b[38;2;0;255;0m", "\x1b[48;2;0;255;0m"},
		{"NoColor", NoColor{}, "", ""},
		{"Hex Invalid", Hex("ZZZ"), "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.color.String(); got != tt.wantFg {
				t.Errorf("String() = %q, want %q", got, tt.wantFg)
			}
			if got := tt.color.Background(); got != tt.wantBg {
				t.Errorf("Background() = %q, want %q", got, tt.wantBg)
			}
		})
	}
}

func TestStyle_Render(t *testing.T) {
	red := Red()
	blue := Blue()

	tests := []struct {
		name  string
		style Style
		text  string
		want  string
	}{
		{
			name:  "Plain",
			style: NewStyle(),
			text:  "Hello",
			want:  "Hello",
		},
		{
			name:  "Red Foreground",
			style: NewStyle().Foreground(red),
			text:  "Hello",
			want:  "\x1b[31mHello\x1b[0m",
		},
		{
			name:  "Bold Italic",
			style: NewStyle().Bold(true).Italic(true),
			text:  "World",
			want:  "\x1b[1m\x1b[3mWorld\x1b[0m",
		},
		{
			name:  "Complex",
			style: NewStyle().Foreground(red).Background(blue).Underline(true),
			text:  "Test",
			want:  "\x1b[31m\x1b[44m\x1b[4mTest\x1b[0m",
		},
		{
			name:  "Turn Off",
			style: NewStyle().Bold(true).Bold(false),
			text:  "Plain",
			want:  "Plain",
		},
		{
			name:  "Hidden",
			style: NewStyle().Hidden(true),
			text:  "Secret",
			want:  "\x1b[8mSecret\x1b[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.style.Render(tt.text)
			if got != tt.want {
				t.Errorf("Render() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestStringUtilities(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantLen int
	}{
		{
			name:    "Plain",
			input:   "Hello",
			want:    "Hello",
			wantLen: 5,
		},
		{
			name:    "Simple Color",
			input:   "\x1b[31mHello\x1b[0m",
			want:    "Hello",
			wantLen: 5,
		},
		{
			name:    "Complex",
			input:   "\x1b[31m\x1b[1mBold Red\x1b[0m",
			want:    "Bold Red",
			wantLen: 8,
		},
		{
			name:    "Cursor Move",
			input:   "Move\x1b[1AUp",
			want:    "MoveUp",
			wantLen: 6,
		},
		{
			name:    "Unicode",
			input:   "\x1b[32m世界\x1b[0m",
			want:    "世界",
			wantLen: 2, // 2 runes
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Strip(tt.input); got != tt.want {
				t.Errorf("Strip() = %q, want %q", got, tt.want)
			}
			if got := VisibleLen(tt.input); got != tt.wantLen {
				t.Errorf("VisibleLen() = %d, want %d", got, tt.wantLen)
			}
		})
	}
}

func TestCursorAndScreen(t *testing.T) {
	// Basic smoke tests to ensure constants/formatting are correct
	if got := CursorUp(5); got != "\x1b[5A" {
		t.Errorf("CursorUp(5) = %q", got)
	}
	if got := CursorPosition(10, 20); got != "\x1b[10;20H" {
		t.Errorf("CursorPosition(10, 20) = %q", got)
	}
	if got := ClearScreen(); got != "\x1b[2J\x1b[H" {
		t.Errorf("ClearScreen() = %q", got)
	}
}
