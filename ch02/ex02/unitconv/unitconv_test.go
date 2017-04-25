// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package unitconv

import (
	"testing"
)

func TestLengthUnitConv(t *testing.T) {
	var tests = []struct {
		ft Feet
		m  Meter
	}{
		{1, 0.3048},
		{6.5616, 2},
	}

	for _, test := range tests {
		m := FeetToMeter(test.ft)
		if m != test.m {
			t.Errorf("FeetToMeter(%v) = %v, want %v", test.ft, m, test.m)
		}

		ft := MeterToFeet(test.m)
		if ft != test.ft {
			t.Errorf("MeterToFeet(%v) = %v, want %v", test.m, ft, test.ft)
		}
	}
}

func TestWeightUnitConv(t *testing.T) {
	var tests = []struct {
		lb Pond
		kg Kilogram
	}{
		{1, 0.4536},
		{4.4092, 2},
	}

	for _, test := range tests {
		kg := PondToKilogram(test.lb)
		if kg != test.kg {
			t.Errorf("PondToKilogram(%v) = %v, want %v", test.lb, kg, test.kg)
		}

		lb := KilogramToPond(test.kg)
		if lb != test.lb {
			t.Errorf("KilogramToPond(%v) = %v, want %v", test.kg, lb, test.lb)
		}
	}
}
