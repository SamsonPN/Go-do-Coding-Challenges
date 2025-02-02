package main

import (
	"bytes"
	"testing"
)

func TestCccat(t *testing.T) {
	t.Run("", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		wanted := "Your heart is the size of an ocean. Go find yourself in its hidden depths. - Rumi"
		ReadFile(buffer, "./snippet.txt")
		got := buffer.String()

		if wanted != got {
			t.Errorf("got %v, wanted %v", got, wanted)
		}
	})
}
