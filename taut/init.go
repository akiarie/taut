package main

import (
	"amisi/taut/input"
	"fmt"
	"log"
)

func main() {
	table, err := input.Table("A -> NOT[10]")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(table)
}
