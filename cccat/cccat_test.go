package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestCccat(t *testing.T) {
	t.Run("reading from text file", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		want := "Your heart is the size of an ocean. Go find yourself in its hidden depths. - Rumi"
		ReadFile(buffer, "./snippet.txt")
		got := buffer.String()
		assertEqual(t, want, got)
	})
	t.Run("reading from standard input", func(t *testing.T) {
		want := "I Don'T Believe In Failure. It Is Not Failure If You Enjoyed The Process. - Oprah Winfrey"

		// simulates command <arg> | new_command <another_arg>
		// which passes output from command as input to new_command
		r, w, _ := os.Pipe()
		originalStdin := os.Stdin

		os.Stdin = r
		defer func() { os.Stdin = originalStdin }()
		fmt.Fprintln(w, want)
		w.Close()

		/*
			 * in order for this test to pass:
			 * 1. we must pass the text from "want" to standard input
				created from os.Pipe()
			 * 2. we pass a pointer of buffer to ReadStdIn
			 * 3. and that function will write to the buffer
			 * 4. so if ReadStdIn CANNOT read from os.Stdin,
				then this test fails
		*/
		buffer := &bytes.Buffer{}
		ReadStdIn(buffer)
		got := buffer.String()
		assertEqual(t, want, got)
	})
	t.Run("concat content of two files", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		want := "Your heart is the size of an ocean. Go find yourself in its hidden depths. - RumiI Don'T Believe In Failure. It Is Not Failure If You Enjoyed The Process. - Oprah Winfrey"
		ConcatFiles(buffer, "./snippet.txt", "snippet2.txt")
		got := buffer.String()
		assertEqual(t, want, got)
	})
}

func assertEqual(t testing.TB, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("got: %v, wanted: %v", got, want)
	}
}
