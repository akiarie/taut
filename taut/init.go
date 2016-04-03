package main

import (
	"amisi/taut/input"
	"fmt"
)

func main() {
	table, err := input.Table("A[0011], B[0101] -> OR[0111], AND[0001]")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(table)
}
