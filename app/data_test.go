package app

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestData(t *testing.T) {
	dominionApp := DominionApp{}

	fmt.Printf("Creating AppData\n")
	data := Data{}
	for i := 0; i < 3; i++ {
		NewCopper(&data)
	}

	fmt.Printf("Encode AppData\n")
	b := bytes.Buffer{}
	data.Encode(&b)

	fmt.Printf("Decode AppData\n")
	data2, _ := dominionApp.DecodeData(&b)

	fmt.Printf("Check Equal\n")
	assert.Equal(t, data.NumAllCards, data2.NumAllCards, "NumAllCards")
	assert.Equal(t, data.NextActor, data2.NextActor, "NextActor")
	for i := uint8(0); i < data.NumAllCards; i++ {
		assert.True(t, data.AllCards[i].equals(data2.AllCards[i]), "Card different")
	}

	fmt.Printf("Done\n")
}
