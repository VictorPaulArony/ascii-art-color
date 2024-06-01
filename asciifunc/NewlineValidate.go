package asciifunc

import "strings"

/*
The function IsTheStringConsistOfOnlyNewLines checks if a given string consists solely of the sequence \n repeated one or more times.
The function first checks if the input string m is exactly equal to \\n.
If not, it initializes two counters, countL for backslashes (\\) and CountN for the letter n.
It splits the string into individual characters and iterates over them, updating the counters based on the character type.
After the loop, it checks if the sum of both counters equals the length of the string and if the counts of both characters are equal.
Returns true if both conditions are met, otherwise returns false.
*/
func IsTheStringConsistOfOnlyNewLines(m string) bool {
	if m == "\\n" {
		return true
	}

	countL := 0
	CountN := 0

	if len(m) == 0 {
		return false
	}

	mystring := strings.Split(m, "")

	for i := 0; i < len(mystring); i++ {
		if mystring[i] == "n" {
			CountN++
		} else if mystring[i] == "\\" {
			countL++
		}
	}

	if countL+CountN == len(m) && countL == CountN {
		return true
	}

	return false
}
