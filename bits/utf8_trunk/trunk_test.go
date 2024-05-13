package bits_test

import (
    "fmt"
    "os"
    "testing"
    "bytes"
)

func TestTrunk(t *testing.T) {
    casesbytes, err := os.ReadFile("cases")
    if err != nil{ fmt.Println(err); t.Fail(); return}

    response, _ := os.Create("expected_")

    // prepare [][]byte
    cases := bytes.Split(casesbytes, []byte{0x0a})
    cases = cases[:len(cases)-1]

    for _, cas := range cases {
        size := int(cas[0])
        cas = cas[1:]

        if size >= len(cas) {
            response.Write(cas)
            response.Write([]byte{0x0a})

            continue
        }

        i := 0

        for {
            nums := f(cas[i])

            if i + nums > size {
                break
            }

            i += nums
        }

        response.Write(cas[:i])
        response.Write([]byte{0x0a})
    }

    response.Sync()
    response.Close()
}

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

