// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// 引数をビットシフトしながら最下位ビットの検査を64回繰り返す方式
func PopCountByBitShift(x uint64) int {
	var lsb uint64 = 0x00000001

	pc := 0
	for i := 0; i < 64; i++ {
		if x&lsb == 1 {
			pc++
		}

		x = x >> 1
	}

	return pc
}
