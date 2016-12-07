package truth

import "fmt"

type Tcol struct {
	Name   string
	Values []bool
}

type Table struct {
	Input  []Tcol
	Output []Tcol
}

// Valid returns a boolean indicating whether or not the
// statement(s) represented by the able are well-defined.
func (t Table) Valid() bool {
	// at least one I/O
	if len(append(t.Input, t.Output...)) < 2 {
		return false
	}

	// confirm lengths identical
	l := len(t.Input[0].Values)
	for _, col := range append(t.Input[1:], t.Output...) {
		if len(col.Values) != l {
			return false
		}
	}

	// TODO: check for well-defined-ness of values
	return true
}

// Response gives the output state corresponding to the
// given input values.
func (t Table) Response(input []bool) ([]bool, error) {
	// validate table
	if !t.Valid() {
		return []bool{}, fmt.Errorf("Invalid table queried for response.")
	}
	return []bool{}, fmt.Errorf("Implement!")
}

func (t Table) String() (s string) {
	cols := append(t.Input, t.Output...)

	// find widest name
	var w int
	for _, col := range cols {
		test := len(col.Name)
		if test > w {
			w = test
		}
	}
	width := fmt.Sprintf("%d", w-1)

	for _, col := range cols {
		s += fmt.Sprintf("% "+width+"s\t", col.Name)
	}
	for i := 0; i < len(cols[0].Values); i++ {
		s += "\n"
		for _, col := range cols {
			var val int
			if col.Values[i] {
				val = 1
			}
			s += fmt.Sprintf("% "+width+"s\t", fmt.Sprintf("%d", val))
		}
	}
	return
}
