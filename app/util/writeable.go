package util

type Writeable interface {
	// ToByte convert corresponding struct to byte array
	ToByte() []byte //TODO change to return error
}
