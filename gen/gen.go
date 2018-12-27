package gen

import (
	"fmt"
	"math/rand"
	"sync"
)

var m sync.Mutex
var count = 0

var firstNames = []string{
	"Nat",
	"Bob",
	"Mark",
	"Jinny",
	"Holly",
	"Alice",
	"Eve",
}

// Email generates unique email address every time it is called.
func Email() string {
	m.Lock()
	defer m.Unlock()
	name := firstNames[rand.Intn(len(firstNames))]
	ret := fmt.Sprintf("%s-%d@example.com", name, count)
	count++
	return ret
}
