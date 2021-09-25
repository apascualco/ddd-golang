package auth

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {

	t.Run("Given valid values should create new user entity", func(t *testing.T) {
		// Given
		idString := "d85ecb79-403c-443a-99cb-9ae3916a4961"
		id, err := uuid.Parse(idString)
		assert.NoError(t, err)
		uuid := NewUUID(id)
		e := "apascualco@gmail.com"
		email, err := NewEmail(e)
		assert.NoError(t, err)
		p := "1234meomfeofm"
		password, err := NewPassword(p)
		assert.NoError(t, err)

		// When
		user, err := NewUser(uuid, email, password)

		// Then
		assert.NoError(t, err)
		assert.Equal(t, idString, user.ID().String())
		assert.Equal(t, e, user.Email())
		assert.True(t, user.ValidatePassword(p))
	})
}
