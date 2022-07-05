package util

import "unsafe"

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

func Uint32Safe(a int) uint32 {
	b := uint32(a)
	if int(b) != a {
		panic("unsafe")
	}
	return b
}

func Uint32ToByteArray(val uint32) [4]byte {
	return *(*[4]byte)(unsafe.Pointer(&val))
}

func IntToByteArray32(val int) [4]byte {
	return Uint32ToByteArray(Uint32Safe(val))
}

func Uint16Safe(a int) uint16 {
	b := uint16(a)
	if int(b) != a {
		panic("unsafe")
	}
	return b
}

func Uint16ToByteArray(val uint16) [2]byte {
	return *(*[2]byte)(unsafe.Pointer(&val))
}

func IntToByteArray16(val int) [2]byte {
	return Uint16ToByteArray(Uint16Safe(val))
}

func ByteArrayToUint32(val [4]byte) uint32 {
	return *(*uint32)(unsafe.Pointer(&val))
}

func ByteArrayToUint16(val [2]byte) uint16 {
	return *(*uint16)(unsafe.Pointer(&val))
}
