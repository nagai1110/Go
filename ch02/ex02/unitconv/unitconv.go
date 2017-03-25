// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package unitconv

import (
	"fmt"
)

type Feet float64
type Meter float64
type Pond float64
type Kilogram float64

func (ft Feet) String() string     { return fmt.Sprintf("%gft", ft) }
func (m Meter) String() string     { return fmt.Sprintf("%gm", m) }
func (lb Pond) String() string     { return fmt.Sprintf("%glb", lb) }
func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg) }

func FeetToMeter(ft Feet) Meter { return Meter(ft / 3.2808) }

func MeterToFeet(m Meter) Feet { return Feet(m * 3.2808) }

func PondToKilogram(lb Pond) Kilogram { return Kilogram(lb / 2.2046) }

func KilogramToPond(kg Kilogram) Pond { return Pond(kg * 2.2046) }
