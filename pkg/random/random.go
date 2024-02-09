package random

import (
	"math/rand"
	"time"
)

// Choice returns a random element from provided array.
func Choice(arr []string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	idx := r.Intn(len(arr))
	return arr[idx]
}
