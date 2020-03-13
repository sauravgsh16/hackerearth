package main

import (
	"fmt"
)

func main() {
	hm := NewHashMap()

	hm.put(1, 100)
	hm.put(2, 200)

	fmt.Printf("%d\n", hm.get(1))
	fmt.Printf("%d\n", hm.get(2))
	hm.remove(2)
	fmt.Printf("%d\n", hm.get(2))
}
