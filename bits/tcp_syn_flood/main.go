package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type TS_FORMAT uint

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

type LinkLayerType uint

const (
	LINKTYPE_NULL LinkLayerType = iota
	LINKTYPE_ETHERNET
	// ... many more types
)

// PcapHeader definition
// https://www.ietf.org/archive/id/draft-gharris-opsawg-pcap-01.html
// 24 octets
type FileHeader struct {
	Magic        uint
	MajorVersion uint
	MinorVersion uint
	// TzOffset is always zero
	TzOffset uint
	// TsAccurasy is a timestampe accuracy, always 0
	TsAccuracy     uint
	SnapshotLength uint

	LinkLayerHeaderType LinkLayerType
}

func NewFileHeader(src []byte) FileHeader {
	return FileHeader{
		Magic:               leToUint(src[0:4]),
		MajorVersion:        leToUint(src[4:6]),
		MinorVersion:        leToUint(src[6:8]),
		TzOffset:            leToUint(src[8:12]),
		TsAccuracy:          leToUint(src[12:16]),
		SnapshotLength:      leToUint(src[16:20]),
		LinkLayerHeaderType: LinkLayerType(leToUint(src[20:24])),
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
	Timestamp   uint
	Tsmicronano uint
	Len         int
	UntruncLen  int
}

func NewPacketHeader(src []byte) PacketHeader {
	return PacketHeader{
		Timestamp:   leToUint(src[0:4]),
		Tsmicronano: leToUint(src[4:8]),
		Len:         int(leToUint(src[8:12])),
		UntruncLen:  int(leToUint(src[12:16])),
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
		TotalLen: int(leToUint(src[2:4])),
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

type TCPSegmentHeader struct {
	Srcport      uint
	Dstport      uint
	Seqnumber    uint
	Acknumber    uint
	Offsetnumber uint
	Flags        uint
}

func NewTCPSegmentHeader(src []byte) TCPSegmentHeader {
	return TCPSegmentHeader{
		Srcport:      beToUint(src[:2]),
		Dstport:      beToUint(src[2:4]),
		Seqnumber:    beToUint(src[4:8]),
		Acknumber:    beToUint(src[8:12]),
		Offsetnumber: uint(src[12]) >> 4,
		Flags:        uint(src[13]),
	}
}

func (segh TCPSegmentHeader) Print() {
	fmt.Println("TCP Segement Header:")
	fmt.Printf("  segment   | source port: %d\n", segh.Srcport)
	fmt.Printf("  segment   | destination port: %d\n", segh.Dstport)
	fmt.Printf("  segment   | sequence number: %d\n", segh.Seqnumber)
	fmt.Printf("  segment   | acknowledgement number: %d\n", segh.Seqnumber)
	fmt.Printf("  segment   | offset number: %d\n", segh.Offsetnumber)
	fmt.Printf("  segment   | flags set: %v\n", ListSetFlags(segh.Flags))
}

type TCPFlag uint

var AllTCPFlags = []TCPFlag{CWR, ECE, URG, ACK, PSH, RST, SYN, FIN}

const (
	CWR TCPFlag = 1 << 7
	ECE TCPFlag = 1 << 6
	URG TCPFlag = 1 << 5
	ACK TCPFlag = 1 << 4
	PSH TCPFlag = 1 << 3
	RST TCPFlag = 1 << 2
	SYN TCPFlag = 1 << 1
	FIN TCPFlag = 1 << 0
)

func (tcpf TCPFlag) Is(i uint) bool {
	return i&uint(tcpf) > 0
}

func ListSetFlags(i uint) []string {
	res := make([]string, 0, len(AllTCPFlags))

	for _, flag := range AllTCPFlags {
		if flag.Is(i) {
			res = append(res, flag.String())
		}
	}

	return res
}

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

	pcap = pcap[24:]

	counter := 0

	initiated := 0
	acked := 0

	for len(pcap) > 0 {
		// per packet header
		ph := NewPacketHeader(pcap[:16])
		pcap = pcap[16:]
		if counter == 0 {
			ph.Print()
		}

		// protocol type description: https://www.tcpdump.org/linktypes/LINKTYPE_NULL.html
		protocolType := leToUint(pcap[:4])
		pcap = pcap[4:]
		ph.Len = ph.Len - 4 // minus 4 bytes we processed above, this is messy.

		if protocolType != 2 {
			panic(fmt.Sprintf("unexpected protocol type %d, expected 2 (IPv4)", protocolType))
		}

		// IPv4 Header
		ipheader := NewIPv4Header(pcap[:20]) // hardcoded to 20, no options, bad
		if counter == 0 {
			ipheader.Print()
		}

		// tcp packet
		tcp := NewTCPSegmentHeader(pcap[20:ph.Len])
		if counter == 0 {
			tcp.Print()
		}

		fmt.Printf("src: %d -> dst: %d  %v\n", tcp.Srcport, tcp.Dstport, ListSetFlags(tcp.Flags))
		if tcp.Dstport == 80 && SYN.Is(tcp.Flags) {
			initiated++
		}

		if tcp.Srcport == 80 && ACK.Is(tcp.Flags) {
			acked++
		}

		// finally truncate packet header length
		pcap = pcap[ph.Len:]

		counter++
	}

	fmt.Println("----------------------------------")
	fmt.Println("----------------------------------")
	fmt.Println("total packets parsed", counter)
	fmt.Println("total initiated connections", initiated)
	fmt.Println("total acknowledged connections", acked)
	fmt.Printf("total acknowledged connections %.2f%%\n", float64(acked)/float64(initiated)*100)

}

// could have used package encoging/binary
// binary.BigEndian and binary.LittleEndian
// but I have to write it myself.

func beToUint(src []byte) uint {
	var res uint

	for i, j := len(src)-1, 0; i >= 0; i, j = i-1, j+1 {
		res |= uint(src[i]) << (8 * j)
	}

	return res
}

func leToUint(src []byte) uint {
	var res uint

	for i := 0; i < len(src); i++ {
		res |= uint(src[i]) << (8 * i)
	}

	return res
}
