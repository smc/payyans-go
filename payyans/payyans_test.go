package payyans

import (
	"log"
	"os"
	"reflect"
	"runtime/debug"
	"strings"
	"testing"
)

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func readFile(path string) string {
	contents, err := os.ReadFile(path)
	checkError(err)
	return string(contents)
}

// AssertEqual checks if values are equal
// Thanks https://gist.github.com/samalba/6059502#gistcomment-2710184
func assertEqual(t *testing.T, value interface{}, expected interface{}) {
	if value == expected {
		return
	}
	debug.PrintStack()
	t.Errorf("Received %v (type %v), expected %v (type %v)", value, reflect.TypeOf(value), expected, reflect.TypeOf(expected))
}

func TestAsciiToUnicodeConversion(t *testing.T) {
	bytes, err := os.ReadFile("testdata/ML-TTKarthika.txt")
	checkError(err)

	lines := strings.Split(string(bytes), "\n\n")

	for _, line := range lines {
		inputAndExpected := strings.Split(line, "\n")
		output, err := AsciiToUnicodeByMapString(
			inputAndExpected[0],
			readFile("../unicode-conversion-maps/maps/ML-TTKarthika.map"),
			readFile("../normalizer/libindic/normalizer/normalizer_ml.rules"),
		)

		checkError(err)

		assertEqual(t, output, inputAndExpected[1])
	}
}

func TestMain(m *testing.M) {
	m.Run()
}
