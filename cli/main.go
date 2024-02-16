package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/smc/payyans-go/payyans"
)

func main() {
	asciiToUnicode := flag.Bool("ascii-to-unicode", true, "ASCII to Unicode conversion. Selected by default")
	unicodeToAscii := flag.Bool("unicode-to-ascii", false, "Unicode to ASCII conversion")

	fontName := flag.String("font", "", "Font name")

	flag.Parse()

	if *fontName == "" {
		fmt.Println("Specify a font name with -font.\n\nUse --help for all available commands.")
		return
	}

	args := flag.Args()

	inputFilename := args[0]
	outputFilename := args[1]

	if *unicodeToAscii {
		*asciiToUnicode = false
	}

	if *asciiToUnicode {
		bytes, err := os.ReadFile(inputFilename)
		if err != nil {
			log.Print("Unable to read file")
			log.Fatal(err.Error())
		}

		outputBytes := payyans.AsciiToUnicode(string(bytes))

		err = os.WriteFile(outputFilename, outputBytes, 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else {

	}
}
