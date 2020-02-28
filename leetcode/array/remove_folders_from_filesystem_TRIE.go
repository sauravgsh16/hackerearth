package main

import (
	"strings"
)

type trieNode struct {
	children map[string]*trieNode
	isEnd    bool
}

type trie struct {
	root *trieNode
}

func (t *trie) insert(word []string) {
	node := t.root

	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			node.children[char] = &trieNode{
				children: make(map[string]*trieNode),
			}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func (t *trie) addSelective(word []string) bool {
	node := t.root
	n := len(word)

	for i, char := range word {
		if _, ok := node.children[char]; !ok {
			return false
		}
		node = node.children[char]
		if node.isEnd && i < n-1 {
			return false
		}
	}
	if !node.isEnd {
		return false
	}
	return true
}

func removeSubfolders(folders []string) []string {
	t := &trie{
		root: &trieNode{
			children: make(map[string]*trieNode),
		},
	}

	for _, f := range folders {
		word := strings.Split(f, "/")[1:]
		if len(word) == 0 {
			continue
		}
		t.insert(word)
	}

	ans := []string{}
	for _, f := range folders {
		word := strings.Split(f, "/")[1:]
		if len(word) == 0 {
			continue
		}
		if !t.addSelective(word) {
			continue
		}
		ans = append(ans, f)
	}
	return ans
}
