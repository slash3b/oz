package main

import (
	"os"
	"syscall"
	"time"

	"golang.org/x/term"
)

/*
notes:
tty
stty
*/

func main() {
	b := make([]byte, 1)

	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	for {
		os.Stdin.Read(b)

		input := int(b[0])

		switch syscall.Signal(input) {
		case syscall.SIGINT, syscall.SIGQUIT:
			return
		}

		bps := input - 0x30 // ascii offset.
		if bps > 9 {
			continue
		}

		for i := 0; i < bps; i++ {
			os.Stdout.Write([]byte{0x07})

			// I guessed that system needs a delay to output
			// a bleep sound, so I've added it. 300ms seems like enough.
			// not sure why delay is needed.
			time.Sleep(time.Millisecond * 200)
		}
	}
}
