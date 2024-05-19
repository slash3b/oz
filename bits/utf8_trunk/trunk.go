package main

import (
    "fmt"
    "os"
    "bytes"
)

func main() {
    b, _ := os.ReadFile("cases")
    cases := bytes.Split(b, []byte{0x0a})
    cases = cases[:len(cases)-1] // strip last new line byte

    for _, line := range cases {
        size := int(line[0])
        line = line[1:]

        if size >= len(line) {
            fmt.Fprintln(os.Stdout, string(line))

            continue
        }

        // loop over line bytes to correctly truncate.

        i := 0

        for {
            nums := f(line[i])

            if i + nums > size {
                break
            }

            i += nums
        }

        fmt.Fprintln(os.Stdout, string(line[:i]))
    }
}

// f tells you how many bytes to follow
func f(b byte) int {
    switch b >> 4 {
	case 0xf:
		return 4
	case 0x3:
		return 3
	}

	if (b >> 5) == 0x6 {
		return 2
	}

	return 1
}

