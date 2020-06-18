package bitset

import (
	"errors"
	"math/bits"
)

const nBits = 64

type Bitset struct {
	bits []uint64
	size int
}

func New(size int) *Bitset {
	b := new(Bitset)
	n := 1 + (size-1)/nBits
	b.bits = make([]uint64, n)
	b.size = size
	return b
}

func (b *Bitset) All() bool {
	return b.Count() == b.size
}

func (b *Bitset) Any() bool {
	return b.Count() > 0
}

func (b *Bitset) Count() int {
	c := 0
	for _, x := range b.bits {
		c += bits.OnesCount64(x)
	}
	return c
}

func (b *Bitset) Flip() {
	for i := range b.bits {
		for j := 0; j < 64; j++ {
			if b.bits[i]&(1<<(j)) == 1 {
				b.bits[i] &^= 1 << (j)
			} else {
				b.bits[i] |= 1 << (j)
			}
		}
	}
}

func (b *Bitset) Reset() {
	for i := range b.bits {
		b.bits[i] = 0
	}
}

func (b *Bitset) Set(pos int, value bool) error {
	if pos >= b.size || pos < 0 {
		return errors.New("out of range")
	}
	idx := pos / nBits
	offset := pos % nBits
	if value {
		b.bits[idx] |= 1 << (offset)
	} else {
		b.bits[idx] &^= 1 << (offset)
	}

	return nil
}

func (b *Bitset) Test(pos int) (bool, error) {
	if pos >= b.size || pos < 0 {
		return false, errors.New("out of range")
	}
	idx := pos / nBits
	offset := pos % nBits

	return b.bits[idx]&(1<<(offset)) != 0, nil
}
