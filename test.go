package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	rs "rs/lib"
	"strconv"
	"strings"
	"time"
)

var (
	mainInput string
	mainCmd   []string
	err       error
	tStopLst  []*rs.BusStop
	tMin      int
	tHr       int
)

// This function should print out the list of created bus stop
func bstList() bool {
	if len(tStopLst) == 0 {
		fmt.Println("Error: no item in bus stop list")
		return false
	}
	fmt.Printf("List of bus stops: ")
	for _, v := range tStopLst {
		fmt.Printf("%v ", v.Name)
	}
	fmt.Printf("\n")
	return true
}

// This function should print out the data of the target bus stop
func bstGet() bool {
	// Check parameters
	if len(mainCmd) < 2 {
		fmt.Println("Error: invalid parameter; bstGet targetStop")
		return false
	}

	// Find and print
	for _, v := range tStopLst {
		if v.Name == mainCmd[1] {
			temp := *v
			fmt.Printf("Bus Stop Name: %v\n", temp.Name)
			fmt.Printf("Total Passengers: %v\n", temp.Q.Size)
			fmt.Printf("List of passengers:\n")
			for i := temp.Q.Head; i != nil; i = i.Next {
				fmt.Printf("Source: %v, Dest: %v, Next: %v\n", i.Source, i.Destination, i.Next)
			}
			return true
		}
	}
	fmt.Println("Error: bus stop with such name does not exist")
	return false
}

// This function should create a new bus stop and add it to the list
func bstCreate() bool {
	// Check parameters
	if len(mainCmd) < 2 {
		fmt.Println("Error: invalid parameter; bstCreate name")
		return false
	}
	// Check for duplicated stop
	for _, v := range tStopLst {
		if v.Name == mainCmd[1] {
			fmt.Println("Error: bus stop with such name already existed")
			return false
		}
	}
	tStopLst = append(tStopLst, &rs.BusStop{Name: mainCmd[1]})
	fmt.Printf("Succesfully created a new bus stop with name %v...\n", mainCmd[1])
	return true
}

// This function should add psg into queue of target bus stop
func bstAdd() bool {
	// Check parameters
	if len(mainCmd) < 3 {
		fmt.Println("Error: invalid parameter; bstAdd targetStop psgValue")
		return false
	}
	// Check target list
	if len(tStopLst) < 2 {
		fmt.Println("Error: insufficient amount of bus stop")
		return false
	}
	// Convert second parameter into int
	psgNum, err := strconv.Atoi(mainCmd[2])
	if err != nil {
		fmt.Printf("Error: invalid argument %v\n", mainCmd[2])
		return false
	}
	// Get target bus stop for correctness check
	var initSize int
	for _, v := range tStopLst {
		if v.Name == mainCmd[1] {
			initSize = v.Q.Size
			break
		}
	}

	// Create a new passenger
	start := time.Now()
	p := rs.NewPassenger()
	rand.Seed(time.Now().UnixNano())
	rs.GnrPsgAt(tStopLst, mainCmd[1], psgNum, p)
	end := time.Since(start)

	// Check feedback
	for _, v := range tStopLst {
		if v.Name == mainCmd[1] {
			temp := *v
			if temp.Q.Size-initSize != psgNum {
				fmt.Println("Error: passengers quantity is incorrect")
				return false
			}
			for temp.Q.Size != 0 {
				check := temp.Q.Pop()
				if check.Source == check.Destination {
					fmt.Printf("Error: duplicate src and dst at %v\n", v)
					return false
				}
			}
			fmt.Printf("Added passengers successsfully\n")
			fmt.Printf("Time taken: %v\n", end)
			return true
		}
	}
	fmt.Println("Error: adding failed")
	return false
}

// This function should add psg into queue of random bus stop
func bstAddRd() bool {
	// Check parameters
	if len(mainCmd) < 2 {
		fmt.Println("Error: invalid parameter; bstAddRd psgValue")
		return false
	}
	// Check target list
	if len(tStopLst) < 2 {
		fmt.Println("Error: insufficient amount of bus stop")
		return false
	}
	// Convert second parameter into int
	psgNum, err := strconv.Atoi(mainCmd[1])
	if err != nil {
		fmt.Printf("Error: invalid argument %v\n", mainCmd[1])
		return false
	}
	// Get target bus stop for correctness check
	var initSize int
	for _, v := range tStopLst {
		initSize += v.Q.Size
	}

	// Create a new passenger
	start := time.Now()
	p := rs.NewPassenger()
	rand.Seed(time.Now().UnixNano())
	rs.GnrPsg(tStopLst, psgNum, p)
	end := time.Since(start)

	// Check feedback
	var finalSize int
	for _, v := range tStopLst {
		finalSize += v.Q.Size
		temp := *v
		for temp.Q.Size != 0 {
			check := temp.Q.Pop()
			if check.Source == check.Destination {
				fmt.Printf("Error: duplicate src and dst at %v\n", v)
				return false
			}
		}
	}
	if finalSize-initSize != psgNum {
		fmt.Println("Error: passengers quantity is incorrect")
		fmt.Printf("%v %v %v\n", finalSize, initSize, psgNum)
		return false
	}
	fmt.Printf("Added passengers successsfully\n")
	fmt.Printf("Time taken: %v\n", end)
	return true
}

// This function should remove psg out of queue of target bus stop
func bstRmv() bool {
	return true
}

func tick() bool {
	start := time.Now()
	fmt.Println("Starting the clock...")
	tHr = 0
	tMin = 0
	rs.ConTimeTick(&tHr, &tMin)
	end := time.Since(start)
	if tHr == 24 && tMin == 0 {
		fmt.Println("Clock check successful")
		fmt.Printf("Time taken: %v\n", end)
		return true
	} else {
		fmt.Println("Error: clock check failed")
		return false
	}
}

// This function should add car into the map
func carAdd() bool {
	return true
}

func help() bool {
	fmt.Println("help")
	fmt.Println("bstList")
	fmt.Println("bstGet")
	fmt.Println("bstCreate")
	fmt.Println("bstAdd")
	fmt.Println("bstAddRd")
	fmt.Println("bstRmv")
	fmt.Println("timeTick")
	fmt.Println("carAdd")
	return true
}

func main() {
	// Setup
	reader := bufio.NewReader(os.Stdin)

	// Map functions
	var cmdMap = map[string](func() bool){
		"bstList":   bstList,
		"bstGet":    bstGet,
		"bstCreate": bstCreate,
		"bstAdd":    bstAdd,
		"bstAddRd":  bstAddRd,
		"bstRmv":    bstRmv,
		"timeTick":  tick,
		"carAdd":    carAdd,
		"help":      help,
	}

	fmt.Println("Test Drive Initiated...!")

	// Simple shell
	for {
		// Take Input and Split Argument
		fmt.Printf("> ")
		mainInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error: keyboard interrupt")
			continue
		}
		mainInput = strings.TrimSpace(mainInput)
		mainCmd = strings.Split(mainInput, " ")
		exeTarget, ok := cmdMap[mainCmd[0]]
		if !ok {
			fmt.Println("Error: command does not exist; Try help for list of commands")
			continue
		}
		exeTarget()
	}
}
