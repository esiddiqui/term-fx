package color

import (
	"strconv"
	"strings"

	"github.com/esiddiqui/term-fx/text"
)

const (
	Fg256_Reset int = iota
)

const (
	Bg256_Reset int = iota
)

// Color256 wraps the supplied text into 256 color background & foreground escape & reset sequences
func Color256(bg, fg int, text string) string {
	var sb strings.Builder
	sb.WriteString(get256Bg(bg))
	sb.WriteString(get256Fg(fg))
	sb.WriteString(text)
	sb.WriteString(get256Fg(Fg256_Reset))
	sb.WriteString(get256Bg(Bg256_Reset))
	return sb.String()
}

func ColorBg256(bg int, text string) string {
	var sb strings.Builder
	sb.WriteString(get256Bg(bg))
	sb.WriteString(text)
	sb.WriteString(get256Bg(Bg256_Reset))
	return sb.String()
}

func ColorFg256(fg int, text string) string {
	var sb strings.Builder
	sb.WriteString(get256Fg(fg))
	sb.WriteString(text)
	sb.WriteString(get256Fg(Fg256_Reset))
	return sb.String()
}

// returns the foreground color esc sequence for the supplied color id
func get256Fg(id int) string {
	var sb strings.Builder
	sb.WriteString("38;5;")
	sb.WriteString(strconv.Itoa(id))
	return text.Escp(sb.String())
}

// returns the background color esc sequence for the supplied color id
func get256Bg(id int) string {
	var sb strings.Builder
	sb.WriteString("48;5;")
	sb.WriteString(strconv.Itoa(id))
	return text.Escp(sb.String())
}
