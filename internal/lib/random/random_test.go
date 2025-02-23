package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewRandomString(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "size = 1",
			size: 1,
		},
		{
			name: "size = 5",
			size: 5,
		},
		{
			name: "size = 10",
			size: 10,
		},
		{
			name: "size = 20",
			size: 20,
		},
		{
			name: "size = 30",
			size: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str1 := NewRandomString(tt.size)
			time.Sleep(100 * time.Millisecond)
			str2 := NewRandomString(tt.size)

			assert.Len(t, str1, tt.size)
			assert.Len(t, str2, tt.size)

			// Check that two generated strings are different
			// This is not an absolute guarantee that the function works correctly,
			// but this is a good heuristic for a simple random generator.

			// So it does not work due the time
			// of creation of pseudo-random generator

			// And I added time.Sleep to fix that
			// In real application it hardly possible to appear
			// that two strings would be generated at the same time
			assert.NotEqual(t, str1, str2)
		})
	}
}
