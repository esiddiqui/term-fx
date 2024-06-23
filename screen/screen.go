package screen

import (
	"fmt"

	"github.com/esiddiqui/tfx/text"
)

// ScreenMode320_200 sets the screen mode to 320x200 256-color
func ScreenMode320_200() {
	fmt.Print(text.EscPrefix("=19h"))
}

// Clear the screen
func Clear() { fmt.Print(text.EscPrefix("2J")) }

// Cls is a short hand for Clear
func Cls() { Clear() }
