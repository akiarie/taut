package input

import (
	"amisi/taut/truth"
	"fmt"
	"log"
	"math"
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

type matchMode string

const (
	modeImplicit matchMode = `([A-Z]+)`
	modeExplicit matchMode = `([A-Z]+)\[([01]+)\]`
)

// namespace returns a function which returns true only for the first call of a particular string
func namespace() func(string) bool {
	var names []string
	return func(s string) bool {
		for _, name := range names {
			if s == name {
				return false
			}
		}
		// unique, add to list
		names = append(names, s)
		return true
	}
}

func parse(in string, mode matchMode, unique func(string) bool) (cols []truth.Tcol, err error) {
	reVal := regexp.MustCompile(string(mode))
	matches := reVal.FindAllStringSubmatch(in, -1)
	// holds 2^len(matches) which is the no. of rows
	nrows := int(math.Pow(2, float64(len(matches))))
	for i, match := range matches {
		var values []bool
		switch mode {
		case modeImplicit:
			// holds 2^i which is the frequency of flips
			freq := int(math.Pow(2, float64(i)))
			for j := 0; j < nrows; j++ {
				v := (j/freq)%2 == 1
				values = append(values, v)
			}
		case modeExplicit:
			for _, v := range match[2] {
				values = append(values, v == '1')
			}
		default:
			return []truth.Tcol{}, fmt.Errorf("Unknown matchmode '%s'", mode)
		}

		name := match[1]
		if !unique(name) {
			return []truth.Tcol{}, fmt.Errorf("Duplicate name '%s' for: %v", name, match)
		}

		cols = append(cols, truth.Tcol{name, values})
	}
	return
}

// Table takes a raw string and returns a truth.Table representing it
func Table(in string) (truth.Table, error) {
	// remove spaces
	nospace := stripSpaces(in)

	// ensure string matches format
	reFormat := regexp.MustCompile(`^[A-Z]+(,[A-Z]+)*->[A-Z]+\[[01]+\](,[A-Z]+\[[01]+\])*`)
	if !reFormat.MatchString(nospace) {
		return truth.Table{}, fmt.Errorf("Invalid input string: '%s'", in)
	}

	// break string into input/output components
	part := strings.SplitN(nospace, "->", 2)
	rawIn := part[0]
	rawOut := part[1]

	// one unique namespace for both variables
	variables := namespace()

	input, err := parse(rawIn, modeImplicit, variables)
	if err != nil {
		return truth.Table{}, err
	}
	output, err := parse(rawOut, modeExplicit, variables)
	if err != nil {
		return truth.Table{}, err
	}

	// confirm lengths match
	l := len(input[0].Values)
	for _, col := range append(input[1:], output...) {
		if len(col.Values) != l {
			return truth.Table{}, fmt.Errorf("Invalid input length: len(%v) != len(%v).", col, input[0])
		}
	}

	return truth.Table{input, output}, nil
}
