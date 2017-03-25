// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package tempconv

import (
	"testing"
)

func TestTempConv(t *testing.T) {
	var tests = []struct {
		c Celsius
		f Fahrenheit
		k Kelvin
	}{
		{20, 68, 293.15},
		{30, 86, 303.15},
	}

	for _, test := range tests {
		f := CToF(test.c)
		if f != test.f {
			t.Errorf("CToF(%f) = %v, want %v", test.c, f, test.f)
		}

		k := CToK(test.c)
		if k != test.k {
			t.Errorf("CToK(%f) = %v, want %v", test.c, k, test.k)
		}

		c := FToC(test.f)
		if c != test.c {
			t.Errorf("FToC(%f) = %v, want %v", test.f, c, test.c)
		}

		k = FToK(test.f)
		if k != test.k {
			t.Errorf("FToK(%f) = %v, want %v", test.f, k, test.k)
		}

		c = KToC(test.k)
		if c != test.c {
			t.Errorf("KToC(%f) = %v, want %v", test.k, c, test.c)
		}

		f = KToF(test.k)
		if f != test.f {
			t.Errorf("KToF(%f) = %v, want %v", test.k, f, test.f)
		}
	}
}
