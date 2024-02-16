package gopayyans



func ASCIIToUnicode(ASCIIText, font string) {
	fmt.Println(ASCIIText, font)
}

func GetRules(font string) (map[string]string, error) {
	return ReadAndCleanFile(font)
}



func Unicode2ASCII(unicodeText string, font string) string {
	// Normalize the Unicode text
	normalizedText := normalizer.Normalize(unicodeText)

	// Create a map of Unicode characters to their ASCII equivalents
	rulesReverse := getRules(font)
	rulesDict := map[string]string{}
	for k, v := range rulesReverse {
		rulesDict[v] = k
	}

	// Initialize the output string
	asciiText := ""
  
	// Iterate over the Unicode text
	for index := 0; index < len(normalizedText); index++ {
	  // Check for a three-character sequence
	  for charNo := 3; charNo >= 1; charNo-- {
		letter := normalizedText[index : index+charNo];
		if letter in rulesDict {
		  asciiLetter := rulesDict[letter]
		  asciiText += asciiLetter
		  index += charNo
		  break
		}
	  }

	  // If we didn't find a match, just add the character to the output string
	  if index < len(normalizedText) {
		asciiText += normalizedText[index : index+1]
	  }
	}
  
	return asciiText
  }
  