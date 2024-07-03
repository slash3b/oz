package main

import (
    "os"
    "fmt"
    "io/fs"
    "log"
    "encoding/binary"
)
/*
plan:

*/

func main() {
    dir := os.DirFS("varint")

    fs.WalkDir(dir, ".",func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

        if path == "." { return nil }

        b, err := os.ReadFile("varint/"+path)
        if err != nil {
            fmt.Println("err", path, err)

            panic(err)
        }

        fmt.Printf("content %08b\n", b)
        fmt.Printf("content %d\n", binary.BigEndian.Uint64(b))


		return nil
	})

}

