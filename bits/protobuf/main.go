package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		return
	}

	a, _ := strconv.ParseUint(args[1], 10, 64)
	fmt.Println("incoming", a)
    prettyBytes("in bytes (BE):", a)
    fmt.Println("------------------------------")

	res := encode(a)

	fmt.Printf("encoded (LE):\t%#x\n\t%08[1]b\n", res)

	res2 := decode(res)

	fmt.Printf("decoded (BE):\t%#x\n", res2)
    prettyBytes("\t", res2)

	return

	_, err := os.Stdout.Write(res)

	if err != nil {
		panic(err)
	}
}

// encode in little endian byte order
func encode(src uint64) []byte {
	if src == 0 {
		return nil
	}

	var res []byte

	for src > 0 {
		r := src & 0x7f // 0x7f a.k.a. 0b0111_1111, how src % 128 also does the same is beyond me.

		src = src >> 7

        // Each byte in the varint has a continuation bit that indicates if the byte that follows it is part of the varint.
        // see: https://protobuf.dev/programming-guides/encoding/#varints
        // note: 
        // a sequence of continuation bits that denotes a group might look like 1, 1, 1, 0,
        // where last 0 tells this is the end of a sequence.
        // a single byte group has no continuation, so continueation bit sequence might look liek 1, 1, 1, 0, 0 
        // where last byte with `0` continuation bit is a standalone int.
        if src > 0 {
            r |= 0x80 // 0x80 a.k.a 0b1000_0000
        }

		res = append(res, byte(r))
	}

	return res
}


func decode(src []byte) uint {
	var res uint

	for i := len(src) - 2; i >= 0; i-- {
		res |= uint(src[i]&0x7f) << uint(7*i)
	}

	return res
}

func prettyBytes(m string, src any) {
	r := fmt.Sprintf("%064b", src)

	res := []string{}

	for i := 0; i < len(r); i += 8 {
		res = append(res, string(r[i:i+8]))
	}

	fmt.Println(m, strings.Join(res, " "))
}
