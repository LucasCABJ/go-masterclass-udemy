package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckUsername(t *testing.T) {
	assert.True(t, CheckUsername("supername123"))
}

func TestCheckUsernameReturnsFalse(t *testing.T) {
	assert.False(t, CheckUsername("abc"))
}

// This is what we call table-driven test
func TestCheckUsernameTable(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{"too short", "abc", false},
		{"empty", "", false},
		{"contains 'admin'", "newadmin", false},
		{"valid username", "username_123", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CheckUsername(tc.input)
			assert.Equal(t, got, tc.want, "got:%v\twant:%v", got, tc.want)
		})
	}
}

func TestLoginTable(t *testing.T) {
	testCases := []struct {
		name      string
		username  string
		password  string
		shouldErr bool
		wantErr   error
	}{
		{"invalid username", "abc", "", true, ErrInvalidUsername},
		{"invalid password", "greatusername", "123456", true, ErrInvalidPassword},
		{"valid credentials", "greausername", "1234", false, nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ok, err := Login(tc.username, tc.password)
			if tc.shouldErr {
				assert.ErrorIs(t, err, tc.wantErr)
				assert.False(t, ok)
			} else {
				assert.NoError(t, err)
				assert.True(t, ok)
			}
		})
	}
}
