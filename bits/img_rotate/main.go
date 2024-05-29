package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// BitmapFileHeader
// https://en.wikipedia.org/wiki/BMP_file_format
type BitmapFileHeader struct {
	// The following entries are possible:
	// BM: Windows 3.1x, 95, NT, ... etc.
	// BA: OS/2 struct bitmap array
	// CI: OS/2 struct color icon
	// CP: OS/2 const color pointer
	// IC: OS/2 struct icon
	// PT: OS/2 pointer
	Type               string
	Size               uint
	Reserved1          []byte
	Reserved2          []byte
	PixelArrayStartIdx uint
}

type CompressionMethod int

const (
	BI_RGB CompressionMethod = iota
	BI_RLE8
	BI_RLE4
	BI_BITFIELDS
	BI_JPEG
	BI_PNG
	BI_ALPHABITFIELDS
	_ = iota
	_ = iota
	_ = iota
	_ = iota
	BI_CMYK
	BI_CMYKRLE8
	BI_CMYKRLE4
)

// DIBHeader of BITMAPINFOHEADER type.
type DIBHeader struct {
	Width  uint
	Height uint
	// ColorDepth means bits per pixel!
	ColorDepth uint
	// Compression
	Compression CompressionMethod
	ImageSize   uint
}

func tole(b []byte) uint {
	var ans uint

	for i := len(b) - 1; i >= 0; i-- {
		ans |= uint(b[i]) // could have just sum it but lets use an OR :)

		if i != 0 {
			ans = ans << 8
		}
	}

	return ans
}

func main() {
	b, _ := os.ReadFile("image-rotate/teapot.bmp")

	bmph := BitmapFileHeader{
		Type:               string(b[:2]),
		Size:               tole(b[2:7]),
		Reserved1:          b[7:9],
		Reserved2:          b[8:10],
		PixelArrayStartIdx: tole(b[10:14]),
	}

	dibh := DIBHeader{
		Width:       tole(b[18:22]),
		Height:      tole(b[22:26]),
		ColorDepth:  tole(b[28:30]),
		Compression: CompressionMethod(tole(b[30:34])),
		ImageSize:   tole(b[34:38]),
	}

	fmt.Println("bmp header:")

	hb, _ := json.MarshalIndent(bmph, " ", "")
	fmt.Println(string(hb))

	fmt.Println("dib header:")

	hb, _ = json.MarshalIndent(dibh, " ", "")
	fmt.Println(string(hb))

	pxstart := bmph.PixelArrayStartIdx
	pxend := bmph.PixelArrayStartIdx + dibh.ImageSize -1

    for pxstart < pxend {
        b[pxstart], b[pxend] = b[pxend], b[pxstart]

        pxstart++
        pxend--
    }

	os.WriteFile("image-rotate/teapot_.bmp", b, 0o600)
}

