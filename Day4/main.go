// ==== Going to make Browser tab manager
package main

import (
	"fmt"
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

// method to add node at front of the linkedlist
func (bl *BrowserList) addFront(tabdata Tab) {
	newTabNode := &Node{Data: tabdata} // creating a new node

	if bl.Head == nil { // means no node
		bl.Head = newTabNode
		bl.Tail = newTabNode
		return
	}
	newTabNode.Next = bl.Head
	bl.Head.Prev = newTabNode
	bl.Head = newTabNode
	fmt.Printf("Tab %v added successfully\n", tabdata.Title)

}

// method to add node at back of the linkedlist

func (bl *BrowserList) addBack(tabData Tab) {
	newTabNode := &Node{Data: tabData} // new Node

	if bl.Head == nil { // means no node in list
		bl.Head = newTabNode
		bl.Tail = newTabNode
		return
	}

	bl.Tail.Next = newTabNode
	newTabNode.Prev = bl.Tail
	bl.Tail = newTabNode
	fmt.Printf("Tab %v added successfully in Back\n", tabData.Title)
}

// method to close means delete tab from the linkedlist
// here we need to iterate hence delete tab is of o(n)
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

			fmt.Printf("Tab %v deleted successfully\n", title)
			return
		}
		currNode = currNode.Next
	}
	fmt.Printf("No tab %v found to delete\n", title)
}

// method to search particular tab from the linkedlist in which i used prefix tree i.e trie which we already learn

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

	bl := BrowserList{}
	bl.addBack(Tab{"google.com", "https://google.com"})
	bl.addBack(Tab{"github.com", "https://github.com"})
	bl.addFront(Tab{"golang.org", "https://golang.org"})
	bl.addBack(Tab{"stackoverflow.com", "https://stackoverflow.com"})

	bl.display() // print all tabs

	fmt.Println("Deleting github.com")
	bl.close("github.com")
	bl.display()

}
