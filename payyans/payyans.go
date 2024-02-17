package payyans

import (
	"log"
)

var preBase = []string{"േ", "െ", "ൈ", "ോ", "ൊ", "ൌ", "്ര"}
var postBase = []string{"്യ", "്വ"}

func AsciiToUnicode(input string, rulesMap map[string]string) string {
	reverseRulesMap := reverseMap(rulesMap)

	preBaseAsciiLetters := make(map[string]bool)
	postBaseAsciiLetters := make(map[string]bool)

	// ആദ്യത്തെ ഓട്ടം: മുമ്പേ ഗമിക്കും പ്രീബേസിനെ പിടിച്ച് തോളില്‍ കയറ്റുക
	for _, char := range preBase {
		preBaseAsciiLetters[reverseRulesMap[char]] = true
	}

	for _, char := range postBase {
		postBaseAsciiLetters[reverseRulesMap[char]] = true
	}

	runes := []rune(input)

	transposedText := ""
	prebase := ""
	i := 0
	for i < len(runes) {
		currentChar := string(runes[i])

		if keyExists(preBaseAsciiLetters, currentChar) {
			prebase = currentChar + prebase
		} else if keyExists(postBaseAsciiLetters, currentChar) {
			transposedText += currentChar + prebase
			prebase = ""
		} else {
			transposedText += currentChar
			if i+1 < len(runes) {
				nextCharacter := string(runes[i+1])
				if !keyExists(postBaseAsciiLetters, nextCharacter) {
					transposedText += prebase
					prebase = ""
				}
			}
		}

		if prebase != "" {
			transposedText += prebase
		}

		// രണ്ടാമത്തെ ഓട്ടം: പച്ച മലയാളം
		unicodeText := ""
		i := 0
		for i < len(runes) {
			currentChar := string(runes[i])
			if mappedChar, ok := rulesMap[currentChar]; ok {
				unicodeText += mappedChar
			} else {
				unicodeText += currentChar
			}
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

func contains[T comparable](s []T, e T) bool {
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
