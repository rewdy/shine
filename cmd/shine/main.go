package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/spf13/pflag"
)

const (
	// Version
	version = "0.0.2"

	// Env var to use for start color
	startEnvVar = "SHINE_START_COLOR"

	// Env var to use for end color
	endEnvVar = "SHINE_END_COLOR"

	envVarRandom = "SHINE_RANDOM"

	// Default start color
	masterDefaultStart = "pink"

	// Default end color
	masterDefaultEnd = "blue"
)

//
// HELPERS
//

// helper function to get all keys from a color map
func getKeys(m map[string][3]int) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// helper function to get a random integer between min and max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// helper function to return n number of spaces
func spaces(n int) string {
	return fmt.Sprintf("%*s", n, " ")
}

// helper function to print the color options for the --help command
func printColorsOptions(colors map[string][3]int) {
	var colorStrings = make([][2]string, len(colors))
	maxLength := 0
	index := 0
	for color, rgb := range colors {
		// Update maxLength so we know how wide to make our columns
		if len(color) > maxLength {
			maxLength = len(color)
		}
		// Now get color values
		r := rgb[0]
		g := rgb[1]
		b := rgb[2]

		// Colorize the text to print
		colorStrings[index] = [2]string{color, fmt.Sprintf("\033[48;2;%d;%d;%dm %v \033[0m", r, g, b, color)}
		// countie countie
		index++
	}

	colWidth := maxLength + 4
	numCols := 5

	fmt.Print("Color swatches for you to choose from:\n\n")

	// Print the values in columns
	for i, value := range colorStrings {
		// Print the value with padding
		padLength := colWidth - len(value[0])
		fmt.Printf("%s%s", value[1], spaces(padLength))

		// Print a newline after every 'columns' values
		if (i+1)%numCols == 0 {
			fmt.Println()
			fmt.Println()
		}
	}

	// Print a newline if the last row is not complete
	if len(colorStrings)%numCols != 0 {
		fmt.Println()
	}
}

// Function to split string into characters, respecting unicode
// so we can keep emojis!!!! xoxo 💋
func splitByRunes(s string) []string {
	var result []string
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		result = append(result, string(r))
		s = s[size:]
	}
	return result
}

// Helper to returns a string with spaced added around it or not
func padOrNot(text string, pad bool) string {
	if pad {
		return " " + text + " "
	}
	return text
}

//
// REAL STUFF
//

// The main gradient function
func generateGradient(text string, startColor, endColor [3]int) string {
	// calculate the color of the steps based on length of text
	chars := splitByRunes(text)
	steps := len(chars)
	gradientText := ""

	for i := 0; i < steps; i++ {
		// Calculate the current color
		r := startColor[0] + (endColor[0]-startColor[0])*i/(steps-1)
		g := startColor[1] + (endColor[1]-startColor[1])*i/(steps-1)
		b := startColor[2] + (endColor[2]-startColor[2])*i/(steps-1)

		// Append the current character with the background color and reset the text color
		gradientText += fmt.Sprintf("\033[48;2;%d;%d;%dm%v\033[0m", r, g, b, chars[i])
	}

	return gradientText
}

// Main function
func main() {
	// The colors we predefine
	colors := map[string][3]int{
		"black":     {0, 0, 0},
		"blue":      {20, 124, 234},
		"brown":     {139, 69, 19},
		"forest":    {0, 97, 42},
		"gold":      {255, 215, 0},
		"gray":      {128, 128, 128},
		"green":     {15, 158, 0},
		"indigo":    {69, 0, 255},
		"lime":      {59, 215, 0},
		"navy":      {0, 0, 128},
		"orange":    {255, 165, 0},
		"persimmon": {236, 88, 0},
		"pink":      {255, 5, 234},
		"purple":    {128, 0, 128},
		"red":       {255, 5, 32},
		"salmon":    {250, 128, 114},
		"teal":      {20, 124, 127},
		"violet":    {108, 47, 157},
		"white":     {255, 255, 255},
		"yellow":    {255, 255, 0},
	}

	// Figure out defaults
	defaultStartColor := masterDefaultStart
	defaultEndColor := masterDefaultEnd
	defaultRandom := false
	if os.Getenv(startEnvVar) != "" {
		defaultStartColor = os.Getenv(startEnvVar)
	}
	if os.Getenv(endEnvVar) != "" {
		defaultEndColor = os.Getenv(endEnvVar)
	}
	if os.Getenv(envVarRandom) != "" {
		defaultRandom = true
	}

	// Setup flags
	startColorFlag := pflag.StringP("start-color", "s", defaultStartColor, "The start color of the gradient")
	endColorFlag := pflag.StringP("end-color", "e", defaultEndColor, "The end color of the gradient")
	padFlag := pflag.BoolP("pad", "p", true, "Adds extra padding around the test to let it breathe")
	randomFlag := pflag.BoolP("random", "r", defaultRandom, "Selects random start and end colors")
	versionFlag := pflag.BoolP("version", "v", false, "Prints the version of the program")

	defaultText := "Hello, World!"
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <text>\n\n", os.Args[0])
		pflag.PrintDefaults()
		fmt.Println()
		printColorsOptions(colors)
	}

	// Parse the flags (Note: you can't access flags before this call)
	pflag.Parse()

	// Get the positional args, the text is positional
	args := pflag.Args()

	// Make a list of acceptable color names
	acceptableColors := getKeys(colors)

	// Use the first positional argument as the text input if provided
	text := defaultText
	if len(args) > 0 {
		text = args[0]
	}
	if *versionFlag {
		text = version
	}

	// Lowercase the colors so we don't have to worry about it
	startColorCased := strings.ToLower(*startColorFlag)
	endColorCased := strings.ToLower(*endColorFlag)

	// If the random flag is set, pick random colors. This overrides
	// the values passed in start and end colors/default colors
	if *randomFlag {
		startColorCased = acceptableColors[randomInt(0, len(acceptableColors))]
		endColorCased = acceptableColors[randomInt(0, len(acceptableColors))]
	}

	// Use the built-in defaults always for version
	if *versionFlag {
		startColorCased = masterDefaultStart
		endColorCased = masterDefaultEnd
	}

	// Get the start and end color rgb values
	startColor := colors[startColorCased]
	endColor := colors[endColorCased]

	// Prepare the text
	preparedText := padOrNot(text, *padFlag)

	// Generate the gradient text
	gradientText := generateGradient(preparedText, startColor, endColor)

	// Print the gradient text
	fmt.Println(gradientText)
}
