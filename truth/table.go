package truth

type Tcol struct {
	Values []bool
	Name   string
}

type Table struct {
	Input  []Tcol
	Output []Tcol
}
