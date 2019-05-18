package main

import (
	"fmt"
	"os"
)

// FMT is the colour format which will be printed to stdout
type FMT int

const (
	// ExtendedFmt will print the format "filename:lnum:colstart-colend:hex:line"
	ExtendedFmt FMT = iota
	// ShortFmt will print the format "filename:lnum:colstart-colend:hex"
	ShortFmt

	extendedFmt = "%s:%d:%d-%d:%s:%s\n"
	shortFmt    = "%s:%d:%d-%d:%s\n"
)

var outputFmt = ExtendedFmt

// SetOutputFmt sets the format to print the colour
func SetOutputFmt(fmt FMT) {
	outputFmt = fmt
}

// PrintColours prints clrs to out based on the formatting specificed by
// SetOutputFmt.
func PrintColours(clrs colours, out *os.File) {
	for _, colour := range clrs {
		if outputFmt == ExtendedFmt {
			fmt.Fprintf(out, extendedFmt, colour.Tag, colour.Lnum, colour.ColStart, colour.ColEnd, colour.Hex, colour.Line)
		} else {
			fmt.Fprintf(out, shortFmt, colour.Tag, colour.Lnum, colour.ColStart, colour.ColEnd, colour.Hex)
		}
	}
}
