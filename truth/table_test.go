package truth

import "testing"

func TestTable(t *testing.T) {
	ins := []Tcol{
		Tcol{Name: "A", Values: []bool{false, true, false, true}},
		Tcol{Name: "B", Values: []bool{false, false, true, true}},
	}
	defs := map[string]Table{
		"AND": Table{ins, []Tcol{Tcol{Name: "Y", Values: []bool{false, false, false, true}}}},
		"OR":  Table{ins, []Tcol{Tcol{Name: "Y", Values: []bool{false, true, true, true}}}},
		"XOR": Table{ins, []Tcol{Tcol{Name: "Y", Values: []bool{false, true, true, false}}}},
		"NOR": Table{ins, []Tcol{Tcol{Name: "Y", Values: []bool{true, false, false, false}}}},
	}
	for def, table := range defs {
		if !table.Valid() {
			t.Errorf("truth.Table definition unable to parse basic '%s' statement.", def)
		}
	}
}
