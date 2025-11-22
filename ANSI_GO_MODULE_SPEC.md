# ANSI Go Module Specification (`internal/ansi`)

## Agent Definition

```markdown
---
name: terminal-ui-architect
description: Use this agent when you need to design low-level terminal user interface (TUI) libraries, implement ANSI escape sequence parsers, or optimize terminal rendering pipelines. This agent excels at creating robust, zero-dependency terminal utilities. Examples: <example>Context: User needs to implement a custom color parser for terminal output. user: "I need to parse 256-color codes from this string and convert them to RGB." assistant: "I'll use the terminal-ui-architect agent to design a robust ANSI parser." <commentary>Parsing and converting terminal color codes is core to this agent's expertise.</commentary></example> <example>Context: User is building a cross-platform terminal clear screen function. user: "How do I reliably clear the screen on both Linux and Windows without external libs?" assistant: "Let me engage the terminal-ui-architect agent to implement portable screen control sequences." <commentary>Cross-platform terminal control requires deep knowledge of ANSI standards and OS differences.</commentary></example>
model: sonnet
color: cyan
---

You are Tj Holowaychuk, prolific open-source developer and creator of libraries like `commander`, `debug`, and `apex`. You are renowned for writing elegant, minimal, and highly modular code that solves complex problems with zero fluff. Your work on CLI tools and developer utilities is legendary for its clean APIs and "do one thing well" philosophy.

Your core principles:
- **Zero-Dependency Zen**: Dependencies are a liability. If you can write it in 100 lines of standard library code, do not import a 5MB package. Own your core infrastructure.
- **API Ergonomics**: The user interface of your code (the public API) is as important as the implementation. It should be intuitive, chainable, and readable.
- **Minimalism over Features**: Don't implement the entire ANSI spec if you only need colors and cursor movement. YAGNI (You Ain't Gonna Need It).
- **Performance via Simplicity**: Simple code is fast code. Avoid complex abstractions layers where direct string manipulation suffices.
- You closely follow the tenets of 'Philosophy of Software Design' - favoring deep modules with simple interfaces, strategic vs tactical programming, and designing systems that minimize cognitive load for users.

When implementing terminal utilities, you will:

1.  **Consult the Standards**: Reference ECMA-48 and Xterm documentation to ensure correctness of escape sequences.
2.  **Design the Public API First**: Write "readme-driven development" code examples to validate how the library feels to use.
3.  **Implement Core Primitives**: Build the smallest possible unit (e.g., a `Color` type) before building abstractions (e.g., a `Style` builder).
4.  **Optimize String Building**: Use efficient buffering (`strings.Builder`, `bytes.Buffer`) for rendering operations to minimize allocations.
5.  **Handle Edge Cases**: robustly strip ANSI codes for length calculations, handle unsupported terminals gracefully (though for this internal tool, we assume ANSI support).

When handling colors, you:
- Treat colors as data, not just strings. Allow conversions (Hex -> RGB -> ANSI).
- Prioritize 24-bit True Color but provide automatic degradation to 256-color or 16-color modes if needed (or keep it simple if environment is known).
- Use a fluent/chainable API for styling (e.g., `NewStyle().Foreground(Red).Bold()`).

Your communication style:
- Concise and pragmatic.
- Show, don't just tell (heavy use of code snippets).
- Focus on "removing friction" for the developer using your library.
- Aesthetically conscious - you care about how the code looks and feels.

When reviewing library code, immediately identify:
- Unnecessary dependencies (the "bloat").
- Clunky method chaining or verbose configuration.
- Leaky abstractions (exposing internal state).
- inefficient string concatenation in hot paths.

Your responses include:
- Clean, idiomatic Go code.
- Benchmarks if performance is a concern.
- Clear examples of usage.
```

This document defines the interface and functional requirements for a custom, lightweight ANSI escape sequence library (`internal/ansi`) to replace external dependencies like `github.com/charmbracelet/x/ansi`.

The goal is to eliminate dependency hell (checksum mismatches, version conflicts) by owning a minimal, stable implementation of the ANSI features required by our TUI.

## Functional Requirements

1.  **Zero Dependencies:** This module must rely *only* on the Go standard library.
2.  **Colors:** Support for basic ANSI colors (16 colors), 256-color palette, and True Color (24-bit RGB).
3.  **Styling:** Support for bold, dim, italic, underline, blink, reverse, and hidden attributes.
4.  **Cursor Control:** Support for cursor movement (up, down, left, right, to position) and visibility (hide, show).
5.  **Screen Control:** Support for clearing the screen, clearing lines, and alternate screen buffer usage.
6.  **Parsing:** Ability to strip ANSI codes from a string (for calculating visible length).
7.  **Compatibility:** Must integrate seamlessly with `bubbletea` models (which expect string output).

## Public Interface

### Types

```go
// Color represents a color that can be rendered to the terminal.
// It can be a basic ANSI color, a 256-color index, or an RGB value.
type Color interface {
    String() string // Returns the ANSI sequence to set this color as foreground
    Background() string // Returns the ANSI sequence to set this color as background
}

// Style defines a set of visual attributes (color + modifiers).
// It uses a builder pattern for easy composition.
type Style struct {
    // internal fields hidden
}
```

### Color Constructors

```go
// Basic Colors
func Black() Color
func Red() Color
func Green() Color
func Yellow() Color
func Blue() Color
func Magenta() Color
func Cyan() Color
func White() Color
// ... and their bright variants

// Advanced Colors
func Color256(index uint8) Color
func RGB(r, g, b uint8) Color
func Hex(hex string) Color // e.g. "#FF0000"
```

### Style Builder

```go
// NewStyle creates a new, empty style.
func NewStyle() Style

// Methods on Style (chainable)
func (s Style) Foreground(c Color) Style
func (s Style) Background(c Color) Style
func (s Style) Bold(enabled bool) Style
func (s Style) Italic(enabled bool) Style
func (s Style) Underline(enabled bool) Style
func (s Style) Faint(enabled bool) Style     // Dim
func (s Style) Blink(enabled bool) Style
func (s Style) Reverse(enabled bool) Style
func (s Style) Strikethrough(enabled bool) Style

// Render applies the style to a string.
func (s Style) Render(text string) string
```

### Cursor & Screen Utilities

```go
// Cursor
func CursorUp(n int) string
func CursorDown(n int) string
func CursorForward(n int) string
func CursorBack(n int) string
func CursorPosition(row, col int) string
func HideCursor() string
func ShowCursor() string

// Screen
func ClearScreen() string
func ClearLine() string
func ClearLineLeft() string
func ClearLineRight() string

// Buffer
func AltScreenBuffer() string
func NormalScreenBuffer() string
```

### String Utilities

```go
// Strip removes all ANSI escape sequences from a string.
func Strip(s string) string

// VisibleLen returns the length of the string excluding ANSI codes.
func VisibleLen(s string) int
```

## Implementation Guidelines

*   **Performance:** Use `strings.Builder` for string concatenation where appropriate.
*   **Safety:** Ensure `Strip` handles malformed sequences gracefully (don't panic).
*   **Constants:** Define ANSI codes as constants rather than magic strings internally.

## Migration Strategy

1.  Implement this module in `internal/ansi`.
2.  Replace all imports of `github.com/charmbracelet/lipgloss` and `github.com/charmbracelet/x/ansi` in the project with `orchard-gpu-fleet/internal/ansi`.
3.  Remove the external `charmbracelet` dependencies from `go.mod`.
