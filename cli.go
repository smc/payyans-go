package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"

	"payyans"
)

//go:embed unicode-conversion-maps/maps/*.map
var fontMapsFs embed.FS

//go:embed normalizer/libindic/normalizer/*.rules
var normalizerRulesFs embed.FS

func main() {
	asciiToUnicode := flag.Bool("ascii-to-unicode", true, "ASCII to Unicode conversion")
	unicodeToAscii := flag.Bool("unicode-to-ascii", false, "Unicode to ASCII conversion")

	font := flag.String("font", "", "Font name")
	fonts := flag.Bool("fonts", false, "List all available fonts")

	fontMapFilePath := flag.String("font-map-file", "", "Path to font map file")
	normalizerRulesFilePath := flag.String("normalizer-rules-file", "", "Path to normalizer rules file")

	flag.Parse()

	if *fonts {
		files, err := fs.ReadDir(fontMapsFs, "unicode-conversion-maps/maps")
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			fmt.Println(strings.Replace(file.Name(), ".map", "", 1))
		}

		return
	}

	if *font == "" && *fontMapFilePath == "" {
		fmt.Println("Please specify font with -font or path to font map file with -font-map-file.\n\nUse -fonts to list all available fonts.\n\nUse --help for all available commands.")
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

		var fontMapFileBytes []byte

		if *fontMapFilePath != "" {
			fontMapFileBytes, err = os.ReadFile(*fontMapFilePath)
			if err != nil {
				log.Print("Unable to read font map file from filesystem")
				log.Fatal(err.Error())
			}
		} else {
			fontMapFileBytes, err = fs.ReadFile(fontMapsFs, path.Join("unicode-conversion-maps", "maps", *font+".map"))
			if err != nil {
				log.Print("Unable to find font. See list of available fonts with -font")
				log.Fatal(err.Error())
			}
		}

		var normalizerMapFileBytes []byte

		if *normalizerRulesFilePath != "" {
			normalizerMapFileBytes, err = os.ReadFile(*normalizerRulesFilePath)
			if err != nil {
				log.Print("Unable to read normalizer map file")
				log.Fatal(err.Error())
			}
		} else {
			normalizerMapFileBytes, err = fs.ReadFile(normalizerRulesFs, path.Join("normalizer", "libindic", "normalizer", "normalizer_ml.rules"))
			if err != nil {
				log.Print("Unable to find normalizer map file from binary")
				log.Fatal(err.Error())
			}
		}

		outputString, err := payyans.AsciiToUnicodeByMapString(string(bytes), string(fontMapFileBytes), string(normalizerMapFileBytes))

		if err != nil {
			log.Fatal(err)
		}

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
