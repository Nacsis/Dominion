package app

type Writeable interface {
	ToByte() []byte
}
