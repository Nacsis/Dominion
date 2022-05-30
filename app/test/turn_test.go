package test

import (
	"github.com/stretchr/testify/assert"
	"perun.network/perun-examples/app-channel/app"
	"testing"
)

// Test_Turn_Serialization
func Test_Turn_Serialization(t *testing.T) {
	turn := app.Turn{}

	turn.Init(0)

	bytes := turn.ToByte()

	turnActual := app.Turn{}
	turnActual.Of(bytes[1:])
	assert.Equal(t, turn, turnActual)
}
