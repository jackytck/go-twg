package signal

import "testing"

func TestHandler(t *testing.T) {
	t.Log("Nat 1")
	// t.FailNow()
	t.Log("Nat 2")
	t.Error("Failed")
	t.Error("Failed2")
}
