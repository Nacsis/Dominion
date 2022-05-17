package app

func Uint8safe(a uint16) uint8 {
	b := uint8(a)
	if uint16(b) != a {
		panic("unsafe")
	}
	return b
}
