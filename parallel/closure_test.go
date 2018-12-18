package parallel

import (
	"fmt"
	"testing"
)

func TestClosure(t *testing.T) {
	testCases := []struct {
		arg  int
		want int
	}{
		{2, 4},
		{3, 9},
		{4, 16},
	}
	for _, tc := range testCases {
		tc := tc // copy value for parallel tests - do not delete
		t.Run(fmt.Sprintf("arg=%d", tc.arg), func(t *testing.T) {
			t.Parallel()
			t.Logf("Testing with: arg=%d, want=%d", tc.arg, tc.want)
			if tc.arg*tc.arg != tc.want {
				t.Errorf("%d^2 != %d", tc.arg, tc.want)
			}
		})
	}
}
