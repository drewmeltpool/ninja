package example_bin

import "testing"

func TestHundred(t *testing.T) {
	answer := Ninja()

	if answer != "ninja" {
		t.Errorf("Expected ninja but got %s", answer)
	}
}
