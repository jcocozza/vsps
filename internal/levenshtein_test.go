package internal

import "testing"

func TestLevenshtein(t *testing.T) {
	tests := []struct {
		name     string
		a, b     string
		expected int
	}{
		{"kitten, sitting", "kitten", "sitting", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := levenshteinDistance(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%s, %s) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
