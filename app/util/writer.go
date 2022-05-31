package util

import (
	"io"
)

func Write(w io.Writer, o Writeable) error {
	_, err := w.Write(o.ToByte())
	return err
}
