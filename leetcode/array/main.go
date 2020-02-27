package main

import "fmt"

func main() {
	a := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
	fmt.Printf("%v\n", removeSubFolders(a))
}
