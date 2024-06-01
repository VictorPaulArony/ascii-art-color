package asciifunc

import (
	"fmt"
	"strings"
)

/*
The ColorMyText function colors specific indices of text in a list based on a given list of indices and a specified color code.
1.iterate through each element in the original list.
2.Check if the current index is in the IndexToColor list using the IsItAnIndexToColor function.
3.If the index is in the list, print the text at that index with the specified color and reset the color formatting after the text.
4.If the index is not in the list, print the text without any color formatting.
*/
//var buf bytes.Buffer

func ColorMyText(IndexToColor []int, original []string, color string) {
	for i := 0; i < len(original); i++ {
		if IsItAnIndexToColor(i, IndexToColor) {
			fmt.Print(color + original[i] + "\033[0m")
		} else {
			fmt.Print(original[i])
		}
	}
}

/*
TextToColor function takes two strings, colorttext and original, and returns a list of indices where characters from colorttext are found in the original string.
*/

func TextToColor(colorttext, original string) []int {
	var myIndex []int
	mysliceoriginal := strings.Split(original, "")
	MySliceColorText := strings.Split(colorttext, "")
	for i := 0; i < len(mysliceoriginal); i++ {
		for j := 0; j < len(MySliceColorText); j++ {
			if mysliceoriginal[i] == MySliceColorText[j] {
				myIndex = append(myIndex, i)
			}
		}
	}

	return myIndex
}

/*The IsItAnIndexToColor function checks if a given integer is present in a list of integers.  */
func IsItAnIndexToColor(num int, m []int) bool {
	for i := 0; i < len(m); i++ {
		if num == m[i] {
			return true
		}
	}
	return false
}

/*The CheckWordIsASubstring function determines if any character of a substring is present in a given word string
1.If subString is empty, the function immediately returns true.
2.The function splits subString into individual characters.
3.It iterates through each character of subString.
4.For each character, it checks if this character is contained in wordString.
5.If any character is found in wordString, the function returns true; if none are found, it returns false after the loop.
*/

func CheckWordIsASubstring(subString, wordString string) bool {
	if len(subString) == 0 {
		return true
	}

	subStrings := strings.Split(subString, "")
	for i := 0; i < len(subStrings); i++ {
		if strings.Contains(wordString, subStrings[i]) {
			return true
		}
	}
	return false
}
