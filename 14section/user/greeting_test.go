package user

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Greeting(prefix, name string) {
	fmt.Printf("Hello, %s %s!\n", prefix, name)
}

func Greeting2(out io.Writer, prefix, name string) {
	fmt.Fprintf(out, "Hello, %s %s!\n", prefix, name)
}

func TestGreeting(t *testing.T) {
	originalStdout := os.Stdout

	r, w, err := os.Pipe()
	assert.NoError(t, err)

	os.Stdout = w
	defer func() {
		os.Stdout = originalStdout
	}()

	Greeting("Mr.", "Smith")
	w.Close()

	var buf strings.Builder
	_, err = io.Copy(&buf, r)
	assert.NoError(t, err)

	want := "Hello, Mr. Smith!\n"
	got := buf.String()
	assert.Equal(t, want, got)
}

func TestGreeting2(t *testing.T) {
	var buf strings.Builder

	Greeting2(&buf, "Mr.", "Smith")

	got := buf.String()
	want := "Hello, Mr. Smith!\n"
	assert.Equal(t, want, got)
}
