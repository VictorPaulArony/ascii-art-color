package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art/asciifunc"
)

/*
The main function processes command-line arguments to determine the behavior of an ASCII art generator.
It validates input parameters, checks file existence and content, and applies color and ASCII transformations based on specified conditions.

Inputs:
- os.Args[1:]: Array of strings from the command line excluding the program name. Expected to contain color flags, substrings for coloring, and optionally a filename.

Flow:
1. Validates the number of input arguments and checks for correct command-line syntax.
2. Parses and validates the color flag and checks if the substring to be colored is present in the main input string.
3. Handles special cases like strings consisting only of new lines or containing non-ASCII characters.
4. Determines the filename for ASCII art transformation, validates the file's existence, emptiness, and content integrity.
5. Calls the `AsciiConverter` function to generate the colored ASCII art output based on the validated inputs and conditions.

Outputs:
- Prints error messages directly to the console for various validation failures.
- Outputs transformed ASCII art to the console if all validations pass and conditions are met.
*/

func main() {
	myinputs := os.Args[1:]
	var newSlice []string

	if (len(myinputs) < 1 || len(myinputs) > 4) && len(myinputs) != 2 {
		fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		return
	}

	if len(myinputs) == 1 && !strings.Contains(myinputs[0], "--color=") {
		newSlice = append(newSlice, "--color=white", myinputs[0])
		myinputs = newSlice
	} else if len(myinputs) == 1 && strings.Contains(myinputs[0], "--color=") {
		fmt.Println("This is a flag")
		return
	}

	var inputs string
	var inputColor string

	/*
		Conditionerto check whether the passed argumet has the letters affected in the word, the color and the ASCII format
	*/

	if len(myinputs) == 2 {
		if !strings.Contains(myinputs[0], "--color=") {
			newSlice = append(newSlice, "--color=white", myinputs[0], myinputs[1])
			myinputs = newSlice
		}

		inputColor = myinputs[1]
		inputs = myinputs[1]

	} else if len(myinputs) == 3 && (myinputs[2] == "thinkertoy.txt" || myinputs[2] == "standard.txt" || myinputs[2] == "shadow.txt") {
		inputColor = myinputs[1]
		inputs = myinputs[1]

	} else {

		inputs = myinputs[2]
		inputColor = myinputs[1]
	}

	var color string

	if !asciifunc.IsColorFlagValid(myinputs[0]) {
		fmt.Println("Wrong format in writing the color flag")
		fmt.Println("Please refer to the example below in writing color flag")
		fmt.Println("EX: go run . --color=<color> <letters to be colored> 'something'")
		os.Exit(0)
	}

	if asciifunc.ColorType(myinputs[0]) == "WRONG SPELLING OF COLOR OR COLOR IS NOT THE PROGRAM" {
		fmt.Println("WRONG SPELLING OF COLOR OR COLOR IS NOT THE PROGRAM")
		return
	} else {
		color = asciifunc.ColorType(myinputs[0])
	}

	if !asciifunc.CheckWordIsASubstring(inputColor, inputs) {
		fmt.Printf("ERROR: The word to be colored (%v) is not a substring of the string sentence (%v)\n", myinputs[1], inputs)
		return
	}

	if asciifunc.IsTheStringConsistOfOnlyNewLines(inputs) {
		if len(inputs) == 2 {
			fmt.Println()
			return
		}

		numbersOfnewlines := len(inputs) / 2
		count := 0
		for count < numbersOfnewlines {
			fmt.Println()
			count++
		}
		return
	}

	if len(inputs) == 0 {
		return
	}

	// error handling incase of \a,\b,\t,\v,\f,\r(tab functions)
	if strings.Contains(inputs, "\\a") || strings.Contains(inputs, "\\b") || strings.Contains(inputs, "\\t") || strings.Contains(inputs, "\\v") || strings.Contains(inputs, "\\f") || strings.Contains(inputs, "\\r") {
		fmt.Println("ERROR: THE PROGRAM HANDLES AN INPUT WITH NUMBERS, LETTERS, SPACES, SPECIAL CHARACTERS (Special characters are the punctuation characters on your keyboard, such as: ! @ # $ % ^ & * ( ) - _ = + \\ | [ ] { } ; : / ? . >() ) and \\n ONLY.\n REMOVE THIS CHARACTERS (\\a,\\r,\\f,\\v,\\b,\\t) INCASE YOUR STRING HAS THEM. ")
		return
	}
	// deals with the types of new lines
	inputs = strings.ReplaceAll(inputs, "\n", "\\n")
	// checks if the word is in ascii value range
	if !asciifunc.IsItAnAsciiCharacter(inputs) {
		fmt.Println("ERROR:CONTAINS A WORD OUTSIDE ASCII VALUE RANGE")

		return
	}

	/* Gives out the output according to the condition obtained during the input */
	if len(myinputs) == 4 {
		inputfile := myinputs[3]
		filenames := asciifunc.Filenamevalidate(inputfile)

		if filenames == "FILE DOES NOT EXSIT OR CHECK THE SPELLING OF YOUR FILES" {
			fmt.Println("ERROR: " + filenames)
			return
		} else {
			if asciifunc.CheckFileEmpty(filenames) {
				if asciifunc.CheckNumberOfLinesInTheFile(filenames) {
					asciifunc.AsciiConverter(inputs, filenames, inputColor, color)
				} else {
					fmt.Printf("ERROR: %v FILE HAS BEEN MODIFIED \n DO NOT MODIFY THE FILE \n", filenames)
					return
				}
			} else {
				fmt.Printf("ERROR: %v FILE IS EMPTY \n", filenames)
				return
			}
		}

	} else if len(myinputs) == 3 {
		if myinputs[2] == "thinkertoy.txt" || myinputs[2] == "standard.txt" || myinputs[2] == "shadow.txt" {
			inputfile := myinputs[2]
			filenames := asciifunc.Filenamevalidate(inputfile)

			if filenames == "FILE DOES NOT EXSIT OR CHECK THE SPELLING OF YOUR FILES" {
				fmt.Println("ERROR: " + filenames)
				return
			} else {
				if asciifunc.CheckFileEmpty(filenames) {
					if asciifunc.CheckNumberOfLinesInTheFile(filenames) {
						asciifunc.AsciiConverter(inputs, filenames, inputColor, color)
					} else {
						fmt.Printf("ERROR: %v FILE HAS BEEN MODIFIED \n DO NOT MODIFY THE FILE \n", filenames)
						return
					}
				} else {
					fmt.Printf("ERROR: %v FILE IS EMPTY \n", filenames)
					return
				}
			}

		} else {
			if asciifunc.CheckFileEmpty("standard.txt") {
				if asciifunc.CheckNumberOfLinesInTheFile("standard.txt") {
					asciifunc.AsciiConverter(inputs, "standard.txt", inputColor, color)
				} else {
					fmt.Println("ERROR: standard.txt FILE  HAS BEEN MODIFIED \n DO NOT MODIFY THE FILE")
					return
				}
			} else {
				fmt.Println("ERROR: standard.txt FILE IS EMPTY")
				return
			}
		}
	} else if len(myinputs) == 2 {
		if asciifunc.CheckFileEmpty("standard.txt") {
			if asciifunc.CheckNumberOfLinesInTheFile("standard.txt") {
				asciifunc.AsciiConverter(inputs, "standard.txt", inputColor, color)
			} else {
				fmt.Println("ERROR: standard.txt FILE  HAS BEEN MODIFIED \n DO NOT MODIFY THE FILE")
				return
			}
		} else {
			fmt.Println("ERROR: standard.txt FILE IS EMPTY")
			return
		}
	}
}
