package main

import (
	"encoding/binary"
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
	Size               int
	Reserved1          []byte
	Reserved2          []byte
	PixelArrayStartIdx int
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
	Width  int
	Height int
	// ColorDepth means bits per pixel
	ColorDepth int
	// Compression
	Compression CompressionMethod
}

func main() {
	b, _ := os.ReadFile("image-rotate/teapot.bmp")

	bmpHeader := BitmapFileHeader{
		Type:               string(b[:2]),
		Size:               int(binary.LittleEndian.Uint32(b[2:7])),
		Reserved1:          b[7:9],
		Reserved2:          b[8:10],
		PixelArrayStartIdx: int(binary.LittleEndian.Uint16(b[10:13])),
	}

	dibHeader := DIBHeader{
		Width:      int(binary.LittleEndian.Uint16(b[18:22])),
		Height:     int(binary.LittleEndian.Uint16(b[22:26])),
		ColorDepth: int(binary.LittleEndian.Uint16(b[28:30])),
		Compression: CompressionMethod(binary.LittleEndian.Uint32(b[30:34])),
	}

    hb, _ := json.MarshalIndent(bmpHeader, " ", "")
	fmt.Println(string(hb))

    hb, _ = json.MarshalIndent(dibHeader, " ", "")
	fmt.Println(string(hb))
}
