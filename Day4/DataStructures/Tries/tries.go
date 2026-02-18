package Tries

import (
	"strings"
)

type TrieNode struct {
	IsEnd    bool
	Children [26]*TrieNode
}

// my original Trie struct
type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: &TrieNode{},
	}

}

// in this function we add word to our trie datastructure
func (t *Trie) Insert(word string) {

	currNode := t.Root // my root node

	word = strings.ToLower(word) // lower case it
	for _, ch := range word {
		idx := ch - 'a' // find idx for our children array

		if currNode.Children[idx] == nil {
			// means no node here
			currNode.Children[idx] = &TrieNode{} // initialize our node
		}
		currNode = currNode.Children[idx] // now move to next child node
	}
	currNode.IsEnd = true // word completed marking it true

}

// now search word from Trie
func (t *Trie) Search(word string) bool {
	currNode := t.Root // my root node

	word = strings.ToLower(word) // lower case it
	for _, ch := range word {
		idx := ch - 'a' // current idx

		if currNode.Children[idx] == nil {
			// no word found
			return false
		}
		currNode = currNode.Children[idx]
	}
	if currNode.IsEnd {
		return true // after traversal if end is true then word found
	}
	return false

}

// now deleting word from trie means simply unmark it

func (t *Trie) UnMark(word string) {
	currNode := t.Root // my root node

	word = strings.ToLower(word)
	for _, ch := range word {
		idx := ch - 'a'
		if currNode.Children[idx] == nil {

			return
		}
		currNode = currNode.Children[idx]
	}

	currNode.IsEnd = false // simply unmark it so next time no word found from that

}
