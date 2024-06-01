package asciifunc

import (
	"fmt"
	"regexp"
	"strings"
)

/*
The IsColorFlagValid function checks if a given string is a valid color flag in the format --color=X, where X can be any single letter (either uppercase or lowercase).
1.Convert the input string mycolorflag to lowercase to ensure case-insensitivity.
2.Use a regular expression to check if the string matches the pattern --color=([a-z]|[A-Z]).
3.If there is an error during the regex matching, print the error and return false.
4.Return true if the string matches the pattern, otherwise return false.
*/
func IsColorFlagValid(mycolorflag string) bool {
	mycolorflag = strings.ToLower(mycolorflag)
	match, err := regexp.MatchString("--color=([a-z]|[A-Z])", mycolorflag)
	if err != nil {
		fmt.Printf("WE HAVE THIS ERROR %v", err)
		return false
	}

	return match
}

/*The ColorType function takes a string representing a color name, checks if it is a valid color from a predefined list, and returns the corresponding ANSI escape code for that color. If the color is not found, it returns an error message.
1.Extract the actual color name from the input string using the FindColorType function.
2.Convert the extracted color name to lowercase to ensure case-insensitivity.
3.Check if the color exists in the predefined map of color names to ANSI codes.
4.If the color is not found, return an error message.
5.If the color is found, return the corresponding ANSI escape code.
*/

func ColorType(mycolors string) string {
	mycolors = FindColorType(mycolors)
	mymap := map[string]string{
		"reset":         "\033[0m",
		"black":         "\033[30m",
		"red":           "\033[31m",
		"green":         "\033[32m",
		"yellow":        "\033[33m",
		"blue":          "\033[34m",
		"magenta":       "\033[35m",
		"cyan":          "\033[36m",
		"gray":          "\033[37m",
		"white":         "\033[97m",
		"brightblack":   "\033[90m",
		"brightred":     "\033[91m",
		"brightgreen":   "\033[92m",
		"brightyellow":  "\033[93m",
		"brightblue":    "\033[94m",
		"brightmagenta": "\033[95m",
		"brightcyan":    "\033[96m",
		"lightgray":     "\033[37;1m",
		"darkgray":      "\033[90;1m",
		"brightwhite":   "\033[97;1m",
		"orange":        "\033[38;5;208m",
		"pink":          "\033[38;5;205m",
		"purple":        "\033[38;5;171m",
		"lightpurple":   "\033[38;5;183m",
		"darkblue":      "\033[38;5;21m",
		"lightblue":     "\033[38;5;39m",
		"teal":          "\033[38;5;30m",
		"lightteal":     "\033[38;5;51m",
		"lightgreen":    "\033[38;5;118m",
		"darkgreen":     "\033[38;5;22m",
		"brown":         "\033[38;5;94m",
		"yellowgreen":   "\033[38;5;106m",
		"turquoise":     "\033[38;5;80m",
		"gold":          "\033[38;5;178m",
		"silver":        "\033[38;5;7m",
		"skyblue":       "\033[38;5;117m",
	}
	mycolors = strings.ToLower(mycolors)
	IscolorPresent := func(color string) bool {
		_, ok := mymap[mycolors]
		return ok
	}

	if !IscolorPresent(mycolors) {
		return "WRONG SPELLING OF COLOR OR COLOR IS NOT THE PROGRAM"
	}
	return mymap[mycolors]
}

/*The function takes a single string input m, which is expected to contain an "=" character.
1.The function identifies the position of the first "=" character in the input string using strings.Index.
2.It then extracts and returns the substring that starts immediately after the "=" character.*/

func FindColorType(m string) string {
	num := strings.Index(m, "=")

	return m[num+1:]
}
