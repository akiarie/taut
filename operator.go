package main

import "fmt"

type table []operator

func (t table) String() string {
	var s string
	if len(t) == 0 {
		return "Ø"
	}
	var B int
	for i := 0; ; i++ {
		if pw := 1 << i; pw > len(t[0].bits) {
			panic(fmt.Sprintf("invalid bits length %d", B))
		} else if pw == len(t[0].bits) {
			B = i
			break
		}
	}
	vars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"[:B]
	// print header
	pad := map[string]int{}
	for i := 0; i < len(vars); i++ {
		fmt.Printf("%c ", vars[i])
	}
	for i, op := range t {
		space := " "
		if i+1 == len(t) {
			space = ""
		}
		fmt.Printf("%s%s", op.name, space)
		pad[op.name] = len(op.name)
	}
	fmt.Println()
	// values
	for i := 0; i < 1<<B; i++ { // rows
		for j := 0; j < len(vars); j++ { // cols
			k := len(vars) - 1 - j // the current power of 2
			mask := 1 << k
			fmt.Printf("%d ", (i&mask)>>k)
		}
		for z, op := range t {
			space := " "
			if z+1 == len(t) {
				space = ""
			}
			fmt.Printf("%*c%s", pad[op.name], op.bits[i], space)
		}
		fmt.Println()
	}
	return s
}

type operator struct {
	bits, name string
}
