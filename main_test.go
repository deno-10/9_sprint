package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{"zero size", 0, 0},
		{"negative size", -5, 0},
		{"one element", 1, 1},
		{"many elements", 1000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)
			if data == nil {
				t.Fatalf("got nil slice, want non-nil")
			}
			if len(data) != tt.wantLen {
				t.Fatalf("len = %d, want %d", len(data), tt.wantLen)
			}
			for i, v := range data {
				if v <= 0 {
					t.Fatalf("data[%d] = %d, want positive", i, v)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"one element", []int{7}, 7},
		{"max first", []int{9, 1, 2}, 9},
		{"max last", []int{1, 2, 9}, 9},
		{"max middle", []int{1, 9, 2}, 9},
		{"all equal", []int{5, 5, 5}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Fatalf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"fewer than CHUNKS", []int{3, 8, 1}, 8},
		{"exactly CHUNKS", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"max in tail remainder", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 99}, 99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Fatalf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}

	data := generateRandomElements(100_003)
	if got, want := maxChunks(data), maximum(data); got != want {
		t.Fatalf("maxChunks = %d, maximum = %d", got, want)
	}
}
