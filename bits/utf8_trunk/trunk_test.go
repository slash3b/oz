package bits_test

import (
    "fmt"
    "os"
    "testing"
    "bytes"
)

func TestTrunk(t *testing.T) {
    cases, err := os.ReadFile("cases")
    if err != nil{ fmt.Println(err); t.Fail(); return}

    expected, err := os.ReadFile("expected")
    if err != nil{ fmt.Println(err); t.Fail(); return}

    // strip last new line byte
    cases = cases[:len(cases)-1]
    expected = expected[:len(cases)-1]
    expected_ := bytes.Split(expected, []byte{0x0a})

    debug, _ := os.Create("debug")
    defer debug.Sync(); defer debug.Close()

    for lnum, line := range bytes.Split(cases, []byte{0x0a}) {
        size := int(line[0])
        line = line[1:]

        if size >= len(line) {
            debug.Write(line); debug.Write([]byte{0x0a})
            if !bytes.Equal(line, expected_[lnum]) { t.Fail(); return }

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

        debug.Write(line[:i]); debug.Write([]byte{0x0a})
        if !bytes.Equal(line[:i], expected_[lnum]) { t.Fail(); return }
    }

}

// f tells you how many bytes to follow
func f(b byte) int {
    switch b >> 4 {
	case 0x0f:
		return 4
	case 0x0e:
		return 3
	}

	if (b >> 5) == 0x06 {
		return 2
	}

	return 1
}

