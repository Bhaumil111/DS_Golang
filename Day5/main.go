package main

import (
	"fmt"
	"hash/fnv"
)

type Node struct {
	key string
	val string
}

// one particular bucket make

const LoadFactor = 6.5 // This is my threshold loadfactor . live loadfactor calculate based on division of ( total no of key value pair )  /  total no of buckets

// when live loadfactor exceeds threshold then i double my capacity of Datalist in Custom maps

type Bucket struct { // this is main storage means a bucket hold real data means key value pair // all each bucket has overflow pointer which points to next bucket like linkedlist

	Data            *[8]Node // fixed 8 sized array in which KV stored
	OverflowPointer *Bucket  // overflow pointer which points to new bucket where our LF exceeds load factor

}

// my custom_map
type CustomMap struct {
	BucketList []Bucket // my maps has a list of buckets

}

// Method for calculating hash function based on the slice of bucket

func hashIndex(username string, size int) int {
	h := fnv.New32a()               // creating a hash calculator
	h.Write([]byte(username))       // john --- 'j' , 'o', 'h' , 'n' then based that each it calculate my hash number
	hashValue := h.Sum32()          // it give me hash
	idx := hashValue % uint32(size) // your idx based on size
	fmt.Printf("Calculated hash is %v  and idx is %v for string %v\n", hashValue, size, username)
	return int(idx)
}

func main() {

	fmt.Println("Implementing custom map in golang \n")

	// now

	a := hashIndex("Bhaumil", 8)
	b := hashIndex("Raj", 8)
	c := hashIndex("Manav", 8)
	d := hashIndex("Bhaumil", 8)
	e := hashIndex("Bhaumil", 8)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
