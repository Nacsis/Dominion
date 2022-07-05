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

func SliceToHashByte(slice []byte) [HashSizeByte]byte {
	var byteWithFixSize [HashSizeByte]byte
	copy(byteWithFixSize[:HashSizeByte], slice)
	return byteWithFixSize
}

func SliceToPreImageByte(slice []byte) [PreImageSizeByte]byte {
	var byteWithFixSize [PreImageSizeByte]byte
	copy(byteWithFixSize[:PreImageSizeByte], slice)
	return byteWithFixSize
}
