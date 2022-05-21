package util

type Writeable interface {
	ToByte() []byte
}
