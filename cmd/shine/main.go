package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

func GetKeys(m map[string][3]int) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// spaces returns a string of n spaces
func spaces(n int) string {
	return fmt.Sprintf("%*s", n, " ")
}

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

func generateGradient(text string, startColor, endColor [3]int) string {
	// calculate the color of the steps based on length of text
	steps := len(text)
	gradientText := ""

	for i := 0; i < steps; i++ {
		// Calculate the current color
		r := startColor[0] + (endColor[0]-startColor[0])*i/(steps-1)
		g := startColor[1] + (endColor[1]-startColor[1])*i/(steps-1)
		b := startColor[2] + (endColor[2]-startColor[2])*i/(steps-1)

		// Append the current character with the background color and reset the text color
		gradientText += fmt.Sprintf("\033[48;2;%d;%d;%dm%c\033[0m", r, g, b, text[i])
	}

	return gradientText
}

func padOrNot(text string, pad bool) string {
	if pad {
		return " " + text + " "
	}
	return text
}

func main() {
	colors := map[string][3]int{
		"red":       {255, 5, 32},
		"orange":    {255, 165, 0},
		"yellow":    {255, 255, 0},
		"green":     {0, 128, 0},
		"blue":      {20, 124, 234},
		"indigo":    {75, 0, 130},
		"violet":    {238, 130, 238},
		"pink":      {255, 5, 234},
		"teal":      {20, 124, 127},
		"brown":     {139, 69, 19},
		"black":     {0, 0, 0},
		"white":     {255, 255, 255},
		"gray":      {128, 128, 128},
		"salmon":    {250, 128, 114},
		"persimmon": {236, 88, 0},
	}

	acceptableColors := GetKeys(colors)

	startColorFlag := pflag.StringP("start-color", "s", "red", "The start color of the gradient")
	endColorFlag := pflag.StringP("end-color", "e", "blue", "The end color of the gradient")
	padFlag := pflag.BoolP("pad", "p", true, "Adds extra padding around the test to let it breathe")
	randomFlag := pflag.BoolP("random", "r", false, "Selects random start and end colors")

	defaultText := "Hello, World!"
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <text>\n\n", os.Args[0])
		pflag.PrintDefaults()
		fmt.Println()
		printColorsOptions(colors)
	}

	// Parse the flags
	pflag.Parse()

	// Get the positional args
	args := pflag.Args()

	startColorCased := strings.ToLower(*startColorFlag)
	endColorCased := strings.ToLower(*endColorFlag)

	// Use the first positional argument as the text input if provided
	text := defaultText
	if len(args) > 0 {
		text = args[0]
	}

	if *randomFlag {
		startColorCased = acceptableColors[randomInt(0, len(acceptableColors))]
		endColorCased = acceptableColors[randomInt(0, len(acceptableColors))]
	}

	// Get the start and end colors
	startColor := colors[startColorCased]
	endColor := colors[endColorCased]

	// Prepare the text
	preparedText := padOrNot(text, *padFlag)

	// Generate the gradient text
	gradientText := generateGradient(preparedText, startColor, endColor)

	// Print the gradient text
	fmt.Println(gradientText)
}
