package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode"
)

type stream []token

type parser struct {
	pos   int
	raw   string
	input stream
	table table
}

func (p *parser) error(msg string) {
	panic(fmt.Sprintf("Error: %s at %d %q in %q", msg, p.input[p.pos].pos, p.input[p.pos].value, p.raw))
}

func (p *parser) parsetable() {
	p.table = table{}
	for p.pos < len(p.input) {
		var op operator
		if tk := p.input[p.pos]; tk.class == tkSpace && p.pos > 0 {
			p.pos++
		}
		if tk := p.input[p.pos]; tk.class != tkId {
			p.error("operator must start with identifier")
		} else {
			op.name = tk.value
			p.pos++
		}
		p.punct('[')
		if tk := p.input[p.pos]; tk.class == tkSpace {
			p.pos++
		}
		if tk := p.input[p.pos]; tk.class != tkBits {
			p.error("operator must contain bits")
		} else {
			op.bits = tk.value
			p.pos++
		}
		if tk := p.input[p.pos]; tk.class == tkSpace {
			p.pos++
		}
		p.punct(']')
		p.table = append(p.table, op)
		if p.pos+3 < len(p.input) {
			for p.pos < len(p.input) {
				if tk := p.input[p.pos]; tk.class == tkSpace {
					p.pos++
					continue
				} else if tk.value == "," {
					p.punct(',')
				}
				break
			}
		}
	}
}

func (p *parser) punct(c byte) {
	tk := p.input[p.pos]
	if tk.class != tkPunctuation || tk.value[0] != c {
		p.error(fmt.Sprintf("expected %q got %q", c, tk.value))
	}
	p.pos++
}

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
		return &token{class: tkSpace, value: tkSpace, pos: pos}, st - pos, nil
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
		return &token{class: tkBits, value: match, pos: pos}, len(match), nil
	}
	return nil, -1, fmt.Errorf("Unknown characters %q", input[pos:])
}

func tokenize(untrim string) ([]token, error) {
	input := strings.TrimSpace(untrim)
	tokens := []token{}
	for pos := 0; pos < len(input); {
		tk, shift, err := parsetoken(input, pos)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, *tk)
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
	p := &parser{input: tokens, raw: raw}
	p.parsetable()
	fmt.Println(p.table)
}
