package main

import (
	"payyans"
	"syscall/js"
)

func AsciiToUnicodeByMapString(this js.Value, args []js.Value) interface{} {
	output, err := payyans.AsciiToUnicodeByMapString(args[0].String(), args[1].String(), args[2].String())

	if err != nil {
		return js.ValueOf(err.Error())
	}

	return js.ValueOf(output)
}

func main() {
	js.Global().Set("AsciiToUnicodeByMapString", js.FuncOf(AsciiToUnicodeByMapString))
	select {} // Need to keep main running for wasm
}
