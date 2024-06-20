package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	b, err := os.ReadFile("./color-convert/advanced.css")
	if err != nil {
		panic(err)
	}

	c := []byte{}

	for i := 0; i < len(b); {
		// skip comment line
		if b[i] == byte('/') && b[i+1] == byte('*') {

			for {
				c = append(c, b[i])

				if b[i] == byte('\n') {
					i++
					break
				}

				i++
				continue
			}
		}

		if b[i] != byte('#') {
			c = append(c, b[i])
			i++
			continue
		}

		// found color hex notation

		j := i + 1

		// loop until we find semicolon

		for {
			if b[j] == byte(';') {

				h := b[i+1 : j]

				c = append(c, []byte(conv(h))...)

				break
			}

			j++
		}

		i = j
	}

	fmt.Println(string(c))

	os.WriteFile("./color-convert/advanced_expected__.css", c, 0o644)
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

func conv(c []byte) string {
	c = []byte(strings.ToLower(string(c)))

	for i := 0; i < len(c); i++ {
		v, ok := m[c[i]]
		if !ok {
			panic(fmt.Sprintf("unexpected %v", c[i]))
		}

		c[i] = v
	}

	switch len(c) {
	case 8:
		return fmt.Sprintf("rgba(%v %v %v / %.5f)", (c[0]<<4)|c[1], (c[2]<<4)|c[3], (c[4]<<4)|c[5], float64((c[6]<<4)|c[7])/0xff)
	case 6:
		return fmt.Sprintf("rgb(%v %v %v)", (c[0]<<4)|c[1], (c[2]<<4)|c[3], (c[4]<<4)|c[5])
	case 4:
		return fmt.Sprintf("rgba(%v %v %v / %.5f)", (c[0]<<4)|c[0], (c[1]<<4)|c[1], (c[2]<<4)|c[2], float64((c[3]<<4)|c[3])/0xff)
	case 3:
		return fmt.Sprintf("rgb(%v %v %v)", (c[0]<<4)|c[0], (c[1]<<4)|c[1], (c[2]<<4)|c[2])
	default:
		return "0, 0, 0"
	}
}
