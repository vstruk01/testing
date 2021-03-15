package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountCycle(x uint64) int {
	var counter int

	for ; x != 0; x >>= 8 {
		counter += int(pc[byte(x)])
	}
	return counter
}

func PopCountHard(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountEachBit(x uint64) int {
	var counter int

	for ; x != 0; x >>= 1 {
		counter += int(x & 1)
	}
	return counter
}
