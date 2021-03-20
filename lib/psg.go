package rs

import (
	"math/rand"
)

// Passenger create a passenger object
type Passenger struct {
	Source      string
	Destination string
	Next        *Passenger
	waitingTime int
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

<<<<<<< Updated upstream
//NewPassenger Create a new passenger
func NewPassenger() *Passenger {
	var p *Passenger
	p = new(Passenger)
=======
//NewPassenger1 Create a new passenger
func NewPassenger() Passenger {
	var p Passenger
>>>>>>> Stashed changes
	return p
}

//GnrPsg Generate psg and add to bus stop
func GnrPsg(stopList []*BusStop, random1 int, psgr Passenger) {
	for i := 1; i < random1; i++ {
		psgr.Source = *&stopList[rand.Intn(10)].Name
		psgr.Destination = *&stopList[rand.Intn((len(stopList)-0-1)+1)].Name
		for j := 0; j < len(stopList)-1; j++ {
			if psgr.Source == *&stopList[j].Name {
				stopList[j].Q.Add(psgr)
				// fmt.Println(stopList[i].Name)
				// fmt.Println(stopList[i].Q.Size)
			}
		}
	}
}

//GnrPsgAt Generate psg and add to specific bus stop
<<<<<<< Updated upstream
func GnrPsgAt(stopList []*BusStop, stop string, inputPsg int, psgr *Passenger) {
=======
func GnrPsgAt(stopList []*BusStop, stop string, inputPsg int, psgr Passenger) {
>>>>>>> Stashed changes
	for i := 0; i < len(stopList); i++ {
		if stop == *&stopList[i].Name {
			for j := 0; j < inputPsg; j++ {
				stopList[i].Q.Add(psgr)
				// fmt.Println(stopList[i].Name)
				// fmt.Println(stopList[i].Q.Size)
			}
		}
	}
}

//Random number of int
func Random(min int, max int) int {
	return rand.Intn(max-min) + min
}

// Random code from https://gist.github.com/201Factory/5ef7c2d46cf848db16041cafa17ab054
