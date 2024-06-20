package main

import (
	"fmt"
	"os"
)

func main() {
	b, err := os.ReadFile("./color-convert/simple.css")
	if err != nil {
		panic(err)
	}

	c := []byte{}

	for i := 0; i < len(b); {
		if b[i] != byte('#') {
			c = append(c, b[i])
			i++
			continue
		}

		j := i + 1

		for {
			if b[j] == byte('\n') {

				h := b[i+1 : j-1]

				r, g, b := conv(h)

				s := fmt.Sprintf("rgb(%v %v %v);", r, g, b)

                c = append(c, []byte(s)...)

				break
			}

			j++
		}

		i = j
	}

	fmt.Println(string(c))

	os.WriteFile("./color-convert/simple_expected__.css", c, 0o644)
}

// ascii to byte
var m = map[byte]byte{
	0x30: 0x0,
	0x31: 0x1,
	0x32: 0x2,
	0x33: 0x3,
	0x34: 0x4,
	0x35: 0x5,
	0x36: 0x6,
	0x37: 0x7,
	0x38: 0x8,
	0x39: 0x9,
	0x61: 0xa,
	0x62: 0xb,
	0x63: 0xc,
	0x64: 0xd,
	0x65: 0xe,
	0x66: 0xf,
}

func conv(c []byte) (byte, byte, byte) {
	if len(c)%3 != 0 {
		panic(fmt.Sprintf("invalid byte seq given: %x", c))
	}

	for i := 0; i < len(c); i++ {
		v, ok := m[c[i]]
		if !ok {
			panic(fmt.Sprintf("unexpected %v", c[i]))
		}

		c[i] = v
	}

	switch len(c) {
	case 6:
		return (c[0] << 4) | c[1], (c[2] << 4) | c[3], (c[4] << 4) | c[5]
	default:
		return 0, 0, 0
	}

}
