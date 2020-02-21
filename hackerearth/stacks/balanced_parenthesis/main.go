package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	txt, _ := reader.ReadString('\n')
	txt = strings.TrimSpace(txt)

	minAddition(txt)
}

func minAddition(p string) {
	n := len(p)
	stack := []string{}

	for i := 0; i < n; i++ {

		if string(p[i]) == "(" {
			stack = append(stack, string(p[i]))
		} else {
			if len(stack) > 0 {
				if stack[len(stack)-1] == "(" {
					stack = stack[:len(stack)-1]
					continue
				}
			}
			stack = append(stack, string(p[i]))

		}
	}

	fmt.Printf("%d\n", len(stack))
}
