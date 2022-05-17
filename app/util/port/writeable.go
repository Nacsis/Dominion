package port

type Writeable interface {
	ToByte() []byte
}
