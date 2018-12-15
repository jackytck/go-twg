package parallel

import (
	"testing"
	"time"
)

func TestA(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second * 2)
}

func TestB(t *testing.T) {
	// t.Parallel()
	t.Run("Sub1", func(tt *testing.T) {
		tt.Parallel()
		time.Sleep(time.Second * 1)
	})

	t.Run("Sub2", func(tt *testing.T) {
		tt.Parallel()
		time.Sleep(time.Second * 1)
	})
}

func TestC(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second * 2)
}
