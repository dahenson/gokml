package gokml

import (
	"testing"
)

func TestBoolToInt(t *testing.T) {
	i := boolToInt(true)
	if i != 1 {
		t.Fatalf("[true] returned %d", i)
	}

	i = boolToInt(false)
	if i != 0 {
		t.Fatalf("[false returned %d", i)
	}
}
