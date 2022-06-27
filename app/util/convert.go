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

func SliceToHashByte(slice []byte) [HashSize]byte {
	var byteWithFixSize [HashSize]byte
	copy(byteWithFixSize[:HashSize], slice)
	return byteWithFixSize
}

func SliceToPreImageByte(slice []byte) [PreImageSize]byte {
	var byteWithFixSize [PreImageSize]byte
	copy(byteWithFixSize[:PreImageSize], slice)
	return byteWithFixSize
}
