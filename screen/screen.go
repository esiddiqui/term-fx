package screen

import (
	"fmt"

	"github.com/esiddiqui/term-fx/text"
)

// ScreenMode320_200 sets the screen mode to 320x200 256-color
func ScreenMode320_200() {
	fmt.Print(text.EscPrefix("=19h"))
}

func Clear() {
	fmt.Print(text.EscPrefix("2J"))
}
