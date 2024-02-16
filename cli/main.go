package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"payyans"
)

func main() {
	asciiToUnicode := flag.Bool("ascii-to-unicode", true, "ASCII to Unicode conversion. Selected by default")
	unicodeToAscii := flag.Bool("unicode-to-ascii", false, "Unicode to ASCII conversion")

	fontMapFilePath := flag.String("map", "", "Path to map file")

	flag.Parse()

	if *fontMapFilePath == "" {
		fmt.Println("Please specify the path to font map file with -map.\n\nUse --help for all available commands.")
		return
	}

	args := flag.Args()

	if len(args) < 1 {
		log.Fatal("Please provide input filename")
	}

	inputFilename := args[0]
	outputFilename := ""
	if len(args) > 1 {
		outputFilename = args[1]
	}

	if *unicodeToAscii {
		*asciiToUnicode = false
	}

	if *asciiToUnicode {
		bytes, err := os.ReadFile(inputFilename)
		if err != nil {
			log.Print("Unable to read file")
			log.Fatal(err.Error())
		}

		outputString := payyans.AsciiToUnicodeByMapFile(string(bytes), *fontMapFilePath)

		if outputFilename != "" {
			err = os.WriteFile(outputFilename, []byte(outputString), 0644)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println(outputString)
		}
	} else {

	}
}
