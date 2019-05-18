package main

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	hexColour   = regexp.MustCompile(fmt.Sprintf("#(?:%s{6}|%[1]s{3})", hexDigit))
	hexDisabled = false
)

func parseHex(line string) colours {
	var clrs colours
	if hexDisabled {
		return clrs
	}

	matches := hexColour.FindAllStringIndex(line, -1)
	for _, match := range matches {
		colour := &Colour{
			ColStart: match[0] + 1,
			ColEnd:   match[1],
			Hex:      strings.ToLower(toFullHex(line[match[0]:match[1]])),
			Line:     line,
		}
		clrs = append(clrs, colour)
	}
	return clrs
}
