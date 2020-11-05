package utils

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	red        = color.New(color.FgRed)
	green      = color.New(color.FgGreen)
	yellow     = color.New(color.FgYellow)
	redOnWhite = color.New(color.FgRed).Add(color.BgWhite)
)

// Red is used for printing in red
func Red(x ...interface{}) {
	red.Println(x...)
}

// Green is used for printing in green
func Green(x ...interface{}) {
	green.Println(x...)
}

// Yellow is used for printing in yellow
func Yellow(x ...interface{}) {
	yellow.Println(x...)
}

// RedOnWhite is used for printing in red on white background
func RedOnWhite(x ...interface{}) {
	redOnWhite.Println(x...)
}

// Print is for regular printing
func Print(x ...interface{}) {
	fmt.Println(x...)
}

// DisableColor disables colored output
func DisableColor() {
	color.NoColor = true // disables colorized output
}
