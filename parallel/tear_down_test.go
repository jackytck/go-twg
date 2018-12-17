package parallel

import (
	"fmt"
	"testing"
	"time"
)

func TestTeardown(t *testing.T) {
	fmt.Println("Setup...")
	defer fmt.Println("Deferred teardown")

	t.Run("Group", func(tt *testing.T) {
		tt.Run("SubA", func(tt *testing.T) {
			tt.Parallel()
			time.Sleep(time.Second)
			fmt.Println("SubA done")
		})

		tt.Run("SubB", func(tt *testing.T) {
			tt.Parallel()
			time.Sleep(time.Second)
			fmt.Println("SubB done")
		})
	})

	fmt.Println("Teardown")
}
