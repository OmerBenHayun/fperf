package generator

import (
	"math/rand"
	"time"
)

type keySpace struct {
	r      *rand.Rand
	nbytes int
}

func newKeySpace(nbytes int) *keySpace {
	return &keySpace{
		r:      rand.New(rand.NewSource(time.Now().Unix())),
		nbytes: nbytes,
	}
}

func (ks *keySpace) randKey() string {
	p := make([]byte, ks.nbytes)
	ks.r.Read(p)
	return string(p)
}
func (ks *keySpace) randRange() (string, string) {
	start := []byte(ks.randKey())
	if len(start) == 0 {
		return "", ""
	}
	// the max range is 256 according to the last bytes of start
	end := make([]byte, len(start))
	copy(end, start)
	end[len(end)-1] = 0xFF
	return string(start), string(end)
}
