package main

import (
	"testing"
)

func testAdding(t *testing.T) {
	if add(1, 2) != 3 {
		t.Errorf("big bada boom")
	}
}
