package auth

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUUID(t *testing.T) {
	t.Run("Create a valid uuid entity from string", func(t *testing.T) {
		// Given
		id := "c315c2d8-0df8-475b-9ede-8cd352534c38"

		// When
		ud, err := NewUUIDByString(id)

		// Then
		assert.NoError(t, err)
		assert.Equal(t, ud.String(), id)
	})

	t.Run("Invalid uuid string should return error when try to create entity UUID from string", func(t *testing.T) {
		// Given
		id := "invalid"

		// When
		_, err := NewUUIDByString(id)

		// Then
		assert.Error(t, err)
		assert.EqualError(t, err, "Check the uuid format")
	})

	t.Run("Given a valid uuid from byte should create UUID instance", func(t *testing.T) {
		// Given
		u := uuid.New()
		id := u.String()
		b, err := u.MarshalBinary()
		assert.NoError(t, err)

		// When
		uuid, err := NewUUIDByByte(b)

		// Then
		assert.NoError(t, err)
		assert.Equal(t, uuid.String(), id)
	})

	t.Run("Invalid uuid bytes should return error when try to create entity UUID from string", func(t *testing.T) {
		// Given
		id := []byte("invalid")

		// When
		_, err := NewUUIDByByte(id)

		// Then
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid UUID (got 7 bytes)")
	})

	t.Run("Given a valid uuid should match", func(t *testing.T) {
		// Given
		id := "c315c2d8-0df8-475b-9ede-8cd352534c38"
		u, err := uuid.Parse(id)
		assert.NoError(t, err)

		// When
		uud := NewUUID(u)

		// Then
		assert.Equal(t, uud.String(), id)
	})
}
