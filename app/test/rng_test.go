package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"perun.network/perun-examples/dominion-cli/app"
)

// Test_Rng_Serialization
func Test_Rng_Serialization(t *testing.T) {
	committed := rngCommittedSetUp()
	touched := rngTouchedSetUp()
	released := rngReleaseSetUp()

	bytesCommitted := committed.ToByte()
	bytesTouched := touched.ToByte()
	bytesReleased := released.ToByte()

	committedActual := app.RNG{}
	touchedActual := app.RNG{}
	releasedActual := app.RNG{}
	committedActual.Of(bytesCommitted[1:])
	touchedActual.Of(bytesTouched[1:])
	releasedActual.Of(bytesReleased[1:])

	assert.Equal(t, committed, committedActual)
	assert.Equal(t, touched, touchedActual)
	assert.Equal(t, released, releasedActual)
}
