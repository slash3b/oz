package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type TS_FORMAT uint32

func (t TS_FORMAT) String() string {
	switch t {
	case TS_SEC_AND_MICRO:
		return "seconds and microseconds"
	case TS_SEC_AND_NANO:
		return "seconds and nanoseconds"
	}

	return ""
}

const (
	TS_SEC_AND_MICRO TS_FORMAT = iota
	TS_SEC_AND_NANO
)

type LinkLayerType uint32

const (
	LINKTYPE_NULL LinkLayerType = iota
	LINKTYPE_ETHERNET
	// ... many more types
)

// PcapHeader definition
// https://www.ietf.org/archive/id/draft-gharris-opsawg-pcap-01.html
// 24 octets
type FileHeader struct {
	Magic        uint32
	MajorVersion uint16
	MinorVersion uint16
	// TzOffset is always zero
	TzOffset uint32
	// TsAccurasy is a timestampe accuracy, always 0
	TsAccuracy     uint32
	SnapshotLength uint32

	LinkLayerHeaderType LinkLayerType
}

func NewFileHeader(src []byte) FileHeader {
	return FileHeader{
		Magic:               littleEndianToUint32(src[0:4]),
		MajorVersion:        littleEndianToUint16(src[4:6]),
		MinorVersion:        littleEndianToUint16(src[6:8]),
		TzOffset:            littleEndianToUint32(src[8:12]),
		TsAccuracy:          littleEndianToUint32(src[12:16]),
		SnapshotLength:      littleEndianToUint32(src[16:20]),
		LinkLayerHeaderType: LinkLayerType(littleEndianToUint32(src[20:24])),
	}
}

func (fh FileHeader) Format() TS_FORMAT {
	switch fh.Magic {
	case 0xA1B2C3D4:
		return TS_SEC_AND_MICRO
	case 0xA1B23C4D:
		return TS_SEC_AND_NANO
	default:
		panic("unknown timestamp format")
	}
}

func (fh FileHeader) Print() {
	fmt.Println("file header fields:")
	fmt.Printf("  magic: %x\n", fh.Magic)
	fmt.Printf("  pcap version: %d.%d\n", fh.MajorVersion, fh.MinorVersion)
	fmt.Printf("  tz offset: %v\n", fh.TzOffset)
	fmt.Printf("  ts accuracy: %v\n", fh.TsAccuracy)
	fmt.Printf("  snapshot lenght: %v\n", fh.SnapshotLength)
	fmt.Printf("  ll header type: %s\n", fh.LinkLayerHeaderType.String())
}

// PacketHeader lenght is 16 octets
type PacketHeader struct {
	Timestamp   uint32
	Tsmicronano uint32
	Len         int // int instead of uint32 for simplicity and ergonomics -- easier to manage in a loop, no need to convert
	UntruncLen  int
}

func NewPacketHeader(src []byte) PacketHeader {
	return PacketHeader{
		Timestamp:   littleEndianToUint32(src[0:4]),
		Tsmicronano: littleEndianToUint32(src[4:8]),
		Len:         int(littleEndianToUint32(src[8:12])),
		UntruncLen:  int(littleEndianToUint32(src[12:16])),
	}
}

func (pkh PacketHeader) Time() time.Time {
	return time.Unix(int64(pkh.Timestamp), 0)
}

func (pkh PacketHeader) Print() {
	fmt.Println("packet header:")
	fmt.Printf("  raw %#v\n", pkh)
	fmt.Printf("  len bytes %d\n", pkh.Len)
	fmt.Println("  len == untrunc", pkh.Len == pkh.UntruncLen)
	fmt.Println("  first packet header time", pkh.Time())
}

type IPv4Header struct {
	Version  int
	IHL      int
	TotalLen int
	// TTL is in seconds
	TTL      int
	Protocol InternetProtocol
	Src      string
	Dst      string
}

func NewIPv4Header(src []byte) IPv4Header {
	h := IPv4Header{
		Version:  int((src[0] & 0xf0) >> 4),
		IHL:      int((src[0] & 0x0f) << 2),
		TotalLen: int(littleEndianToUint16(src[2:4])),
		TTL:      int(src[8]),
		Protocol: InternetProtocol(src[9]),
	}

	h.Src = fmt.Sprintf("%d.%d.%d.%d", src[12], src[13], src[14], src[15])
	h.Dst = fmt.Sprintf("%d.%d.%d.%d", src[16], src[17], src[18], src[19])

	return h
}

func (iph IPv4Header) Print() {
	fmt.Println("ipv4 header")
	fmt.Println("  ip header | Version:", iph.Version)
	fmt.Println("  ip header | IHL:", iph.IHL)
	fmt.Println("  ip header | total length:", iph.TotalLen)
	fmt.Println("  ip header | ttl seconds:", iph.TTL)
	fmt.Println("  ip header | IP number:", iph.Protocol)
	fmt.Println("  ip header | source:", iph.Src)
	fmt.Println("  ip header | destination:", iph.Dst)
}

type InternetProtocol int

const (
	HOPOPT InternetProtocol = iota
	ICMP
	IGMP
	GGP
	IPinIP
	ST
	TCP
)

// todo: figure out what % of connections are ACKnowledged?
func main() {
	pcap := make([]byte, 0)

	for {
		b := make([]byte, 1<<10) // 1024

		n, err := os.Stdin.Read(b)

		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}

		pcap = append(pcap, b[0:n]...)
	}

	fmt.Printf("total read bytes in pcap file: %v \n\n", len(pcap))

	// pcap file header first
	fileHeader := NewFileHeader(pcap[:24])
	fileHeader.Print()

	// points to first unread byte in slice
	curr := 24

	c := 0

	for curr < len(pcap) {
		// per packet header
		ph := NewPacketHeader(pcap[curr : curr+16])
		curr += 16

		ph.Print()

		// protocol type description: https://www.tcpdump.org/linktypes/LINKTYPE_NULL.html
		protocolType := littleEndianToUint32(pcap[curr : curr+4])
		curr += 4

		if protocolType != 2 {
			panic(fmt.Sprintf("unexpected protocol type %d, expected 2 (IPv4)", protocolType))
		}

        // reslice to be used in this iteration
        s := pcap[curr : curr+ph.Len]
		curr += ph.Len-4 // minus 4 bytes we processed above, this is messy.

		// IPv4 Header
        ipheader := NewIPv4Header(s[:20])

        // 

        // payload
        

		ipheader.Print()

        tcppacket := pcap[curr + 24 : curr+ph.Len]
        fmt.Printf("%#v\n", tcppacket)

        fmt.Println(littleEndianToUint16(tcppacket[:2]))
        fmt.Println(littleEndianToUint16(tcppacket[2:4]))

        break


		c++
	}

	fmt.Println("total", c)
}

func littleEndianToUint32(src []byte) uint32 {
	var res uint32

	for i := 0; i < len(src); i++ {
		res += uint32(src[i]) << (8 * i)
	}

	return res
}

func littleEndianToUint16(src []byte) uint16 {
	var res uint16

	for i := 0; i < len(src); i++ {
		res += uint16(src[i]) << (8 * i)
	}

	return res
}

func bigEndianToUint16(src []byte) uint16 {
	var res uint16

	for i := len(src)-1; i >= 0; i-- {
		res += uint16(src[i]) << (8 * i)
	}

	return res
}
