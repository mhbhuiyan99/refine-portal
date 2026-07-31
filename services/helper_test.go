package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunkStrings(t *testing.T) {
	tests := []struct {
		name     string
		items    []string
		size     int
		expected [][]string
	}{
		{
			name:  "normal case",
			items: []string{"A", "B", "C", "D", "E"},
			size:  2,
			expected: [][]string{
				{"A", "B"},
				{"C", "D"},
				{"E"},
			},
		},
		{
			name:     "empty slice",
			items:    []string{},
			size:     2,
			expected: nil,
		},
		{
			name:  "size is one",
			items: []string{"A", "B", "C"},
			size:  1,
			expected: [][]string{
				{"A"},
				{"B"},
				{"C"},
			},
		},
		{
			name:  "size larger than slice",
			items: []string{"A", "B", "C"},
			size:  10,
			expected: [][]string{
				{"A", "B", "C"},
			},
		},
		{
			name:  "size equals slice length",
			items: []string{"A", "B", "C"},
			size:  3,
			expected: [][]string{
				{"A", "B", "C"},
			},
		},
		{
			name:     "zero size",
			items:    []string{"A", "B"},
			size:     0,
			expected: nil,
		},
		{
			name:     "negative size",
			items:    []string{"A", "B"},
			size:     -1,
			expected: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := chunkStrings(tc.items, tc.size)
			assert.Equal(t, tc.expected, result)
		})
	}
}