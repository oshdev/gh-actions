package adding_test

import (
	"github.com/oshdev/gh-actions/pkg/adding"
	"testing"
)

func TestAdding(t *testing.T) {
	if adding.Add(1, 2) != 3 {
		t.Errorf("big bada boom")
	}
}
