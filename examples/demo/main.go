package main

import (
	"fmt"
	"time"

	"github.com/orchard9/get-ansi"
)

func main() {
	// 1. Basic Colors
	fmt.Println(ansi.NewStyle().Bold(true).Underline(true).Render("--- Basic Colors ---"))
	colors := []struct {
		name string
		c    ansi.Color
	}{
		{"Red", ansi.Red()},
		{"Green", ansi.Green()},
		{"Blue", ansi.Blue()},
		{"Yellow", ansi.Yellow()},
		{"Magenta", ansi.Magenta()},
		{"Cyan", ansi.Cyan()},
		{"White", ansi.White()},
		{"Black", ansi.Black()},
	}

	for _, c := range colors {
		s := ansi.NewStyle().Foreground(c.c)
		fmt.Printf("%s ", s.Render(c.name))
	}
	fmt.Println()

	// 2. Advanced Colors
	fmt.Println(ansi.NewStyle().Bold(true).Underline(true).Render("--- Advanced Colors ---"))
	fmt.Println(ansi.NewStyle().Foreground(ansi.RGB(255, 165, 0)).Render("Orange (RGB 255,165,0)"))
	fmt.Println(ansi.NewStyle().Foreground(ansi.Hex("#AA00FF")).Render("Purple (Hex #AA00FF)"))
	fmt.Println(ansi.NewStyle().Foreground(ansi.Color256(208)).Render("Color 208 (256 Palette)"))
	fmt.Println()

	// 3. Styles
	fmt.Println(ansi.NewStyle().Bold(true).Underline(true).Render("--- Text Styles ---"))
	fmt.Println(ansi.NewStyle().Bold(true).Render("Bold Text"))
	fmt.Println(ansi.NewStyle().Faint(true).Render("Faint Text"))
	fmt.Println(ansi.NewStyle().Italic(true).Render("Italic Text"))
	fmt.Println(ansi.NewStyle().Underline(true).Render("Underlined Text"))
	fmt.Println(ansi.NewStyle().Blink(true).Render("Blinking Text"))
	fmt.Println(ansi.NewStyle().Reverse(true).Render("Reversed Text"))
	fmt.Println(ansi.NewStyle().Strikethrough(true).Render("Strikethrough Text"))
	fmt.Println(ansi.NewStyle().Hidden(true).Render("Hidden Text (You shouldn't see this)"))
	fmt.Println(" (End of hidden text)")
	fmt.Println()

	// 4. Backgrounds
	fmt.Println(ansi.NewStyle().Bold(true).Underline(true).Render("--- Backgrounds ---"))
	fmt.Println(ansi.NewStyle().Background(ansi.Red()).Foreground(ansi.White()).Render(" White on Red "))
	fmt.Println(ansi.NewStyle().Background(ansi.Blue()).Foreground(ansi.Yellow()).Bold(true).Render(" Bold Yellow on Blue "))
	fmt.Println()

	// 5. Cursor & Screen (Simulated)
	fmt.Println(ansi.NewStyle().Bold(true).Underline(true).Render("--- Cursor & Screen Utilities ---"))
	fmt.Printf("Strip Test: '%s' (len: %d)\n", ansi.Strip("\x1b[31mHello\x1b[0m"), ansi.VisibleLen("\x1b[31mHello\x1b[0m"))

	fmt.Println("Simulating progress bar (Cursor Movement)...")
	ansi.HideCursor()
	defer fmt.Print(ansi.ShowCursor()) // Ensure cursor returns

	for i := 0; i <= 10; i++ {
		width := 20
		filled := int(float64(i) / 10.0 * float64(width))
		bar := ""
		for j := 0; j < width; j++ {
			if j < filled {
				bar += "="
			} else {
				bar += "-"
			}
		}

		fmt.Printf("\r[%s] %d%%", bar, i*10)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println("\nDone!")
}
