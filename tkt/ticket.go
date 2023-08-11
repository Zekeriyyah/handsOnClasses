package tkt

import (
	"fmt"
	"math/rand"
	"os"
)

var marsDist = 62100000

const (
	SecInHour = 3600 //seconds
	HrInDay   = 24   //hours
)

func GenTicket(n int) {
	//call a func to generate random spaceline
	//call a func to calc the duration of the travel base on randomly generated speed b/w 16-30km/s
	//call a func to generate mode of travel and modify the price

	fmt.Printf("%-20v %2v %-11v  %-3v\n", "Spaceline", "Days", "Trip Type", "Price($m)")
	fmt.Println("===============================================")
	for ; n > 0; n-- {
		tripType := TripType()
		speed := rand.Intn(15) + 16 //km/s within 16km/s to 30km/s

		fmt.Printf("%-20v %2d   %-11v  $%-3d\n", SpaceLine(), Duration(speed), tripType, CalPrice(tripType, speed))
	}
}

func SpaceLine() (val string) {
	switch sp := rand.Intn(3); sp {
	case 0:
		val = "Virgin Galactic"
	case 1:
		val = "Space Adventure"
	case 2:
		val = "Space X"
	default:
		fmt.Println("Error: invalid spaceline")
		os.Exit(1)
	}
	return val
}

func Duration(speed int) (noOfDays int) {
	noOfDays = ((marsDist / speed) / SecInHour) / HrInDay
	return
}

func TripType() string {
	t := rand.Intn(2)
	if t == 0 {
		return "Round-trip"
	} else {
		return "One-way"
	}
}

func CalPrice(trip string, speed int) (p int) {
	const minSp, maxSp = 16, 30       //maximum and minimum speed in km/s
	const minPrice, maxPrice = 36, 50 //max and min price

	//To get the price with condition: the faster the higher the price
	price := minPrice + speed - minSp
	if trip == "Round-trip" {
		price *= 2
	}
	return price
}
