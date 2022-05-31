package util

// BoolToByte convert bool to bytes
func BoolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

// ByteToBool convert byte to bool
func ByteToBool(b byte) bool {
	return b == 1
}
