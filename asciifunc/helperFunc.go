package asciifunc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
AsciiConverter processes a string to convert it into ASCII art, optionally coloring substrings, and handles new lines correctly.

Inputs:
- inputs: The string to be converted into ASCII art.
- filename: The name of the file containing ASCII representations of characters.
- substringToColor: The substring within inputs that should be highlighted with a specific color.
- ColorType: The color to apply to the substringToColor.

Example Usage:
AsciiConverter("Hello\nWorld", "fontfile.txt", "lo", "red")
This will process "Hello" and "World" as separate inputs, color the substring "lo" in red, using ASCII art representation from "fontfile.txt".

Flow:
1. Check if the inputs string contains the newline character (\n), and replace it with an actual newline character (\n) if present.
2. If new lines are present, split the inputs into separate lines and process each line individually.
3. For each line (or the entire inputs if no new lines), find the indices where the substringToColor appears.
4. Call CreateAscii for each line or the whole input to generate and optionally color the ASCII art.
5. Handle empty lines specifically by just printing a new line.

Outputs:
- The function does not return any values but outputs the generated ASCII art directly to the console, with specified substrings colored as per ColorType.
*/
func AsciiConverter(inputs, filename, substringToColor, ColorType string) {
	isNewLinePresent := strings.Contains(inputs, "\\n")

	if isNewLinePresent {
		inputs = strings.ReplaceAll(inputs, "\\n", "\n")
	}

	if isNewLinePresent {

		myword := strings.Split(inputs, "\n")

		// var ascii string
		for i := 0; i < len(myword); i++ {
			if len(myword[i]) == 0 {
				fmt.Println()
			} else {
				indexesToBeColored := TextToColor(substringToColor, myword[i])

				CreateAscii(myword[i], filename, ColorType, indexesToBeColored)
			}
		}

	} else {
		indexesToBeColored := TextToColor(substringToColor, inputs)
		CreateAscii(inputs, filename, ColorType, indexesToBeColored)
	}
}

/*
CREATEASCII func
The CreateAscii function generates ASCII art from a given string, using specified line numbers from a file, and applies a color to selected characters.

1.Calculate line numbers for each character in the string n using getLineNumbers.
2.For each line number, retrieve corresponding lines from the file using PrintLine to construct the ASCII representation.
3.Organize these lines into a 2D slice to represent the complete ASCII art.
4.Apply the specified color to the ASCII art at the indices provided.
5.The function iterates through 9 lines for each character to construct the full ASCII representation.

*/

func CreateAscii(n, filename, color string, indexesOfColorToBePrinted []int) {
	numbers := getLineNumbers(n)

	var asciiLines []string
	var asciiArt [][]string
	counterLInes := 0
	moveToNextLine := 0
	for counterLInes <= 8 {
		for i := 0; i < len(numbers); i++ {
			asciiLines = append(asciiLines, PrintLine(numbers[i]+moveToNextLine, filename))
		}
		asciiArt = append(asciiArt, asciiLines)
		asciiLines = []string{}
		moveToNextLine++
		counterLInes++
	}
	ApplyColorToString(asciiArt, color, indexesOfColorToBePrinted)
}

/*
GETLINENUMBERS FUNCTION
The getLineNumbers function calculates the starting line numbers in an ASCII art file for each character in a given string based on its ASCII value.
*/
func getLineNumbers(word string) []int {
	var lineNumbers []int

	for _, char := range word {
		b := int(char-32)*9 + 2
		lineNumbers = append(lineNumbers, b)
	}

	return lineNumbers
}

/**
PrintLine reads a specific line from a file based on the line number provided.

It opens the specified file, iterates through each line until it reaches the desired line number,
and returns that line as a string.

Parameters:
- num (int): The line number to be retrieved from the file.
- filename (string): The name of the file from which to read the line.

Returns:
- string: The content of the line at the specified line number.

Example:
	lineContent := PrintLine(5, "ascii_art.txt")
	fmt.Println(lineContent) // Outputs the fifth line of the file "ascii_art.txt"
*/

func PrintLine(num int, filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic("Error opening the file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	linetoPrint := 0
	var line string
	for scanner.Scan() {
		linetoPrint++
		if linetoPrint == num {
			line = scanner.Text()
		}
	}

	if err := scanner.Err(); err != nil {
		panic("we have an error")
	}
	return line
}

/*
The ApplyColorToString function iterates over a matrix of strings, applying a specified color to certain indices of each row, and prints each colored row to the console.
FLOW
1.Iterate over each row in the matrix m.
2.For each row, except the last one, call the ColorMyText function to apply the color to specified indices.
3.Print the colored row followed by a newline
*/
func ApplyColorToString(m [][]string, colour string, indextobe []int) {
	for i := 0; i < len(m); i++ {
		if i < len(m)-1 {

			ColorMyText(indextobe, m[i], colour)
			fmt.Println()
		}
	}
}
