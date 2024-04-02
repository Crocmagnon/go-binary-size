package a

import (
	"testing"

	"github.com/loopfz/gadgeto/tonic"
)

func TestSomethingInternal(t *testing.T) {
	if tonic.DefaultBindingHook == nil {
		t.Error("uh oh")
	}
}
