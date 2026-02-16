package main

import (
	"fmt"
	"math/rand"
)

// my song struct which
type Song struct {
	ID     int
	Title  string
	Singer string
}

type Node struct {
	Data Song
	Next *Node
}

// My song playlist
type SongPlaylist struct {
	Head *Node // adding from front
	Tail *Node // used for adding from back
}

var counter int        // global list counter . In future i can use atomic package in order to avoid race condition for concurrency purposes.
var shuffleList []Song // i am using this list only for shuffle function .This list will only modified in shuffle function when its len differs from counter

// func for constructing new node
func newNode(songData Song, NextNode *Node) *Node {
	return &Node{Data: songData, Next: NextNode}
}

// Method to implement add song at beginning of playlist

func (sp *SongPlaylist) addFront(songData Song) {
	//newSongNode := newNode(songData, sp.Head.Next)

	newSongNode := newNode(songData, sp.Head)
	if sp.Head == nil {
		sp.Tail = newSongNode
		sp.Head = newSongNode
	} else {
		sp.Head = newSongNode
	}

	counter++

	fmt.Printf("New Song %v added successfully\n", songData.ID)
}

// Method to implement add song at end of the playlist

func (sp *SongPlaylist) addBack(songData Song) {
	newSongNode := newNode(songData, nil)

	if sp.Tail == nil {
		sp.Tail = newSongNode
		sp.Head = newSongNode
	} else {
		sp.Tail.Next = newSongNode
		sp.Tail = newSongNode
	}
	counter++ // increase lenght

	fmt.Printf("New Song %v added successfully\n", songData.ID)

}

// Method to delete particular song from the playlist

func (sp *SongPlaylist) deleteSong(songId int) {
	// we need to travserse to node which has first this song then remove that node
	// if songs are empty

	if sp.Head == nil {
		fmt.Printf("Songs playlist is already empty\n")
		return
	}

	// this means first node itself is ID
	if sp.Head.Data.ID == songId {
		sp.Head = sp.Head.Next
		if sp.Head == nil { // means only element is there earlier now it is also removed
			sp.Tail = nil
		}
		return
	}

	// now for traversal

	currentNode := sp.Head
	for currentNode.Next != nil && currentNode.Next.Data.ID != songId {
		currentNode = currentNode.Next
	}
	// now current node is at my previous

	if currentNode.Next == nil {
		fmt.Printf("No song with ID %v found to delete \n", songId)
		return

	}

	currentNode.Next = currentNode.Next.Next
	if currentNode.Next == nil { // this is for last element case
		sp.Tail = currentNode
	}
	counter-- // decrease length

	fmt.Printf("Song %v deleted successfully \n", songId)
}

// Method to display song playlist

func (sp *SongPlaylist) DisplayPlaylist() {

	fmt.Printf("All %v songs are listed below:-\n", counter)
	currentNode := sp.Head

	for currentNode != nil {
		fmt.Printf("Song %v -- ( %v ) , Singer - %v\n", currentNode.Data.ID, currentNode.Data.Title, currentNode.Data.Singer)
		currentNode = currentNode.Next
	}
	fmt.Printf("\n")
}

// Method to shuffle song
// this function creates a random order of songs and then play but it don't change original order which created by human
func (sp *SongPlaylist) shuffle() {

	if sp.Head == nil {
		return
	}

	// append slices only when it differ from counter

	if len(shuffleList) != counter {
		shuffleList = nil // reset to len  = 0 and cap = 0
		currentNode := sp.Head

		for currentNode != nil {
			shuffleList = append(shuffleList, currentNode.Data)
			currentNode = currentNode.Next
		}

	}
	rand.Shuffle(len(shuffleList), func(i, j int) {

		shuffleList[i].ID, shuffleList[j].ID = shuffleList[j].ID, shuffleList[i].ID

	})
	fmt.Printf("Songs shuffled successfully\n")

	for _, v := range shuffleList {
		fmt.Printf("Song %v -- ( %v ) , Singer - %v\n", v.ID, v.Title, v.Singer)
	}

	fmt.Printf("\n")

}

// Method to merge two playlist and return merged playlist

func (sp *SongPlaylist) merge(p *SongPlaylist) {

	if sp.Head == nil {
		sp.Head = p.Head
		return
	}

	sp.Tail.Next = p.Head
	sp.Tail = p.Tail

	// Linkedlist merge succesfully

}

func main() {

	//fmt.Println("Creating Linkedlist project")
	sp := SongPlaylist{}

	sp.addFront(Song{ID: 1, Title: "Sorry ", Singer: "Justin Bieber"})
	sp.addFront(Song{ID: 2, Title: "Paradise", Singer: "Coldplay"})
	sp.addFront(Song{ID: 3, Title: "Sapphire", Singer: "Ed sheeran"})
	sp.DisplayPlaylist()
	sp.addBack(Song{ID: 4, Title: "On my way", Singer: "Alan Walker"})
	sp.addBack(Song{ID: 5, Title: "Believer", Singer: "Imagine Dragons"})
	sp.addBack(Song{ID: 6, Title: "Jungle book title", Singer: "Vivek Patel"})

	sp.DisplayPlaylist()
	sp.addFront(Song{ID: 7, Title: "Counting Stars", Singer: "xyz"})
	sp.DisplayPlaylist()

	//sp.deleteSong(7) // Tried to delete first element
	//sp.DisplayPlaylist()

	sp.deleteSong(6) // Tried to delete last element
	sp.DisplayPlaylist()

	//sp.deleteSong(1) // Tried to delete mIDdle  element
	//sp.DisplayPlaylist()

	//sp.deleteSong(10) // Tried to delete wrong element
	//sp.DisplayPlaylist()

	sp.shuffle()         // this is my custom shuffle function
	sp.DisplayPlaylist() // orignal list unchanged

	sp.addBack(Song{ID: 8, Title: "Random", Singer: "Bhaumil Panchal"})
	sp.DisplayPlaylist()

	sp.shuffle()
	sp.shuffle()

	sp.DisplayPlaylist()

	fmt.Printf("Now for merging two list\n")

	sp1 := SongPlaylist{}
	sp1.addBack(Song{ID: 9, Title: "Rand1", Singer: "Justin Bieber"})
	sp1.addBack(Song{ID: 10, Title: "Rand2", Singer: "Coldplay"})
	sp1.addFront(Song{ID: 11, Title: "Rand3", Singer: "Ed sheeran"})
	sp.merge(&sp1)

	sp.DisplayPlaylist()

}
