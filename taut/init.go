package main

import (
	"amisi/taut/input"
	"fmt"
	"log"
)

func main() {
	table, err := input.Table("A, B -> OR[0111]")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(table)
}
