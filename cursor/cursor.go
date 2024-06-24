package cursor

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/esiddiqui/tfx/text"
)

// Home moves cursor to 0,0 position
func Home() {
	fmt.Print(text.EscPrefix("H"))
}

// goto moves the cursor to line,col position on the terminal
func Goto(line, col int) {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(line))
	sb.WriteString(";")
	sb.WriteString(strconv.Itoa(col))
	sb.WriteString("f")
	fmt.Print(text.EscPrefix(sb.String()))
}

// Up moves cursor up by n lines
func Up(n int) {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(n))
	sb.WriteString("A")
	fmt.Print(text.EscPrefix(sb.String()))
}

// Down moves cursor down by n lines
func Down(n int) {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(n))
	sb.WriteString("B")
	fmt.Print(text.EscPrefix(sb.String()))
}

// Rright moves cursor right by n cols
func Right(n int) {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(n))
	sb.WriteString("C")
	fmt.Print(text.EscPrefix(sb.String()))
}

// Left moves cursor left by n cols
func Left(n int) {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(n))
	sb.WriteString("D")
	fmt.Print(text.EscPrefix(sb.String()))
}

// DonwCr moves cursor down by n lines and to the beginning of that line
func DownCr(n int) {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(n))
	sb.WriteString("E")
	fmt.Print(text.EscPrefix(sb.String()))
}

// Up moves cursor by n lines & to the beginning of that line
func UpCr(n int) {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(n))
	sb.WriteString("F")
	fmt.Print(text.EscPrefix(sb.String()))
}

// Col moves cursor to col n
func Col(n int) {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(n))
	sb.WriteString("G")
	fmt.Print(text.EscPrefix(sb.String()))
}

// TODO fix this
func On() {
	// ESC[?25h
	fmt.Print(text.EscPrefix("?25h"))
}

// TODO fix this
func Off() {
	// ESC[?25l
	fmt.Print(text.EscPrefix("?25l"))
}

// ClearDisplay clears everything on display viewport of the terminal
func ClearDisplay() { fmt.Print(text.EscPrefix("2J")) }

// ClearToEndScreen erases from cursor to end of screen
func ClearToEndScreen() { fmt.Print(text.EscPrefix("0J")) }

// ClearToStartScreen erases from cursor to start of screen
func ClearToStartScreen() { fmt.Print(text.EscPrefix("1J")) }

// Clear the screen; alias for ClearDisplay
func Clear() { ClearDisplay() }

// Cls is a short hand for Clear
func Cls() { Clear() }

// Clear saved lines
func ClearSavedLines() { fmt.Print(text.EscPrefix("3J")) }

// ClearToEndOfLine erases from cursor to end of current line
func ClearToEndOfLine() { fmt.Print(text.EscPrefix("0K")) }

// ClearToStartOfLine ereases from cursor to beginning of current line
func ClearToStartOfLine() { fmt.Print(text.EscPrefix("1K")) }

// ClearLine erases the current line
func ClearLine() { fmt.Print(text.EscPrefix("2K")) }
