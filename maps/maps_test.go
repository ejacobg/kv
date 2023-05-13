package maps

import (
	"github.com/ejacobg/kv/kvtest"
	"testing"
)

func TestAcceptance(t *testing.T) {
	kvtest.TestStore(t, NewStore())
}
