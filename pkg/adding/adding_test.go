package adding_test

import (
	"testing"

	"github.com/oshdev/gh-actions/pkg/adding"
)

func TestAdding(t *testing.T) {
	if adding.Add(1, 2) != 3 {
		t.Errorf("big bada boom")
	}
}
