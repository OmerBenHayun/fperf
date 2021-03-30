package generator

import (
	"math/rand"
	"time"
)

type KeySpace struct {
	r      *rand.Rand
	nbytes int
}

func NewKeySpace(nbytes int) *KeySpace {
	return &KeySpace{
		r:      rand.New(rand.NewSource(time.Now().Unix())),
		nbytes: nbytes,
	}
}

func (ks *KeySpace) RandKey() string {
	p := make([]byte, ks.nbytes)
	ks.r.Read(p)
	return string(p)
}
func (ks *KeySpace) RandRange() (string, string) {
	start := []byte(ks.RandKey())
	if len(start) == 0 {
		return "", ""
	}
	// the max range is 256 according to the last bytes of start
	end := make([]byte, len(start))
	copy(end, start)
	end[len(end)-1] = 0xFF
	return string(start), string(end)
}
