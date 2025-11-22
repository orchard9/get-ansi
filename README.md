# Get Ansi (`get-ansi`)

> A zero-dependency, lightweight ANSI escape sequence library for Go.

`get-ansi` provides a clean, fluent API for terminal styling, colors (True Color support), and cursor/screen control. It is designed to be a drop-in, standard-library-only alternative to heavier TUI libraries for basic styling needs.

## Features

- **Zero Dependencies**: Pure Go standard library.
- **Colors**: 4-bit (16 colors), 8-bit (256 colors), and 24-bit (True Color RGB).
- **Styling**: Bold, Dim, Italic, Underline, Blink, Reverse, Hidden, Strikethrough.
- **Control**: Cursor movement, visibility, and screen clearing.
- **Utilities**: Strip ANSI codes, calculate visible string length.

## Installation

```bash
go get github.com/orchard9/get-ansi
```

## Usage

### Styling & Colors

```go
package main

import (
	"fmt"
	"github.com/orchard9/get-ansi"
)

func main() {
    // Fluent Style Builder
    style := ansi.NewStyle().
        Foreground(ansi.Red()).
        Background(ansi.Black()).
        Bold(true).
        Underline(true)

    fmt.Println(style.Render("Hello, Styled World!"))

    // RGB & Hex Colors
    orange := ansi.NewStyle().Foreground(ansi.RGB(255, 165, 0))
    fmt.Println(orange.Render("True Color Orange"))
    
    purple := ansi.NewStyle().Foreground(ansi.Hex("#AA00FF"))
    fmt.Println(purple.Render("Hex Color Purple"))
}
```

### Cursor & Screen Control

```go
// Move cursor up 2 lines
fmt.Print(ansi.CursorUp(2))

// Clear the entire screen
fmt.Print(ansi.ClearScreen())

// Hide/Show cursor
fmt.Print(ansi.HideCursor())
defer fmt.Print(ansi.ShowCursor())
```

### String Utilities

```go
text := "\x1b[31mHello\x1b[0m"

// Strip ANSI codes -> "Hello"
clean := ansi.Strip(text)

// Get visible length -> 5
length := ansi.VisibleLen(text)
```

### Advanced Composition

When using `get-ansi` with other TUI libraries (like `bubbles/table`) that handle their own styling or background colors, using `Render()` might cause issues because it appends a `Reset` code, which can clear the external library's background style (causing "gaps" in selection bars).

To avoid this, use `String()` to get the escape codes *without* the reset, and manually concatenate your text:

```go
// Instead of:
// style.Render("Content") // -> Codes + Content + Reset

// Use:
cell := style.String() + "Content" 
// -> Codes + Content (No Reset)
// The external library will handle the reset at the end of the line/block.
```

## License

MIT
