package main

import (
	"fmt"
    "math"
)

func main() {

	// encode("hello")
	encode("bye")
}

func encode(s string) float64 {

	fmt.Printf("input %s %#x\n", s, []byte(s))

	if len(s) > 6 {
		return 0
	}

	var c uint

	c |= 0x7f_f0_00_00_00_00_00_00 // set 11 bits on
	c |= 0x00_08_00_00_00_00_00_00 // set 1 bit permanently

	fmt.Printf("%064b\n", c)

	fmt.Println("len is:", len(s))

	c |= uint(len(s)) << 48

	if len(s) > 6 {
		return 0
	}

	for i := 0; i < len(s); i++ {
		c |= uint(s[i]) << uint(8*i)
	}

	fmt.Printf("encoded: %#x\n", c)

	fmt.Println("reverting")

	res := []byte{}

	encodedlen := int(c & (0x07 << 48) >> 48)

	for i := 0; i < encodedlen; i++ {
		// mask out particular byte
		b := c & uint(0xff<<(8*i))

		// shift it back
		b >>= (8 * i)

		fmt.Printf(">>>> %T %#[1]v\n", b)

		res = append(res, byte(b))
	}

	fmt.Printf("reverted %s\n", string(res))

	return c
}
