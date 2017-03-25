// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nagai1110/Go/ch02/ex02/unitconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			value, err := strconv.ParseFloat(arg, 64)
			if err == nil {
				convertUnit(value)
			}
		}
	} else {
		var input float64
		fmt.Print("Input Number: ")
		fmt.Scan(&input)

		convertUnit(input)
	}
}

func convertUnit(unit float64) {
	fmt.Printf("Feet : %v -> Meter : %v\n", unit, unitconv.FeetToMeter(unitconv.Feet(unit)))
	fmt.Printf("Meter : %v -> Feet : %v\n", unit, unitconv.MeterToFeet(unitconv.Meter(unit)))
	fmt.Printf("Pond : %v -> Kilogram : %v\n", unit, unitconv.PondToKilogram(unitconv.Pond(unit)))
	fmt.Printf("Kilogram : %v -> Pond : %v\n", unit, unitconv.KilogramToPond(unitconv.Kilogram(unit)))
}
