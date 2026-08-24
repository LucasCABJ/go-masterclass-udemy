package user

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestValidateUsername(t *testing.T) {
// 	got := CheckUsername("lucascabj")
// 	if got != true {
// 		t.Errorf("got: %v \t want: %v", got, true)
// 	}
// }

// func TestCheckUsernameTable(t *testing.T) {
// 	testCases := []struct {
// 		name  string
// 		input string
// 		want  bool
// 	}{
// 		{name: "valid username", input: "lucascabj", want: true},
// 		{name: "invalid username - empty string", input: "", want: false},
// 		{name: "invalid username - too short", input: "lucas", want: false},
// 		{name: "invalid username - contains admin", input: "adminlucas", want: false},
// 	}
// 	for _, tc := range testCases {
// 		got := CheckUsername(tc.input)
// 		if got != tc.want {
// 			t.Errorf("%s: got: %v \t want: %v", tc.name, got, tc.want)
// 		}
// 	}
// }

func TestCheckUsernameWithSubtest(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "valid username", input: "lucascabj", want: true},
		{name: "invalid username - empty string", input: "", want: false},
		{name: "invalid username - too short", input: "lucas", want: false},
		{name: "invalid username - contains admin", input: "adminlucas", want: false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CheckUsername(tc.input)
			message := fmt.Sprintf("%s: got: %v \t want: %v", tc.name, got, tc.want)
			assert.Equal(t, tc.want, got, message)
		})
	}
}
