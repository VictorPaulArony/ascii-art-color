package asciifunc

import (
	"bufio"
	"fmt"
	"os"
)

/*
Filenamevalidate
 The Filenamevalidate function checks if a given filename exists among predefined options and returns the correct filename format.
*/

func Filenamevalidate(m string) string {
	if !filenameExist(m) {
		return "FILE DOES NOT EXSIT OR CHECK THE SPELLING OF YOUR FILES (e.g standard.txt)"
	}

	if m == "shadow" || m == "shadow.txt" {
		return "shadow.txt"
	} else if m == "thinkertoy" || m == "thinkertoy.txt" {
		return "thinkertoy.txt"
	} else {
		return "standard.txt"
	}
}

func filenameExist(m string) bool {
	return m == "shadow.txt" || m == "thinkertoy.txt" || m == "standard.txt"
}

/*
CheckFileEmpty
The CheckFileEmpty function verifies if a specified file is empty or not by checking its content length
*/

func CheckFileEmpty(m string) bool {
	file, err := os.Open(m)
	if err != nil {
		// fmt.Println("error opening the file")
		return false
	}

	defer file.Close()

	content, err := os.ReadFile(m)
	if err != nil {
		fmt.Println("error reading the file")
		return false
	}

	t := string(content)

	if len(t) == 0 || (len(t) == 1 && t == "") {
		return false
	}
	return true
}

/*
CheckNumberOfLinesInTheFile
The CheckNumberOfLinesInTheFile function verifies if a file contains exactly 855 lines, returning true if it does, and false otherwise.
*/
func CheckNumberOfLinesInTheFile(m string) bool {
	file, err := os.Open(m)
	if err != nil {
		fmt.Println("error opening the file")
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	countLines := 0

	for scanner.Scan() {
		countLines++
		if countLines > 855 {
			break
		}

	}

	return countLines == 855
}

/*
IsItAnAsciiCharacter func

	The IsItAnAsciiCharacter function checks if all characters in a given string are valid ASCII characters, specifically within the printable range from 32 to 126 inclusive.
*/
func IsItAnAsciiCharacter(m string) bool {
	for _, word := range m {
		if word < 32 || word > 126 {
			return false
		}
	}
	return true
}
