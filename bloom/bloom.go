package bloom

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

const iterations = 3
const size = 3
const cap = 257257257

// filter describes the interface for the bloom filter.
type filter interface {
	Add(string) error
	Check(string) bool
}

// Filter implements the bloom filter.
type Filter struct {
	bitmap [cap]bool
}

// Add inserts new strings into the filter.
func (f *Filter) Add(str string) error {
	hashed := hash(str)

	for i := 0; i < iterations; i++ {
		f.bitmap[hashToIndex(hashed, i)] = true
	}
	return nil
}

// Check returns if the string is in the filter.
func (f Filter) Check(str string) bool {
	hashed := hash(str)

	for i := 0; i < iterations; i++ {
		if !f.bitmap[hashToIndex(hashed, i)] {
			return false
		}
	}

	return true
}

func hashToIndex(hash [16]byte, i int) int {
	index := ""
	for ii := 0; ii < size; ii++ {
		index += fmt.Sprintf("%03d", hash[i*size+ii])
	}
	computed, _ := strconv.Atoi(index)
	return computed
}

func hash(str string) [16]byte {
	return md5.Sum([]byte(str))
}
