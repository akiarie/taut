package input

import (
	"amisi/taut/truth"
	"reflect"
	"testing"
)

func checkRange(t *testing.T, ins []string, msg string, format ...interface{}) {
	for _, in := range ins {
		_, err := Table(in)
		if err == nil {
			t.Errorf(msg, in)
		}
	}
}

func TestTable(t *testing.T) {
	// basic should-fails
	formats := []string{
		"", "A[] -> B[00]", "A <> B", "A[000], B[111]",
		"A[0011], B[0101] - OR[0111], AND[0001]",
	}
	checkRange(t, formats, "Table() allowed invalid string to pass: '%s'")
	dups := []string{
		"A[0]->A[1]", "A[110],B[011]->C[001],B[011]", "AMISI[001],IS[010] -> NOT[011], AMISI[101]",
	}
	checkRange(t, dups, "Table() allowed duplicate string to pass: '%s'")

	// check basic functions
	ins := []truth.Tcol{
		truth.Tcol{Name: "A", Values: []bool{false, false, true, true}},
		truth.Tcol{Name: "B", Values: []bool{false, true, false, true}},
	}
	defs := map[string]truth.Table{
		"AND": truth.Table{ins, []truth.Tcol{truth.Tcol{Name: "Y", Values: []bool{false, false, false, true}}}},
		"OR":  truth.Table{ins, []truth.Tcol{truth.Tcol{Name: "Y", Values: []bool{false, true, true, true}}}},
		"XOR": truth.Table{ins, []truth.Tcol{truth.Tcol{Name: "Y", Values: []bool{false, true, true, false}}}},
		"NOR": truth.Table{ins, []truth.Tcol{truth.Tcol{Name: "Y", Values: []bool{true, false, false, false}}}},
	}
	and, _ := Table("A[0011],B[0101] -> Y[0001]")
	or, _ := Table("A[0011],B[0101] -> Y[0111]")
	xor, _ := Table("A[0011],B[0101] -> Y[0110]")
	nor, _ := Table("A[0011],B[0101] -> Y[1000]")
	tables := map[string]truth.Table{
		"AND": and, "OR": or, "XOR": xor, "NOR": nor,
	}
	for key, _ := range tables {
		if !reflect.DeepEqual(defs[key], tables[key]) {
			t.Errorf("Table() unable to parse basic table '%s': \n%v\n%v.", key, defs[key], tables[key])
		}
	}
}
