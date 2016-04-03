package truth

import "fmt"

type Tcol struct {
	Values []bool
	Name   string
}

type Table struct {
	Input  []Tcol
	Output []Tcol
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
