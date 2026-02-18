package main

import "fmt"

type TrieNode struct {
	IsEnd    bool
	Children [26]*TrieNode
}

type Trie struct {
	Root *TrieNode
}

func newTrie() *Trie {
	return &Trie{
		Root: &TrieNode{},
	}
}

// inserting new word in trie
func (t *Trie) insert(word string) {
	current := t.Root // i am starting from the root of the Trie

	for _, ch := range word {

		idx := ch - 'a'
		if current.Children[idx] == nil {
			current.Children[idx] = &TrieNode{}
		}
		current = current.Children[idx]

	}
	current.IsEnd = true

}

// now searching word from the trie

func (t *Trie) search(word string) bool {
	current := t.Root

	for _, ch := range word {
		idx := ch - 'a'

		if current.Children[idx] == nil {
			return false
		}
		current = current.Children[idx]
	}
	if current.IsEnd == true {
		return true
	}
	return false
}

func main() {
	trie := newTrie()
	//fmt.Println(trie)
	//fmt.Println(trie.Root)

	words := []string{"aqua", "jack", "card"}
	for _, word := range words {
		trie.insert(word)
	}

	wordFind := "hello"
	if trie.search(wordFind) {
		fmt.Println("Found this word")
	} else {
		fmt.Println("not found this ")
	}
}
