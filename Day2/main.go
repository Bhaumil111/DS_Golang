package main

import (
	"fmt"
	"sort"
	"time"
)

type videoPacket struct {
	Data       string // here using byte because it is of 8 bits and it is common convention to use bytes slices to store raw data
	SequenceNo uint64 // order of this packet in stream
	Duration   uint64 // how long this segment should be played in seconds
}

type VideoBuffer struct {
	Buffer []videoPacket // This is main buffer which store all the packets
}

func (videoBuffer *VideoBuffer) addPacket(packet videoPacket) {
	videoBuffer.Buffer = append(videoBuffer.Buffer, packet)
	fmt.Println("Sending packet ", packet.SequenceNo)
}

func (videoBuffer *VideoBuffer) lengthOfBuffer() {
	fmt.Println("Current Size of the Buffer ", len(videoBuffer.Buffer))
	fmt.Println("Capacity of the Buffer ", cap(videoBuffer.Buffer))
}

func (videoBuffer *VideoBuffer) playVideo() {
	// first task is to sort it according to the sequence number

	sort.Slice(videoBuffer.Buffer, func(seq1, seq2 int) bool {
		return videoBuffer.Buffer[seq1].SequenceNo < videoBuffer.Buffer[seq2].SequenceNo
	}) // sort my buffer

	// just loop through and print data according to duration (seconds)

	fmt.Println("Playing Video....")
	for i := 0; i < len(videoBuffer.Buffer); i++ {
		fmt.Println(videoBuffer.Buffer[i].Data)
		//time.Sleep(videoBuffer.Buffer[i].Duration * time.Second)

		time.Sleep(time.Duration(videoBuffer.Buffer[i].Duration) * time.Second)

		//time.Sleep(-1 *time.Second)

	}

}

// this function send numOfPackets to slices

func main() {

	fmt.Println("Video streaming buffer")

	videoBuffer := VideoBuffer{}

	packet1 := videoPacket{
		Data:       "Hello my name is Bhaumil",
		SequenceNo: 1,
		Duration:   1,
	}
	packet2 := videoPacket{
		Data:       "Please review my movie",
		SequenceNo: 3,
		Duration:   3,
	}

	videoBuffer.addPacket(packet1)
	//videoBuffer.lengthOfBuffer()
	videoBuffer.addPacket(packet2)
	//videoBuffer.lengthOfBuffer()

	videoBuffer.playVideo()

}
