package main

import (
	"amisi/taut/input"
	"fmt"
	"log"
)

func main() {
	table, err := input.Table("A, B -> Y[0001]")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(table)
}
