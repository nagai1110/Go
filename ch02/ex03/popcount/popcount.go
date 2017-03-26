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

// 合計のポピュレーションカウントをループで計算する方式
func PopCountByLoop(x uint64) int {
	sumPC := 0
	for i := uint64(0); i < 8; i++ {
		sumPC += int(pc[byte(x>>(i*8))])
	}

	return sumPC
}
