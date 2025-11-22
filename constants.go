package ansi

// ANSI Escape Sequences
const (
	Reset = "\x1b[0m"

	// Modifiers
	Bold          = "\x1b[1m"
	Faint         = "\x1b[2m"
	Italic        = "\x1b[3m"
	Underline     = "\x1b[4m"
	Blink         = "\x1b[5m"
	Reverse       = "\x1b[7m"
	Hidden        = "\x1b[8m"
	Strikethrough = "\x1b[9m"

	// Modifiers Reset
	NoBold          = "\x1b[22m"
	NoFaint         = "\x1b[22m"
	NoItalic        = "\x1b[23m"
	NoUnderline     = "\x1b[24m"
	NoBlink         = "\x1b[25m"
	NoReverse       = "\x1b[27m"
	NoHidden        = "\x1b[28m"
	NoStrikethrough = "\x1b[29m"
)

// Foreground Colors
const (
	FgBlack   = "\x1b[30m"
	FgRed     = "\x1b[31m"
	FgGreen   = "\x1b[32m"
	FgYellow  = "\x1b[33m"
	FgBlue    = "\x1b[34m"
	FgMagenta = "\x1b[35m"
	FgCyan    = "\x1b[36m"
	FgWhite   = "\x1b[37m"

	FgBrightBlack   = "\x1b[90m"
	FgBrightRed     = "\x1b[91m"
	FgBrightGreen   = "\x1b[92m"
	FgBrightYellow  = "\x1b[93m"
	FgBrightBlue    = "\x1b[94m"
	FgBrightMagenta = "\x1b[95m"
	FgBrightCyan    = "\x1b[96m"
	FgBrightWhite   = "\x1b[97m"
)

// Background Colors
const (
	BgBlack   = "\x1b[40m"
	BgRed     = "\x1b[41m"
	BgGreen   = "\x1b[42m"
	BgYellow  = "\x1b[43m"
	BgBlue    = "\x1b[44m"
	BgMagenta = "\x1b[45m"
	BgCyan    = "\x1b[46m"
	BgWhite   = "\x1b[47m"

	BgBrightBlack   = "\x1b[100m"
	BgBrightRed     = "\x1b[101m"
	BgBrightGreen   = "\x1b[102m"
	BgBrightYellow  = "\x1b[103m"
	BgBrightBlue    = "\x1b[104m"
	BgBrightMagenta = "\x1b[105m"
	BgBrightCyan    = "\x1b[106m"
	BgBrightWhite   = "\x1b[107m"
)
