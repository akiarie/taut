package main

import (
	"fmt"
	"log"
	"regexp"
	"unicode"
)

type tokenclass string

const (
	tkPunctuation tokenclass = "punct"
	tkSpace                  = "[space]"
	tkId                     = "id"
	tkBits                   = "bits"
)

type token struct {
	class tokenclass
	value string
	pos   int // position of lexeme in input stream
}

func (tk token) String() string {
	return tk.value
}

func parsetoken(input string, pos int) (*token, int, error) {
	st := pos
	for _, c := range input[pos:] {
		if !unicode.IsSpace(c) {
			break
		}
		st++
	}
	if st > pos {
		if st >= len(input) {
			return &token{class: tkSpace, value: tkSpace, pos: pos}, len(input[pos:]), nil
		}
		// recurse & increment
		tk, shift, err := parsetoken(input, st)
		return tk, (st - pos) + shift, err
	}

	switch c := input[pos]; c {
	case '[', ']', ',':
		return &token{class: tkPunctuation, value: fmt.Sprintf("%c", c), pos: pos}, 1, nil
	}

	re := regexp.MustCompile(`^[A-Z]+`)
	if match := re.FindString(input[pos:]); match != "" {
		return &token{class: tkId, value: match, pos: pos}, len(match), nil
	}

	re = regexp.MustCompile(`^[01]+`)
	if match := re.FindString(input[pos:]); match != "" {
		return &token{class: tkId, value: match, pos: pos}, len(match), nil
	}
	return nil, -1, fmt.Errorf("Unknown characters %q", input[pos:])
}

func tokenize(input string) ([]token, error) {
	tokens := []token{}
	for pos := 0; pos < len(input); {
		tk, shift, err := parsetoken(input, pos)
		if err != nil {
			return nil, err
		}
		if tk.class != tkSpace {
			tokens = append(tokens, *tk)
		}
		pos += shift
	}
	return tokens, nil
}

func main() {
	raw := "ID[01], NOT[10], OR[0111], AND[0001], XOR[0110]"
	tokens, err := tokenize(raw)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tokens)
}
