package payyans

import (
	"log"
	"strings"
)

var preBase = []string{"േ", "െ", "ൈ", "ോ", "ൊ", "ൌ", "്ര"}
var postBase = []string{"്യ", "്വ"}

func AsciiToUnicode(input string, rulesMap map[string]string) string {
	const symbols = " 	,;.:?01234'\"56789\n"
	const mixCharactors = "CDH"

	var output = ""

	runes := []rune(input)

	i := 0
	for i < len(runes) {
		currentCharacter := string(runes[i])
		p := rulesMap[currentCharacter]

		nextCharacter := ""
		if i+1 < len(runes) {
			nextCharacter = string(runes[i+1])
		}

		nextNextCharacter := ""
		if i+2 < len(runes) {
			nextNextCharacter = string(runes[i+2])
		}

		if strings.Contains(symbols, currentCharacter) {
			output += currentCharacter
		} else if p == "" {
			output += currentCharacter
		} else if strings.Contains(mixCharactors, currentCharacter) {
			// to deal with "ഈ,ഊ,ഔ,ഓ"

			nextCharacter := string(runes[i+1])
			if nextCharacter == "u" {
				if currentCharacter == "C" {
					output += rulesMap["Cu"] // ഈ
				} else if currentCharacter == "D" {
					output += rulesMap["Du"] // ഊ
				} else if currentCharacter == "H" {
					output += rulesMap["Hu"] // ഔ
				}
				i++
			} else if nextCharacter == "m" {
				if currentCharacter == "H" {
					output += rulesMap["Hm"] // ഓ
				} else if currentCharacter == "t" {
					output += rulesMap["tm"] // ോ
				}
				i++
			} else {
				output += rulesMap[currentCharacter]
			}

		} else if currentCharacter == "s" && nextCharacter == "F" {
			output += rulesMap["sF"] // ഐ
			i++
		} else if currentCharacter == "s" && nextCharacter == "s" {
			output += rulesMap[nextNextCharacter+rulesMap["ss"]] // ൈ
			i = i + 2
		} else if Contains[string](preBase, p) {
			preBaseCharacter := p
			mainCharactor := rulesMap[nextCharacter]
			qCharacter := rulesMap[nextNextCharacter]

			if Contains[string](postBase, qCharacter) {
				output += mainCharactor + qCharacter + preBaseCharacter
				i = i + 2
			} else {
				if strings.Contains(symbols, nextNextCharacter) {
					output += mainCharactor + preBaseCharacter + nextNextCharacter
					i = i + 2
				} else {
					if nextCharacter == "{" {
						output += qCharacter + rulesMap["{"] + preBaseCharacter
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
		i++
	}

	return output
}

func AsciiToUnicodeByMapFile(input string, mapFilePath string) string {
	rulesMap, err := ReadAndCleanFile(mapFilePath)

	if err != nil {
		log.Println("error: rule file not found")
		log.Fatal(err.Error())
	}

	return AsciiToUnicode(input, rulesMap)
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
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
