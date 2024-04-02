package a_test

import (
	"testing"

	"github.com/loopfz/gadgeto/tonic"
)

func TestSomethingExternal(t *testing.T) {
	if tonic.DefaultBindingHook == nil {
		t.Error("uh oh")
	}
}
