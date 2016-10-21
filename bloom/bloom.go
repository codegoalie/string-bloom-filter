package bloom

import "crypto/md5"

const iterations = 3

// filter describes the interface for the bloom filter.
type filter interface {
	Add(string) error
	Check(string) bool
}

// Filter implements the bloom filter.
type Filter struct {
	bitmap [256]bool
}

// Add inserts new strings into the filter.
func (f *Filter) Add(new string) error {
	hash := md5.Sum([]byte(new))

	for i := 0; i < iterations; i++ {
		f.bitmap[hash[i]] = true
	}
	return nil
}

// Check returns if the string is in the filter.
func (f Filter) Check(str string) bool {
	hash := md5.Sum([]byte(str))

	for i := 0; i < iterations; i++ {
		if !f.bitmap[hash[i]] {
			return false
		}
	}

	return true
}
