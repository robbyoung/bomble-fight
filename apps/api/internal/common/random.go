package common

import (
	"math/rand"
	"time"
)

type IRandom interface {
	RandInt(int, int) int
	RandArrayEntry([]string) string
}

type Random struct {
	rand *rand.Rand
}

func NewRandom() *Random {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return &Random{
		rand: r,
	}
}

func (r *Random) RandInt(min int, max int) int {
	return r.rand.Intn(max-min) + min
}

func (r *Random) RandArrayEntry(arr []string) string {
	i := r.RandInt(0, len(arr))

	return arr[i]
}
