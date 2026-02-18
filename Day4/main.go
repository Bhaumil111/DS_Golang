// ==== Going to make Browser tab manager
package main

import (
	"fmt"

	"example.com/Day4/DataStructures/Tries"
)

type Tab struct {
	Title string
	Url   string
}

type Node struct {
	Data Tab
	Next *Node
	Prev *Node
}

// my main root double linkedlist in which i add tabs
type BrowserList struct {
	Head *Node
	Tail *Node // using specific tail node in list for o(1) insertion at back
}

var trie = Tries.NewTrie() // global init trie

// method to add node at front of the linkedlist

// overal time complexity is O(length of title ) and space is O(1)
func (bl *BrowserList) addFront(tabData Tab) {
	newTabNode := &Node{Data: tabData} // creating a new node

	if bl.Head == nil { // means no node
		bl.Head = newTabNode
		bl.Tail = newTabNode
	} else {

		newTabNode.Next = bl.Head
		bl.Head.Prev = newTabNode
		bl.Head = newTabNode
	}
	trie.Insert(tabData.Title)
	fmt.Printf("Tab %v added successfully\n", tabData.Title)

}

// method to add node at back of the linkedlist
// overal time complexity is O(length of title ) and space is O(1)

func (bl *BrowserList) addBack(tabData Tab) {
	newTabNode := &Node{Data: tabData} // new Node

	if bl.Head == nil { // means no node in list
		bl.Head = newTabNode
		bl.Tail = newTabNode
	} else {

		bl.Tail.Next = newTabNode
		newTabNode.Prev = bl.Tail
		bl.Tail = newTabNode
	}
	trie.Insert(tabData.Title)
	fmt.Printf("Tab %v added successfully in Back\n", tabData.Title)
}

// method to close means delete tab from the linkedlist
// here we need to iterate hence delete tab is of o(n)
// overall time complexity is O(n+l) and space O(1)

func (bl *BrowserList) close(title string) {
	currNode := bl.Head

	for currNode != nil {

		if currNode.Data.Title == title {

			if currNode.Prev != nil {
				currNode.Prev.Next = currNode.Next
			} else {
				bl.Head = currNode.Next
			}

			if currNode.Next != nil {
				currNode.Next.Prev = currNode.Prev
			} else {
				bl.Tail = currNode.Prev
			}

			trie.UnMark(title)
			fmt.Printf("Tab %v deleted successfully\n", title)
			return
		}
		currNode = currNode.Next
	}
	fmt.Printf("No tab %v found to delete\n", title)
}

// method to search particular tab from the linkedlist in which i used prefix tree i.e trie which we already learn
// overall time complexity is O(lenght of title)
func (bl *BrowserList) SearchTab(title string) {
	if trie.Search(title) { // calling my trie search function means o(title length)

		fmt.Printf("Word %v found\n", title)
	} else {
		fmt.Printf("Word %v not found\n", title)
	}

}

// method to display all tabs

func (bl *BrowserList) display() {

	fmt.Println("All Tabs ")
	if bl.Head == nil {
		fmt.Printf("There are no tabs open in browser\n")
		return
	}
	currentNode := bl.Head

	for currentNode != nil {
		fmt.Printf("Tab %v with url %v\n ", currentNode.Data.Title, currentNode.Data.Url)
		currentNode = currentNode.Next
	}
	fmt.Printf("\n")
}

func main() {

	bl := BrowserList{} // my browserlist

	bl.addBack(Tab{"github", "https://github.com"})
	bl.addBack(Tab{"google", "https://google.com"})
	bl.addFront(Tab{"golang", "https://golang.org"})
	bl.addBack(Tab{"StackOverflow", "https://stackoverflow.com"})

	bl.display() // print all tabs

	fmt.Println("Deleting github.com")
	bl.display()
	bl.SearchTab("stackoverflow")
	bl.SearchTab("github")
	bl.SearchTab("cpp")    // not found case
	bl.SearchTab("google") // finding google it found it
	bl.close("google")     // delete
	bl.SearchTab("google") // now google not found

}
