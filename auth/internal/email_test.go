package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {

	t.Run("Correct format email should return Email entity without error", func(t *testing.T) {
		// Given
		es := "apascualco@gmail.com"

		// When
		e, err := NewEmail(es)

		// Then
		assert.NoError(t, err)
		address := e.email
		assert.NotNil(t, address)
		assert.Equal(t, es, address.Address)
		assert.Equal(t, "", address.Name)
	})

	t.Run("Email without domain should return err", func(t *testing.T) {
		// Given
		es := "apascualco@.com"

		// When
		_, err := NewEmail(es)

		// Then
		assert.Error(t, err)
		assert.EqualError(t, err, "The email address is invalid")
	})

	t.Run("Email without username should return err", func(t *testing.T) {
		// Given
		es := "@gmail.com"

		// When
		_, err := NewEmail(es)

		// Then
		assert.Error(t, err)
		assert.EqualError(t, err, "The email address is invalid")
	})

	t.Run("Email without @ should return err", func(t *testing.T) {
		// Given
		es := "usernamegmail.com"

		// When
		_, err := NewEmail(es)

		// Then
		assert.Error(t, err)
		assert.EqualError(t, err, "The email address is invalid")
	})
}
