package bst

import (
	"fmt"
)

func init() {
	fmt.Println("busstop package initialize...")
}

// Node object contain a pointer type data
type Node struct {
	data string // *Passenger
	next *Node
}

// Queue implementation using Node
type Queue struct {
	head *Node
	tail *Node
	size int
}

// Adder add node to queue
type Adder interface{ Add(node Node) }

// Add does not return anything
func (q *Queue) Add(node Node) {
	if q.head == nil {
		q.head = &node
		q.tail = &node
	} else {
		q.tail.next = &node
		q.tail = &node
	}
	q.size++
}

// Popper remove node from queue
type Popper interface{ Pop() *Node }

// Pop return pointer of the removed node
func (q *Queue) Pop() *Node {
	if q.head == nil {
		fmt.Println("Error: queue is empty")
		return nil
	}
	temp := q.head
	if q.size == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}
	q.size--
	return temp
}

// Printer print the value inside the struct
type Printer interface{ PrintD() }

// PrintD for q struct
func (q Queue) PrintD() {
	fmt.Printf("Current Queue: ")
	for i := q.head; i != nil; i = i.next {
		fmt.Printf("%v ", i.data)
	}
	fmt.Printf("\nCurrent Queue Info: %v\nHead: %v\nTail: %v\nSize: %v\n", q, q.head, q.tail, q.size)
}

// PrintD for bus stop struct
func (bStop BusStop) PrintD() {
	fmt.Println("------------------------------------")
	fmt.Printf("Bus Stop Name: %v\n", bStop.name)
	fmt.Printf("Waiting People: %v\n", bStop.waitingPassenger)
	bStop.q.PrintD()
	fmt.Println("------------------------------------")
}

// BusStop create a bus stop object
type BusStop struct {
	name             string
	waitingPassenger int
	q                Queue
}

// Test get feedback from busstop.go
func Test() {
	fmt.Println("busstop package succesfully initialize...")
}
