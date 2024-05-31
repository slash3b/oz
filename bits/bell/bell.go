package main

import (
	"os"
	"time"
)

func main() {
	b := make([]byte, 1, 1)

	for {
		os.Stdin.Read(b)

		n := int(b[0]) - 0x30

		for i := 0; i < n; i++ {
			os.Stderr.Write([]byte{0x07})

			// I guessed that system needs a delay to output
			// a bleep sound, so I've added it. 300ms seems like enough.
            // not sure why delay is needed.
			time.Sleep(time.Millisecond * 300)
		}
	}
}
