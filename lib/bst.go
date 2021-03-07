package rs

import "fmt"

// Queue implementation using Passenger node
type Queue struct {
	head *Passenger
	tail *Passenger
	size int
}

// Adder add node to queue
type Adder interface{ Add(node Passenger) }

// Add does not return anything
func (q *Queue) Add(node Passenger) {
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
type Popper interface{ Pop() *Passenger }

// Pop return pointer of the removed node
func (q *Queue) Pop() *Passenger {
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
type Printer interface{ printD() }

// printQ does not return anything
func (q Queue) printD() {
	fmt.Printf("Current Queue: ")
	for i := q.head; i != nil; i = i.next {
		fmt.Printf("%v ", i)
	}
	fmt.Printf("\nCurrent Queue Info: %v\nHead: %v\nTail: %v\nSize: %v\n", q, q.head, q.tail, q.size)
}

func (bStop BusStop) printD() {
	fmt.Println("------------------------------------")
	fmt.Printf("Bus Stop Name: %v\n", bStop.Name)
	fmt.Printf("Waiting People: %v\n", bStop.WaitingPassenger)
	bStop.q.printD()
	fmt.Println("------------------------------------")
}

// BusStop create a bus stop object
type BusStop struct {
	Name             string
	WaitingPassenger int
	q                Queue
	TimeTaken        int
}
