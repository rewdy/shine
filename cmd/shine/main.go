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
	acceptableColorsStr := strings.Join(acceptableColors, ", ") + "."

	startColorFlag := pflag.StringP("start-color", "s", "red", "The start color of the gradient. Can be: "+acceptableColorsStr)
	endColorFlag := pflag.StringP("end-color", "e", "blue", "The end color of the gradient. Can be: "+acceptableColorsStr)
	padFlag := pflag.BoolP("pad", "p", true, "Adds extra padding around the test to let it breathe")
	randomFlag := pflag.BoolP("random", "r", false, "Selects random start and end colors")

	defaultText := "Hello, World!"
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <text>\n", os.Args[0])
		pflag.PrintDefaults()
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
