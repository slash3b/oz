package main

import (
	"fmt"
	"math"
)

func main() {
	f := encode("bye")
	fmt.Println("encoded:", f)

	res := decode(f)
	fmt.Println("decoded:", res)
}

func encode(s string) float64 {
	if len(s) > 6 {
		return 0
	}

	var c uint64

	c |= 0x7f_f0_00_00_00_00_00_00 // set 11 bits on
	c |= 0x00_08_00_00_00_00_00_00 // set 1 bit permanently

	c |= uint64(len(s)) << 48

	for i := 0; i < len(s); i++ {
		c |= uint64(s[i]) << (8 * i)
	}

	return math.Float64frombits(c)
}

func decode(f float64) string {
	fbs := math.Float64bits(f)

	res := []byte{}

	encodedlen := int(fbs & (0x07 << 48) >> 48)

	for i := 0; i < encodedlen; i++ {
		// mask out particular byte and shift it back
		b := fbs & uint64(0xff<<(8*i)) >> (8 * i)

		res = append(res, byte(b))
	}

	return string(res)
}
