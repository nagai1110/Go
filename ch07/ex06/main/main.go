// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"flag"
	"fmt"

	"github.com/nagai1110/Go/ch07/ex06/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
