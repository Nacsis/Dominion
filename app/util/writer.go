package util

import (
	"io"
)

func Write(w io.Writer, o Writeable) error {
	_, err := w.Write(o.ToByte())
	return err
}

func AppendLength(in []byte) []byte {
	length := IntToByteArray16(len(in))
	return append(length[:], in...)
}
