package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var pair = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
}

func isBalanced(txt string) bool {
	s := []string{}

	for _, t := range txt {

		switch string(t) {
		case "(":
			s = append(s, string(t))
		case "{":
			s = append(s, string(t))
		case "[":
			s = append(s, string(t))
		default:
			if len(s) <= 0 {
				return false
			}
			b := s[len(s)-1]

			switch string(t) {
			case ")":
				if pair[b] != ")" {
					return false
				}
			case "}":
				if pair[b] != "}" {
					return false
				}
			case "]":
				if pair[b] != "]" {
					return false
				}
			}

			s = s[:len(s)-1]
		}
	}
	if len(s) > 0 {
		return false
	}
	return true
}

func main() {
	var loop int
	fmt.Scanf("%d", &loop)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 1; i++ {
		txt, _ := reader.ReadString('\n')
		txt = strings.TrimSpace(txt)
		if isBalanced(txt) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
