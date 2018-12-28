package quick

import (
	"testing"
	"testing/quick"
)

func TestSquareAndAdd(t *testing.T) {
	f := func(a, b int32) bool {
		got := SquareAndAdd(a, b)
		// t.Logf("SquareAndAdd(%d, %d) = %d\n", a, b, got)
		return got >= 0
	}
	cfg := quick.Config{
		MaxCount: 1000000,
	}
	if err := quick.Check(f, &cfg); err != nil {
		t.Error(err)
	}
}
