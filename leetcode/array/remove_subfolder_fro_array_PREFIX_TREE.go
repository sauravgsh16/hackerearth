package main

import (
	"fmt"
	"strings"
)

type trieNode struct {
	childeren map[string]*trieNode
	isEnd     bool
}

type trie struct {
	root *trieNode
}

func (t *trie) insert(words []string) {
	node := t.root

	n := len(words)
	for i := 0; i < n; i++ {
		char := string(words[i])
		//  if not present in node's children
		if _, ok := node.childeren[char]; !ok {
			node.childeren[char] = &trieNode{
				childeren: make(map[string]*trieNode),
			}
		}
		node = node.childeren[char]
	}
	node.isEnd = true
}

func (t *trie) find() []string {
	ans := []string{}
	dfsTrie(ans, t.root)
	return ans
}

func dfsTrie(ans []string, node *trieNode) {
	if node.isEnd {
		ans = append(ans, "/"+strings.Join(ans, "/"))
		return
	}

	for char := range node.childeren {
		ans = append(ans, char)
		dfsTrie(ans, node.childeren[char])
	}
}

func removeSubFolders(folders []string) []string {
	t := &trie{
		root: &trieNode{
			childeren: make(map[string]*trieNode),
		},
	}
	for _, f := range folders {
		flist := strings.Split(f, "/")[1:]
		t.insert(flist)
	}

	fmt.Printf("%+v\n", t.root)

	return t.find()
}
