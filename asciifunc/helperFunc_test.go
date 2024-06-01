package asciifunc

import (
	"testing"
)

// Function returns true when num is found in the list m
func TestIsItAnIndexToColor(t *testing.T) {
	indexList := []int{1, 2, 3, 4, 5}
	num := 3
	result := IsItAnIndexToColor(num, indexList)
	if !result {
		t.Errorf("Expected true, got %v", result)
	}
} // subString contains characters not present in wordString
func TestCheckWordIsASubstrin(t *testing.T) {
	result := CheckWordIsASubstring("x", "hello")
	if result {
		t.Errorf("Expected false for substring 'x' not in 'hello', got %v", result)
	}
}

// Return error message for non-existent file names
func TestFilenameValidateNonExistentFile(t *testing.T) {
	result := Filenamevalidate("nonexistentfile.txt")
	expected := "FILE DOES NOT EXSIT OR CHECK THE SPELLING OF YOUR FILES (e.g standard.txt)"
	if result != expected {
		t.Errorf("Filenamevalidate(\"nonexistentfile.txt\") = %s; want %s", result, expected)
	}
}

// Input string containing characters below ASCII 32 returns false
func TestIsItAnASciiCharacter(t *testing.T) {
	input := "\x1F" // ASCII character below 32
	result := IsItAnAsciiCharacter(input)
	if result != false {
		t.Errorf("Expected false, got %v for input '%s'", result, input)
	}
}

// func TestAsciiConverter(t *testing.T) {
// 	AsciiConverter("h", "standard.txt", "h", "red")
// 	output := buf.String()
// 	fmt.Println(output)
// 	// expected1 := " _              _   _          \n" +
// 	// 	"| |            | | | |         \n" +
// 	// 	"| |__     ___  | | | |   ___   \n" +
// 	// 	"|  _ \\   / _ \\ | | | |  / _ \\  \n" +
// 	// 	"| | | | |  __/ | | | | | (_) | \n" +
// 	// 	"|_| |_|  \\___| |_| |_|  \\___/  \n" +
// 	// 	"                               \n" + "                               \n"

// 	// 	result2 := "\n" + AsciiConverter("hello\\nThere", "standard.txt")
// 	// 	expected2 := "\n" + " _              _   _          \n" +
// 	// 		"| |            | | | |         \n" +
// 	// 		"| |__     ___  | | | |   ___   \n" +
// 	// 		"|  _ \\   / _ \\ | | | |  / _ \\  \n" +
// 	// 		"| | | | |  __/ | | | | | (_) | \n" +
// 	// 		"|_| |_|  \\___| |_| |_|  \\___/  \n" +
// 	// 		"                               \n" +
// 	// 		"                               \n" +
// 	// 		" _______   _                           \n" +
// 	// 		"|__   __| | |                          \n" +
// 	// 		"   | |    | |__     ___   _ __    ___  \n" +
// 	// 		"   | |    |  _ \\   / _ \\ | '__|  / _ \\ \n" +
// 	// 		"   | |    | | | | |  __/ | |    |  __/ \n" +
// 	// 		"   |_|    |_| |_|  \\___| |_|     \\___| \n" +
// 	// 		"                                       \n" +
// 	// 		"                                       \n"

// 	// if output != expected1 {
// 	// 	t.Errorf("expected this  %v  \nbut got this  %v , the length of result is \n %v \n while length of expected is \n %v", expected1, output, len(output), len(expected1))
// 	// } /*else if result2 != expected2 {
// 	// // 		t.Errorf("expected this  %v  \nbut got this  %v , the length of result is \n %v \n while length of expected is \n %v", expected2, result2, len(result2), len(expected2))
// 	// 	}*/
// }
