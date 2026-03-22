package uniqid

import (
	"math/rand/v2"
	"strconv"
	"sync/atomic"
	"time"
)

var counter atomic.Int64

func init() {
	counter.Store(rand.Int64N(0x75bcd15))
}

// New returns a 13-character unique ID based on the current timestamp and an atomic counter.
func New() string {
	c := counter.Add(1)
	return format(time.Now().Unix(), 8) + format(c, 5)
}

// NewWithPrefix returns a unique ID prefixed with the given string.
func NewWithPrefix(prefix string) string {
	return prefix + New()
}

// NewWithEntropy returns a 23-character unique ID with additional random entropy.
func NewWithEntropy() string {
	c := counter.Add(1)
	id := format(time.Now().Unix(), 8) + format(c, 5)
	n := rand.Float64() * 10
	id += strconv.FormatFloat(n, 'f', 8, 64)
	return id
}

// NewWithPrefixAndEntropy returns a unique ID with a prefix and additional random entropy.
func NewWithPrefixAndEntropy(prefix string) string {
	return prefix + NewWithEntropy()
}

func format(number int64, width int) string {
	hex := strconv.FormatInt(number, 16)
	if width <= len(hex) {
		return hex[:width]
	}
	for len(hex) < width {
		hex = "0" + hex
	}
	return hex
}
