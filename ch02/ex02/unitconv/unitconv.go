// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package unitconv

import (
	"fmt"
	"math"
)

type Feet float64
type Meter float64
type Pond float64
type Kilogram float64

func (ft Feet) String() string     { return fmt.Sprintf("%gft", ft) }
func (m Meter) String() string     { return fmt.Sprintf("%gm", m) }
func (lb Pond) String() string     { return fmt.Sprintf("%glb", lb) }
func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg) }

func FeetToMeter(ft Feet) Meter { return Meter(Round(float64(ft) / 3.2808)) }

func MeterToFeet(m Meter) Feet { return Feet(Round(float64(m) * 3.2808)) }

func PondToKilogram(lb Pond) Kilogram { return Kilogram(Round(float64(lb) / 2.2046)) }

func KilogramToPond(kg Kilogram) Pond { return Pond(Round(float64(kg) * 2.2046)) }

// 小数点以下4桁目まで
func Round(f float64) float64 {
	shift := math.Pow(10, 4)
	return math.Floor(f*shift+.5) / shift
}
