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

	res := encode(a)

	fmt.Printf("encoded:\t%#x\n\t%08[1]b\n", res)

	res2 := decode(res)

	fmt.Printf("decoded:\t%#x\n\t%08[1]b\n\t%[1]d\n", res2)

	return

	_, err := os.Stdout.Write(res)

	if err != nil {
		panic(err)
	}
}

func encode(src uint64) []byte {
	if src == 0 {
		return nil
	}

	var res []byte

	for src > 0 {
		r := src & 0x7f // 0x7f a.k.a. 0b01111111

		src = src >> 7

		res = append(res, byte(r))
	}

	if len(res) > 1 {
		res[0] = res[0] | 0x80
	}

	return res
}

func decode(src []byte) uint {
	var res uint

	for i := len(src) - 1; i >= 0; i-- {
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
