package payyans

import (
	"fmt"
	"strings"
)

var preBase = []string{"േ", "െ", "ൈ", "ോ", "ൊ", "ൌ", "്ര"}
var postBase = []string{"്യ", "്വ"}

func AsciiToUnicode(input string, font string) string {
	mapString, err := GetRules(font)

	if err != nil {
		fmt.Println("error: rule file not found")
	}

	const symbols = " 	,;.:?01234\\'\"56789\n"
	const mixCharactors = "CDH"

	var output = ""

	for i := 0; i < len(input); i++ {
		stripedCharacter := input[i : i+1]
		p := mapString[stripedCharacter]
		sizeJ := i + 1
		sizeK := i + 2

		if strings.Contains(symbols, stripedCharacter) {
			output += stripedCharacter
		} else if p == "" {
			output += stripedCharacter
		} else if strings.Contains(mixCharactors, stripedCharacter) {
			nextCharactor := input[sizeJ : sizeJ+1]
			if nextCharactor == "u" {
				if stripedCharacter == "C" {
					output += mapString["Cu"] // ഈ
				} else if stripedCharacter == "D" {
					output += mapString["Du"] // ഊ
				} else if stripedCharacter == "H" {
					output += mapString["Hu"] // ഔ
				}
				i++
			} else if nextCharactor == "m" {
				if stripedCharacter == "H" {
					output += mapString["Hm"] // ഓ
				} else if stripedCharacter == "t" {
					output += mapString["tm"] // ോ
				}
				i++
			} else {
				output += mapString[stripedCharacter]
			}

		} else if stripedCharacter == "s" && input[sizeJ:sizeJ+1] == "F" {
			output += mapString["sF"] // ഐ
			i++
		} else if stripedCharacter == "s" && input[sizeJ:sizeJ+1] == "s" {
			output += mapString[input[sizeK:sizeK+1]+mapString["ss"]] // ൈ
			i = i + 2
		} else if Contains[string](preBase, p) {
			preBaseCharacter := p
			mainCharactor := mapString[input[sizeJ:sizeJ+1]]
			qCharacter := mapString[input[sizeK:sizeK+1]]

			if Contains[string](postBase, qCharacter) {
				output += mainCharactor + qCharacter + preBaseCharacter
				i = i + 2
			} else {
				if strings.Contains(symbols, input[sizeK:sizeK+1]) {
					output += mainCharactor + preBaseCharacter + input[sizeK:sizeK+1]
					i = i + 2
				} else {
					if input[sizeJ:sizeJ+1] == "{" {
						output += qCharacter + mapString["{"] + preBaseCharacter
						i = i + 2
					} else {
						output += mainCharactor + preBaseCharacter
						i++
					}
				}
			}

		} else {
			output += p
		}

	}

	return output

}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func GetRules(font string) (map[string]string, error) {
	return ReadAndCleanFile(font)
}

// func Unicode2ASCII(unicodeText string, font string) string {
// 	// Normalize the Unicode text
// 		normalizedText := normalizer.Normalize(unicodeText)

// 	// Create a map of Unicode characters to their ASCII equivalents
// 	rulesReverse := getRules(font)
// 	rulesDict := map[string]string{}
// 	for k, v := range rulesReverse {
// 		rulesDict[v] = k
// 	}

// 	// Initialize the output string
// 	asciiText := ""

// 	// Iterate over the Unicode text
// 	for index := 0; index < len(normalizedText); index++ {
// 	  // Check for a three-character sequence
// 	  for charNo := 3; charNo >= 1; charNo-- {
// 		letter := normalizedText[index : index+charNo];
// 		if letter in rulesDict {
// 		  asciiLetter := rulesDict[letter]
// 		  asciiText += asciiLetter
// 		  index += charNo
// 		  break
// 		}
// 	  }

// 	  // If we didn't find a match, just add the character to the output string
// 	  if index < len(normalizedText) {
// 		asciiText += normalizedText[index : index+1]
// 	  }
// 	}

// 	return asciiText
//   }
