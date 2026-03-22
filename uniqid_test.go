package uniqid

import (
	"strings"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	id := New()
	if len(id) != 13 {
		t.Errorf("New() length = %d, want 13", len(id))
	}
}

func TestNewUniqueness(t *testing.T) {
	seen := make(map[string]struct{}, 10000)
	for i := 0; i < 10000; i++ {
		id := New()
		if _, ok := seen[id]; ok {
			t.Fatalf("duplicate ID found: %s at iteration %d", id, i)
		}
		seen[id] = struct{}{}
	}
}

func TestNewWithPrefix(t *testing.T) {
	id := NewWithPrefix("test")
	if !strings.HasPrefix(id, "test") {
		t.Errorf("NewWithPrefix() = %q, want prefix 'test'", id)
	}
	if len(id) != 17 {
		t.Errorf("NewWithPrefix() length = %d, want 17", len(id))
	}
}

func TestNewWithPrefixEmpty(t *testing.T) {
	id := NewWithPrefix("")
	if len(id) != 13 {
		t.Errorf("NewWithPrefix('') length = %d, want 13", len(id))
	}
}

func TestNewWithEntropy(t *testing.T) {
	id := NewWithEntropy()
	if len(id) != 23 {
		t.Errorf("NewWithEntropy() length = %d, want 23", len(id))
	}
}

func TestNewWithEntropyUniqueness(t *testing.T) {
	seen := make(map[string]struct{}, 10000)
	for i := 0; i < 10000; i++ {
		id := NewWithEntropy()
		if _, ok := seen[id]; ok {
			t.Fatalf("duplicate ID found: %s at iteration %d", id, i)
		}
		seen[id] = struct{}{}
	}
}

func TestNewWithPrefixAndEntropy(t *testing.T) {
	id := NewWithPrefixAndEntropy("pfx")
	if !strings.HasPrefix(id, "pfx") {
		t.Errorf("NewWithPrefixAndEntropy() = %q, want prefix 'pfx'", id)
	}
	if len(id) != 26 {
		t.Errorf("NewWithPrefixAndEntropy() length = %d, want 26", len(id))
	}
}

func TestConcurrency(t *testing.T) {
	const goroutines = 100
	const idsPerGoroutine = 100
	ch := make(chan string, goroutines*idsPerGoroutine)

	var wg sync.WaitGroup
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < idsPerGoroutine; j++ {
				ch <- New()
			}
		}()
	}
	wg.Wait()
	close(ch)

	seen := make(map[string]struct{}, goroutines*idsPerGoroutine)
	for id := range ch {
		if _, ok := seen[id]; ok {
			t.Fatalf("concurrent duplicate ID: %s", id)
		}
		seen[id] = struct{}{}
	}
}

func TestFormat(t *testing.T) {
	tests := []struct {
		number int64
		width  int
		want   string
	}{
		{255, 4, "00ff"},
		{255, 2, "ff"},
		{255, 1, "f"},
		{0, 3, "000"},
		{16, 2, "10"},
	}
	for _, tt := range tests {
		got := format(tt.number, tt.width)
		if got != tt.want {
			t.Errorf("format(%d, %d) = %q, want %q", tt.number, tt.width, got, tt.want)
		}
	}
}

func BenchmarkNew(b *testing.B) {
	for b.Loop() {
		New()
	}
}

func BenchmarkNewWithPrefix(b *testing.B) {
	for b.Loop() {
		NewWithPrefix("test")
	}
}

func BenchmarkNewWithEntropy(b *testing.B) {
	for b.Loop() {
		NewWithEntropy()
	}
}

func BenchmarkNewWithPrefixAndEntropy(b *testing.B) {
	for b.Loop() {
		NewWithPrefixAndEntropy("test")
	}
}
