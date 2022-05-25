package util

type Readable interface {
	// Of create corresponding struct out of bytes
	Of(dataBytes []byte) //TODO change to return error
}
