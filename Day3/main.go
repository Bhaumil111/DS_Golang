package main

import "fmt"

type Song struct {
	Id     int
	Name   string
	Singer string
}

type Node struct {
	Data Song
	Next *Node
}

// My song playlist
type SongPlaylist struct {
	Head *Node
	Tail *Node
}

// func for constructing new node
func newNode(songData Song, NextNode *Node) *Node {
	return &Node{Data: songData, Next: NextNode}
}

// Method to implement add song at beginning of playlist

func (songPlaylist *SongPlaylist) addSongAtBeginning(songData Song) {
	//newSongNode := newNode(songData, songPlaylist.Head.Next)

	newSongNode := newNode(songData, songPlaylist.Head)
	if songPlaylist.Head == nil {
		songPlaylist.Tail = newSongNode
		songPlaylist.Head = newSongNode
	} else {
		songPlaylist.Head = newSongNode
	}

	fmt.Printf("New Song %v added successfully\n", songData.Id)
}

// Method to implement add song at end of the playlist

func (songPlaylist *SongPlaylist) addSongAtLast(songData Song) {
	newSongNode := newNode(songData, nil)

	if songPlaylist.Tail == nil {
		songPlaylist.Tail = newSongNode
		songPlaylist.Head = newSongNode
	} else {
		songPlaylist.Tail.Next = newSongNode
		songPlaylist.Tail = newSongNode
	}

	fmt.Printf("New Song %v added successfully\n", songData.Id)

}

// Method to delete particular song from the playlist

func (songPlaylist *SongPlaylist) deleteSong(songId int) {
	// we need to travserse to node which has first this song then remove that node
	// if songs are empty

	if songPlaylist.Head == nil {
		fmt.Printf("Songs playlist is already empty\n")
		return
	}

	// this means first node itself is id
	if songPlaylist.Head.Data.Id == songId {
		songPlaylist.Head = songPlaylist.Head.Next
		if songPlaylist.Head == nil { // means only element is there earlier now it is also removed
			songPlaylist.Tail = nil
		}
		return
	}

	// now for traversal

	currentNode := songPlaylist.Head
	for currentNode.Next != nil && currentNode.Next.Data.Id != songId {
		currentNode = currentNode.Next
	}
	// now current node is at my previous

	if currentNode.Next == nil {
		fmt.Printf("No song with id %v found to delete \n", songId)
		return

	}

	currentNode.Next = currentNode.Next.Next
	if currentNode.Next == nil { // this is for last element case
		songPlaylist.Tail = nil
	}

	fmt.Printf("Song %v deleted successfully \n", songId)
}

// Method to display song playlist

func (songPlaylist *SongPlaylist) DisplayPlaylist() {
	fmt.Printf("All Songs are listed below:- \n")
	currentNode := songPlaylist.Head

	for currentNode != nil {
		fmt.Printf("Song %v -- ( %v ) , Singer - %v\n", currentNode.Data.Id, currentNode.Data.Name, currentNode.Data.Singer)
		currentNode = currentNode.Next
	}
}

// Method to shuffle song
// Method to merge two playlist and return merged playlist

func main() {

	//fmt.Println("Creating Linkedlist project")
	songPlaylist := SongPlaylist{}

	songPlaylist.addSongAtBeginning(Song{Id: 1, Name: "Sorry ", Singer: "Justin Bieber"})
	songPlaylist.addSongAtBeginning(Song{Id: 2, Name: "Paradise", Singer: "Coldplay"})
	songPlaylist.addSongAtBeginning(Song{Id: 3, Name: "Sapphire", Singer: "Ed sheeran"})
	songPlaylist.DisplayPlaylist()
	songPlaylist.addSongAtLast(Song{Id: 4, Name: "On my way", Singer: "Alan Walker"})
	songPlaylist.addSongAtLast(Song{Id: 5, Name: "Believer", Singer: "Imagine Dragons"})
	songPlaylist.addSongAtLast(Song{Id: 6, Name: "Jungle book title", Singer: "Vivek Patel"})

	songPlaylist.DisplayPlaylist()
	songPlaylist.addSongAtBeginning(Song{Id: 7, Name: "Counting Stars", Singer: "xyz"})
	songPlaylist.DisplayPlaylist()

	//songPlaylist.deleteSong(7) // Tried to delete first element
	//songPlaylist.DisplayPlaylist()

	//songPlaylist.deleteSong(6) // Tried to delete last element
	//songPlaylist.DisplayPlaylist()

	//songPlaylist.deleteSong(1) // Tried to delete middle  element
	//songPlaylist.DisplayPlaylist()

	songPlaylist.deleteSong(10) // Tried to delete wrong element
	songPlaylist.DisplayPlaylist()
}
