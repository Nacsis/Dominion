package app

import "testing"

func TestData(t *testing.T) {
	data := AppData{}
	NewCopper(&data)
	data.Encode(&Printer{})
}
