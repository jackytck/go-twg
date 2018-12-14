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

}

func TestC(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second * 2)
}
