// Code generated by "stringer -type=LinkLayerType"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LINKTYPE_NULL-0]
	_ = x[LINKTYPE_ETHERNET-1]
}

const _LinkLayerType_name = "LINKTYPE_NULLLINKTYPE_ETHERNET"

var _LinkLayerType_index = [...]uint8{0, 13, 30}

func (i LinkLayerType) String() string {
	if i >= LinkLayerType(len(_LinkLayerType_index)-1) {
		return "LinkLayerType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LinkLayerType_name[_LinkLayerType_index[i]:_LinkLayerType_index[i+1]]
}
