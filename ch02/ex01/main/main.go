// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"

	"github.com/nagai1110/Go/ch02/ex01/tempconv"
)

func main() {
	c := tempconv.Celsius(20)
	f := tempconv.Fahrenheit(68)
	k := tempconv.Kelvin(293.15)

	fmt.Printf("C : %v -> F : %v, K : %v\n", c, tempconv.CToF(c), tempconv.CToK(c))
	fmt.Printf("F : %v -> C : %v, K : %v\n", f, tempconv.FToC(f), tempconv.FToK(f))
	fmt.Printf("K : %v -> C : %v, F : %v\n", k, tempconv.KToC(k), tempconv.KToF(k))
}
