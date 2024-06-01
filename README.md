

# ascii-arts
## Introduction

This Go program converts input strings into ASCII art.

## Usage

1. **Run the program:** Run the Go source file directly with your desired input string.

Replace `"Your input string"` with the string you want to convert to ASCII art.

2. **Example:**

go run . "Hello" | cat -e
  
This will generate ASCII art for the word "Hello".

  ```

 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $

 ```
## Features

- Converts input strings into ASCII art.
- Handles input strings with newline characters (`\n`).
- Validates input characters to ensure they are within the ASCII range.
- Limits the length of input strings for better printing.

## Limitations

- The program currently limits input strings to a maximum length of 25 characters. Longer strings may result in distorted output.
- The program assumes ASCII characters within the range of printable characters (32 to 126). Characters outside this range may not be rendered correctly.


## Contributors

- [viarony]()
- [jerootieno]()



 