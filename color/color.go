package color

import (
	"strconv"
	"strings"

	"github.com/esiddiqui/term-fx/text"
)

const (
	Fg256_Reset int = iota // 0
)

const (
	Bg256_Reset int = iota // 0
)

const (
	Bg16_Reset   int = iota      // 0
	Bg16_Black   int = iota + 39 // 30
	Bg16_Red                     // 31
	Bg16_Green                   // 32
	Bg16_Yellow                  // 33
	Bg16_Blue                    // 34
	Bg16_Magenta                 // 35
	Bg16_Cyan                    // 36
	Bg16_White                   // 37
	_                            // 38
	Bg16_Default                 // 39
)

const (
	Fg16_Reset   int = iota      // 0
	Fg16_Black   int = iota + 29 // 40
	Fg16_Red                     // 41
	Fg16_Green                   // 42
	Fg16_Yellow                  // 43
	Fg16_Blue                    // 44
	Fg16_Magenta                 // 45
	Fg16_Cyan                    // 46
	Fg16_White                   // 47
	_                            // 48
	Fg16_Default                 // 49
)

// Black wraps txt in ascii escape sequence for 16 color black foreground color
func Black(txt string) string { return Fg16(Fg16_Black, txt) }

// Red wraps txt in ascii escape sequence for 16 color red foreground color
func Red(txt string) string { return Fg16(Fg16_Red, txt) }

// Green wraps txt in ascii escape sequence for 16 color green foreground color
func Green(txt string) string { return Fg16(Fg16_Green, txt) }

// Yellow wraps txt in ascii escape sequence for 16 color yellow foreground color
func Yellow(txt string) string { return Fg16(Fg16_Yellow, txt) }

// Blue wraps txt in ascii escape sequence for 16 color blue foreground color
func Blue(txt string) string { return Fg16(Fg16_Blue, txt) }

// Magenta wraps txt in ascii escape sequence for 16 color magenta foreground color
func Magenta(txt string) string { return Fg16(Fg16_Magenta, txt) }

// Cyan wraps txt in ascii escape sequence for 16 color cyan foreground color
func Cyan(txt string) string { return Fg16(Fg16_Cyan, txt) }

// White wraps txt in ascii escape sequence for 16 color white foreground color
func White(txt string) string { return Fg16(Fg16_White, txt) }

// Apply16 wraps the supplied text with 8-bit (256-color) background & reset sequences
func Apply16(bg, fg int, txt string) string {
	var sb strings.Builder
	sb.WriteString("0;") // reset text styling, since we do not want it here...
	sb.WriteString(strconv.Itoa(fg))
	sb.WriteString(";")
	sb.WriteString(strconv.Itoa(bg))

	var xsb strings.Builder
	xsb.WriteString(text.Escp(sb.String()))
	xsb.WriteString(txt)
	xsb.WriteString(text.Escpi(Bg16_Reset))
	xsb.WriteString(text.Escpi(Fg16_Reset))
	return xsb.String()
}

// Bg16 wraps the supplied text with 16-color background & reset sequences
func Bg16(bg int, txt string) string {
	var sb strings.Builder
	sb.WriteString(text.Escpi(bg))
	sb.WriteString(txt)
	sb.WriteString(text.Escpi(Bg16_Reset))
	return sb.String()
}

// Foreground256 wraps the supplied text with 8-bit (256-color) foreground/font & reset sequenes
func Fg16(fg int, txt string) string {
	var sb strings.Builder
	sb.WriteString(text.Escpi(fg))
	sb.WriteString(txt)
	sb.WriteString(text.Escpi(Fg16_Reset))
	return sb.String()
}

// Apply256 wraps the supplied text into 8-bit (256-color) background & foreground escape & reset sequences
func Apply256(bg, fg int, text string) string {
	var sb strings.Builder
	sb.WriteString(get256Bg(bg))
	sb.WriteString(get256Fg(fg))
	sb.WriteString(text)
	sb.WriteString(get256Fg(Fg256_Reset))
	sb.WriteString(get256Bg(Bg256_Reset))
	return sb.String()
}

// Background256 wraps the supplied text with 8-bit (256-color) background & reset sequences
func Background256(bg int, text string) string {
	var sb strings.Builder
	sb.WriteString(get256Bg(bg))
	sb.WriteString(text)
	sb.WriteString(get256Bg(Bg256_Reset))
	return sb.String()
}

// Foreground256 wraps the supplied text with 8-bit (256-color) foreground/font & reset sequenes
func Foreground256(fg int, text string) string {
	var sb strings.Builder
	sb.WriteString(get256Fg(fg))
	sb.WriteString(text)
	sb.WriteString(get256Fg(Fg256_Reset))
	return sb.String()
}

// SetBackgroundRgb sets the r,g,b true-color background
func SetBackgroundRgb(r, g, b int) {
	// ESC[38;2;{r};{g};{b}m
	var sb strings.Builder
	sb.WriteString("48;2;")
	sb.WriteString(strconv.Itoa(r))
	sb.WriteString(";")
	sb.WriteString(strconv.Itoa(g))
	sb.WriteString(";")
	sb.WriteString(strconv.Itoa(b))
	print(text.Escp(sb.String()))
	// var xsb strings.Builder
	// xsb.WriteString()
}

// SetForegroundRgb sets the r,g,b true-color foreground
func SetForegroundRgb(r, g, b int) {
	var sb strings.Builder
	sb.WriteString("38;2;")
	sb.WriteString(strconv.Itoa(r))
	sb.WriteString(";")
	sb.WriteString(strconv.Itoa(g))
	sb.WriteString(";")
	sb.WriteString(strconv.Itoa(b))
	print(text.Escp(sb.String()))
}

// returns the foreground color esc sequence for the supplied color id (8 bit, 256 color)
func get256Fg(id int) string {
	var sb strings.Builder
	sb.WriteString("38;5;")
	sb.WriteString(strconv.Itoa(id))
	return text.Escp(sb.String())
}

// returns the background color esc sequence for the supplied color id (8 bit, 256 color)
func get256Bg(id int) string {
	var sb strings.Builder
	sb.WriteString("48;5;")
	sb.WriteString(strconv.Itoa(id))
	return text.Escp(sb.String())
}
