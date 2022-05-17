package util

import (
	"io"
	"perun.network/perun-examples/app-channel/app/util/port"
)

func WriteUInt8(w io.Writer, v uint8) error {
	_, err := w.Write([]byte{v})
	return err
}

func Write(w io.Writer, o port.Writeable) error {
	_, err := w.Write(o.ToByte())
	return err
}
