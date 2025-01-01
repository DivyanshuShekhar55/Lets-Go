package main

import (
	"fmt"
	"time"
)

// **********************************************************************
// Go Routines
func basicRoutine() {

	// by using the keyword 'go' we can delegate tasks on a new thread
	go func() {
		fmt.Println("Hello from the new thread")

	}()
	go func() {
		fmt.Println("Hello from another thread")
	}()
}

// ************************************************************************
// Go Channels = store output of a thread in a Queue, this Queue is channel
func basicChannel() {

	// queue will hold int type of data, note that queue is stored on the main thread
	myChannel := make(chan int)
	go func() {
		// some heavy calculations in a new thread occurs, the result = 56 is stored
		myChannel <- 56
	}()

	// x var on the main thread is created and the value from myChannel is stored in it
	x := <-myChannel
	fmt.Println(x)
}

// ************************************************************************
func advancedChannel() {

	// note that following channel is Unbuffered Channel i.e, can hold one value only at a time
	// once a value is stored, we need to "pop" it from queue, before storing another one
	myChannel := make(chan int)

	go func() {
		// storing the value from this thread
		myChannel <- 88
	}()
	go func() {
		myChannel <- 21
	}()
	go func() {
		myChannel <- 34
	}()

	var x int

	// pop out first value that is stored in the queue
	// only then we can store the next value, so storing / popping-out is a 'blocking action'
	// i.e, the other 2 threads will keep waiting to store their values until this part has been finished
	x = <-myChannel
	fmt.Println(x)
	// the other two thread will now continue their work of storage

	// pop the result of the thread that finished next and stored its value in the channel
	x = <-myChannel
	fmt.Println(x)

	x = <-myChannel
	fmt.Println(x)

}

// **********************************************************************
// multiple channels and select ...
func multiChannels() {

	// with select we can do pretty much ... following is like .WaitAny method of C# or
	channel_one := make(chan string)
	channel_two := make(chan string)

	go func() {
		// do some work ...
		channel_one <- "thank you"
	}()

	go func() {
		channel_two <- "arigato"
	}()

	select {
	case msg_from_channel_one := <-channel_one:
		fmt.Println("first thread did its work", msg_from_channel_one)
	case msg_from_channel_two := <-channel_two:
		fmt.Println("second thread did its work", msg_from_channel_two)
	}
}

func main() {
	basicRoutine()
	time.Sleep(time.Second * 2) // wait for the thread actions to finish
	fmt.Println("Hello from the main thread")

	basicChannel()

	advancedChannel()

	multiChannels()
}
