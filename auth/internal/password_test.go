package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {

	t.Run("Given correct password should return password entity without error", func(t *testing.T) {
		// Given
		up := "EFmofmeomddo"

		// When
		p, err := NewPassword(up)

		// Then
		assert.NoError(t, err)
		assert.Equal(t, up, p.password)
	})

	t.Run("Given correct hash should return password entity without error", func(t *testing.T) {
		// Given
		up := "EFmofmeofmfdmmddo"

		// When
		p, err := NewHashPassword("$2a$14$lYyZ5gLotq/NQjxovauxf.WCYcT49nmpwJM33b9lfgpsDOdzWV4Nq")

		// Then
		assert.NoError(t, err)
		assert.True(t, p.ValidatePassowrd(up))
	})

	t.Run("Correct password should return password entity without error and hashing should be valid", func(t *testing.T) {
		// Given
		up := "EFmofmeomddo"

		// When
		p, err := NewPassword(up)
		assert.NoError(t, err)

		// Then
		assert.Equal(t, up, p.password)
		assert.True(t, p.ValidatePassowrd(up))
	})
}
