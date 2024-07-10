package main

import "fmt"

type TrieNode struct {
	childNode [26]*TrieNode
	wordEnd   bool
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) insert(data string) {
	node := t.root
	for _, char := range data {
		if node.childNode[char-'a'] == nil {
			node.childNode[char-'a'] = &TrieNode{}
		}
		node = node.childNode[char-'a']
	}
	node.wordEnd = true
}

func getChildNodes(n [26]*TrieNode) int32 {
	var i int32 = 0
	for _, node := range n {
		if node != nil {
			i += getChildNodes(node.childNode)
			if node.wordEnd {
				return 1
			}
		}
	}
	return i
}

func (t *Trie) search(data string) int32 {
	node := t.root
	for _, char := range data {
		if node.childNode[char-'a'] == nil {
			return 0
		}
		node = node.childNode[char-'a']
	}
	return getChildNodes(node.childNode)
}

func contacts(queries [][]string) []int32 {
	trie := Trie{root: &TrieNode{}}
	for _, v := range queries {
		str := v[0]
		if str[:1] == "a" {
			trie.insert(str[4:])
		} else if str[:1] == "f" {
			fmt.Println(trie.search(str[5:]))
		}
	}
	return []int32{}
}

func main() {
	contacts([][]string{
		{"add hack"},
		{"add hackerrank"},
		{"find hac"},
		{"find hak"},
	})
}
