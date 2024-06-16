package text

import (
	"strconv"
	"strings"
)

const (
	o_seq byte   = '\x1b'
	csi   string = "["
	sgr   string = "m"
)

type TextStyle int

const (
	Style_Reset            TextStyle = iota      // 0
	Style_Bold                                   // 1
	Style_Faint                                  // 2
	Style_Italic                                 // 3
	Style_Underline                              // 4
	Style_Blink                                  // 5
	_                                            // 6
	Style_Inverse                                // 7
	Style_Hidden                                 // 8
	Style_StrikeThrough                          // 9
	Style_FaintEnd         TextStyle = iota + 12 // 22
	Style_BoldEnd          TextStyle = iota + 11 // 22
	Style_ItalicEnd                              // 23
	Style_UnderlineEnd                           // 24
	Style_BlinkEnd                               // 25
	_                                            // 26
	Style_InverseEnd                             // 27
	Style_HiddenEnd                              // 28
	Style_StrikeThroughEnd                       // 29
)

// EscPrefix prepends ASCII escape sequence ESC + [  to the supplied <text>
func EscPrefix(text string) string {
	var sb strings.Builder
	sb.WriteByte(o_seq)
	sb.WriteString(csi)
	sb.WriteString(text)
	return sb.String()
}

// Escp wraps the supplied text with an ASCII escape sequence ESC + [ <text> + m
func Escp(text string) string {
	var sb strings.Builder
	sb.WriteByte(o_seq)
	sb.WriteString(csi)
	sb.WriteString(text)
	sb.WriteString(sgr)
	return sb.String()
}

// Escpi converts the int value to string & then wraps it with an ASCII escape sequence ESC + [ <text> + m
func Escpi(val int) string {
	return Escp(strconv.Itoa(val))
}

// Italicize wraps supplied text in italicize & reset sequence
func Italicize(text string) string {
	return Stylize(Style_Italic, Style_ItalicEnd, text)
}

// It is shortcut for Italicize
func It(text string) string { return Italicize(text) }

// Bolden wraps supplied text in bold & reset sequence
func Bolden(text string) string {
	return Stylize(Style_Bold, Style_BoldEnd, text)
}

// Weaken wraps supplied text in weak/faint & reset sequence
func Weaken(text string) string {
	return Stylize(Style_Faint, Style_FaintEnd, text)
}

// Bol is shortcut for Bolden
func Bol(text string) string { return Bolden(text) }

// Underline wraps supplied text in underline & reset sequence
func Underline(text string) string {
	return Stylize(Style_Underline, Style_UnderlineEnd, text)
}

// Ul is shortcut for Underline
func Ul(text string) string { return Underline(text) }

// Blink wraps supplied text in blink & reset sequence
func Blink(text string) string {
	return Stylize(Style_Blink, Style_BlinkEnd, text)
}

// Hidden wraps supplied text in hide & reset sequence
func Hidden(text string) string {
	return Stylize(Style_Hidden, Style_HiddenEnd, text)
}

// Strikethrough wraps supplied text in striketrhough & reset sequence
func Strikethrough(text string) string {
	return Stylize(Style_StrikeThrough, Style_StrikeThroughEnd, text)
}

// Strike is shorthand for Strikethrough
func Strike(text string) string {
	return Strikethrough(text)
}

// Stylize the supplied text by added start & reset styling sequences
func Stylize(start, reset TextStyle, text string) string {
	var sb strings.Builder
	sb.WriteString(Escpi(int(start)))
	sb.WriteString(text)
	sb.WriteString(Escpi(int(reset)))
	return sb.String()
}
