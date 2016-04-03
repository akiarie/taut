package input

import (
	"amisi/taut/truth"
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode"
)

// stripSpaces returns a copied version of str without
// whitespace as defined by unicode.IsSpace()
// credit: http://stackoverflow.com/a/32082217/3928922
func stripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// intArrayToBool converts an array of integers
// with values of 1/0 to an array of booleans
func intArrayToBool(ints []int) (bools []bool) {
	for _, i := range ints {
		if i < 0 || i > 1 {
			log.Fatalf("Error converting ints to bool with value: %d", i)
		}
		if i == 0 {
			bools = append(bools, false)
		} else {
			bools = append(bools, true)
		}
	}
	return
}

// Table takes a raw string and returns a truth.Table representing it
func Table(in string) (truth.Table, error) {
	// remove spaces
	nospace := stripSpaces(in)

	// check string for format
	reFormat := regexp.MustCompile(`^[A-Z]+\[[01]+\](,[A-Z]+\[[01]+\])*->[A-Z]+\[[01]+\](,[A-Z]+\[[01]+\])*`)
	if !reFormat.MatchString(nospace) {
		return truth.Table{}, fmt.Errorf("Invalid input strings: '%s'", in)
	}

	part := strings.SplitN(nospace, "->", 2)
	rawIn := part[0]
	rawOut := part[1]

	reVal := regexp.MustCompile(`([A-Z]+)\[([01]+)\]`)

	var err error
	var names []string
	unavailable := func(s string) bool {
		for _, test := range names {
			if s == test {
				return true
			}
		}
		// available, but remove for future
		names = append(names, s)
		return false
	}
	// builds truth cols based on each item matched by expression
	parse := func(vals string) (cols []truth.Tcol) {
		for _, match := range reVal.FindAllStringSubmatch(vals, -1) {
			var values []bool
			for _, v := range match[2] {
				values = append(values, v == '1')
			}
			name := match[1]
			if unavailable(name) {
				err = fmt.Errorf("Duplicate name '%s' for: %v", name, match)
			}
			cols = append(cols, truth.Tcol{Name: name, Values: values})
		}
		return
	}

	var input, output []truth.Tcol = parse(rawIn), parse(rawOut)
	if err != nil {
		return truth.Table{}, err
	}

	// confirm lengths match
	l := len(input[0].Values)
	for _, col := range append(input, output...) {
		if len(col.Values) != l {
			return truth.Table{}, fmt.Errorf("Invalid input length: len(%v) != len(%v).", col, input[0])
		}
	}

	return truth.Table{input, output}, nil
}
