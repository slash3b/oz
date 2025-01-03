// Code generated by "stringer -type=TCPFlag"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CWR-128]
	_ = x[ECE-64]
	_ = x[URG-32]
	_ = x[ACK-16]
	_ = x[PSH-8]
	_ = x[RST-4]
	_ = x[SYN-2]
	_ = x[FIN-1]
}

const (
	_TCPFlag_name_0 = "FINSYN"
	_TCPFlag_name_1 = "RST"
	_TCPFlag_name_2 = "PSH"
	_TCPFlag_name_3 = "ACK"
	_TCPFlag_name_4 = "URG"
	_TCPFlag_name_5 = "ECE"
	_TCPFlag_name_6 = "CWR"
)

var (
	_TCPFlag_index_0 = [...]uint8{0, 3, 6}
)

func (i TCPFlag) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _TCPFlag_name_0[_TCPFlag_index_0[i]:_TCPFlag_index_0[i+1]]
	case i == 4:
		return _TCPFlag_name_1
	case i == 8:
		return _TCPFlag_name_2
	case i == 16:
		return _TCPFlag_name_3
	case i == 32:
		return _TCPFlag_name_4
	case i == 64:
		return _TCPFlag_name_5
	case i == 128:
		return _TCPFlag_name_6
	default:
		return "TCPFlag(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
