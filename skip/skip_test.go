package skip

import "testing"

func TestSkip(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	t.Log("This test ran!")
}
