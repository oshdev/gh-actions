package main

import (
	"testing"
)

func TestAdding(t *testing.T) {
	if add(1, 2) != 3 {
		t.Errorf("big bada boom")
	}
}
