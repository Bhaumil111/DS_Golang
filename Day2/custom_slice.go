package main

import (
	"fmt"
)

// struct of my custom slice which has data means underlying array and len ( actual no of elements) and capacity ( how much slice allowed to grow)
type CustomSlice struct {
	Data []int
	Len  int
	Cap  int
}

// creating function to create new slice with inital capacity
func CreateNewSlice(initialCapacity int) *CustomSlice {
	return &CustomSlice{
		Data: make([]int, initialCapacity),
		Len:  0,
		Cap:  initialCapacity,
	}
}

// Append function append new element into slice. If len == cap here then reallocation ( Growth Phase ) happens
func (customSlice *CustomSlice) append(val int) {
	// when len == cap Growth phase happen
	if customSlice.Len == customSlice.Cap {
		newCapacity := customSlice.Cap * 2
		if newCapacity == 0 {
			newCapacity = 1
		}
		newSlice := make([]int, newCapacity)
		copy(newSlice, customSlice.Data)
		customSlice.Data = newSlice
		customSlice.Cap = newCapacity
		fmt.Println("Growth Phase completed . New Capacity is", customSlice.Cap)
	}
	customSlice.Data[customSlice.Len] = val // adding val
	customSlice.Len++

}

// This method prints about len  and capacity of the map
func (customSlice *CustomSlice) printDetails() {

	fmt.Printf("Length of slice is %v and Capacity of slice is %v\n", customSlice.Len, customSlice.Cap)
	fmt.Println("DATA of slice are", customSlice.Data)

}

func main() {

	initialCapacity := 0
	customSlice := CreateNewSlice(initialCapacity)

	fmt.Println("Initializing new custom slice with intiial Capacity ", customSlice)

	// Just adding element to simulate behavior
	n := 5

	for i := 0; i < n; i++ {
		customSlice.append(i + 1)  // appending element one by one
		customSlice.printDetails() //printing Details to show result

	}

}
