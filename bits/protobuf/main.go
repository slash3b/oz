package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		return
	}

	a := uintsfromargs(args)

	fmt.Println("incoming", a)
	fmt.Println("------------------------------")

	res := encode(a)

	fmt.Printf("varint encoded:\t%#x\n", res)

	b := decode(res)

	fmt.Printf("decoded:\t%#x\n", b)

	fmt.Println("------------------------------")
	fmt.Println("outgoing", b)

	if !slices.Equal(a, b) {
		os.Exit(1)
	}
}

// encode in little endian byte order
func encode(src []uint) []byte {
	if len(src) == 0 {
		return nil
	}

	var res []byte

	for _, v := range src {
		for v > 0 {
			r := v & 0x7f // 0x7f a.k.a. 0b0111_1111, how src % 128 also does the same is beyond me.

			v = v >> 7

			// Each byte in the varint has a continuation bit that indicates if the byte that follows it is part of the varint.
			// see: https://protobuf.dev/programming-guides/encoding/#varints
			// note:
			// a sequence of continuation bits that denotes a group might look like 1, 1, 1, 0,
			// where last 0 tells this is the end of a sequence.
			// a single byte group has no continuation, so continueation bit sequence might look liek 1, 1, 1, 0, 0
			// where last byte with `0` continuation bit is a standalone int.
			if v > 0 {
				r |= 0x80 // 0x80 a.k.a 0b1000_0000
			}

			res = append(res, byte(r))
		}
	}

	return res
}

func decode(src []byte) []uint {
	var res []uint

	a := uint(0)
	counter := 0 // byte shift counter

	for i := 0; i < len(src); {
		if src[i]&0x80 == 0 {
			res = append(res, uint(src[i]&0x7f))

			i++

			continue
		}

		a |= uint(src[i]&0x7f) << uint(7*counter)
		counter++
		i++

		// rely on fact that termination bit should be set to 0
		if src[i]&0x80 == 0 {
			a |= uint(src[i]&0x7f) << uint(7*counter)

			res = append(res, a)

			// reset
			a = 0
			counter = 0

			i++
		}
	}

	return res
}

func uintsfromargs(args []string) []uint {
	res := make([]uint, len(args[1:]))

	for i, v := range args[1:] {
		a, _ := strconv.ParseUint(v, 10, 64)

		res[i] = uint(a)
	}

	return res
}
