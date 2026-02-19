package main

import (
	"fmt"
	"hash/fnv"
)

// user struct which hold username and password to store
type User struct {
	Username string
	Password string
}

// one particular bucket make

var counter = 0        // live counter for total no of key value pair
const loadFactor = 6.5 // This is my threshold load factor . live load factor calculate based on division of ( total no of key value pair )  /  total no of buckets
const bucketSize = 8   // there are total 8 elements inside bucket

// when live load factor exceeds threshold then I double my capacity of Datalist in Custom maps

type Bucket struct { // this is main storage means a bucket hold real data means key value pair // all each bucket has overflow pointer which points to next bucket like linkedlist

	Data            [bucketSize]User // fixed 8 sized array in which KV stored
	OverflowPointer *Bucket          // overflow pointer which points to new bucket where our LF exceeds load factor
}

// my custom_map
type CustomMap struct {
	BucketList []*Bucket // my maps have a list of buckets
	Cap        int       // Size of my Bucket
}

// initialize new map with initial capacity
func newMap(initialBuckets int) *CustomMap {
	mp := &CustomMap{
		BucketList: make([]*Bucket, initialBuckets),
		Cap:        initialBuckets,
	}
	// initialize map for each index
	for i := 0; i < initialBuckets; i++ {
		mp.BucketList[i] = &Bucket{}
	}
	return mp
}

// Method for calculating hash function based on the slice of bucket

func hashIndex(username string, size int) int {
	h := fnv.New32a()               // creating a hash calculator
	h.Write([]byte(username))       // john --- 'j' , 'o', 'h' , 'n' then based that each it calculates my hash number
	hashValue := h.Sum32()          // it gives me hash
	idx := hashValue % uint32(size) // your idx based on size
	//fmt.Printf("Calculated hash is %v  and idx is %v for string %v\n", hashValue, size, username)
	return int(idx)
}

// method for appending elements to bucket

func (mp *CustomMap) add(userData User) {

	liveLoadFactor := float64(counter) / float64(mp.Cap) // gettting live loadfactor
	//fmt.Printf("live load factor is %.1f\n", liveLoadFactor)

	if liveLoadFactor >= loadFactor {
		fmt.Printf("Load Factor limit exceeded \n")

		newCap := mp.Cap * 2
		if newCap == 0 {
			newCap = 1
		}

		newMapData := make([]*Bucket, newCap)
		copy(newMapData, mp.BucketList)
		mp.BucketList = newMapData
		mp.Cap = newCap

		for i := 0; i < mp.Cap; i++ {
			if mp.BucketList[i] == nil {
				mp.BucketList[i] = &Bucket{}
			}
		}
		fmt.Printf("Growth phase done successfully\n")

	}

	// calculating idx using hashIndex function
	idx := hashIndex(userData.Username, mp.Cap)
	//fmt.Println("Got this idx for storing data", idx)

	Bucket := mp.BucketList[idx] //   particular bucket per index

	for i := 0; i < 8; i++ {
		if Bucket.Data[i].Username == "" {
			Bucket.Data[i] = userData
			break
		}

	}
	counter++ // increasing my counter ok
	//fmt.Println("User Data entered successfully")
}

// method to authenticate user

func (mp *CustomMap) authenticateUser(userData User) {
	idx := hashIndex(userData.Username, mp.Cap) // go to particular hash

	Bucket := mp.BucketList[idx] // getting particular bucket

	flag := false
	for i := 0; i < 8; i++ {
		if Bucket.Data[i] == userData {
			flag = true
			break

		}
	}

	if flag {
		fmt.Printf("User %v found\n", userData.Username)
	} else {
		fmt.Printf("User %v not found\n", userData.Username)
	}

}

// method to display map elements

func (mp *CustomMap) display() {
	fmt.Println("Display map data")
	for i := 0; i < mp.Cap; i++ {
		fmt.Println(mp.BucketList[i].Data)
	}
	fmt.Println()
}

func main() {

	fmt.Printf("Implementing custom map in golang \n")

	mp := newMap(3) // initialize new map

	Users := []User{
		{Username: "alice", Password: "pass123"},
		{Username: "bob", Password: "bob@321"},
		{Username: "charlie", Password: "charlie12"},
		{Username: "david", Password: "david!45"},
		{Username: "emma", Password: "emma2024"},
		{Username: "frank", Password: "frank#77"},
		{Username: "grace", Password: "grace$88"},
		{Username: "harry", Password: "harry007"},
		{Username: "isla", Password: "isla111"},
		{Username: "jack", Password: "jack_999"},
		{Username: "kate", Password: "kate@000"},
		{Username: "leo", Password: "leoabc"},
		{Username: "maria", Password: "maria321"},
		{Username: "nathan", Password: "nat!456"},
		{Username: "olivia", Password: "olivia12"},
		{Username: "peter", Password: "peter777"},
		{Username: "quinn", Password: "quinn!1"},
		{Username: "rachel", Password: "rach@22"},
		{Username: "sam", Password: "sam_pass"},
		{Username: "tina", Password: "tina#45"},
		{Username: "uma", Password: "uma321"},
		{Username: "victor", Password: "vic!909"},
		{Username: "will", Password: "will000"},
		{Username: "xavier", Password: "xav@123"},
		{Username: "yara", Password: "yara777"},
		{Username: "zane", Password: "zane000"},
		{Username: "aaron", Password: "aaron123"},
		{Username: "bella", Password: "bella321"},
		{Username: "cody", Password: "cody!99"},
		{Username: "diana", Password: "diana88"},
		{Username: "edward", Password: "ed@456"},
		{Username: "fiona", Password: "fio777"},
		{Username: "george", Password: "geo321"},
		{Username: "henry", Password: "hen!222"},
		{Username: "ivan", Password: "ivan001"},
		{Username: "julia", Password: "jul@123"},
		{Username: "kevin", Password: "kev456"},
		{Username: "linda", Password: "lin789"},
		{Username: "mason", Password: "mas@333"},
		{Username: "noah", Password: "noa222"},
		{Username: "oscar", Password: "osc!111"},
		{Username: "paula", Password: "pau000"},
		{Username: "quentin", Password: "que777"},
		{Username: "rose", Password: "ros321"},
		{Username: "steve", Password: "ste!222"},
		{Username: "tom", Password: "tom123"},
		{Username: "ursula", Password: "urs456"},
		{Username: "vince", Password: "vin789"},
		{Username: "wendy", Password: "wen000"},
		{Username: "zoe", Password: "zoe111"},
	}

	for i := 0; i < 20; i++ {
		mp.add(Users[i])
	}

	mp.display()
	mp.authenticateUser(Users[0])
	mp.authenticateUser(Users[1])
	mp.authenticateUser(Users[2])
	fakeUser1 := User{"fake", "1234"}
	fakeUser2 := User{"Vishrut", "1234"}

	mp.authenticateUser(fakeUser1)
	mp.authenticateUser(fakeUser2)

	for i := 20; i < 40; i++ {
		mp.add(Users[i])
	}

	mp.display() // display current data

	mp.authenticateUser(Users[10])

}
