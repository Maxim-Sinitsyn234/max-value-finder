package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{
			name: "нулевой размер",
			size: 0,
			want: 0,
		},
		{
			name: "отрицательный размер",
			size: -5,
			want: 0,
		},
		{
			name: "один элемент",
			size: 1,
			want: 1,
		},
		{
			name: "тысяча элементов",
			size: 1000,
			want: 1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			if len(got) != tt.want {
				t.Fatalf("generateRandomElements(%d): получено %d элементов, ожидалось %d", tt.size, len(got), tt.want)
			}

			for i, value := range got {
				if value <= 0 {
					t.Fatalf("generateRandomElements(%d): элемент %d равен %d, ожидалось положительное число", tt.size, i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	big := make([]int, 10_000)
	for i := range big {
		big[i] = i
	}

	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "nil",
			data: nil,
			want: 0,
		},
		{
			name: "пустой слайс",
			data: []int{},
			want: 0,
		},
		{
			name: "один элемент",
			data: []int{42},
			want: 42,
		},
		{
			name: "максимум первый",
			data: []int{100, 2, 3},
			want: 100,
		},
		{
			name: "максимум последний",
			data: []int{1, 2, 100},
			want: 100,
		},
		{
			name: "максимум в середине",
			data: []int{1, 100, 2},
			want: 100,
		},
		{
			name: "все отрицательные",
			data: []int{-10, -3, -20},
			want: -3,
		},
		{
			name: "одинаковые значения",
			data: []int{7, 7, 7},
			want: 7,
		},
		{
			name: "большой возрастающий слайс",
			data: big,
			want: 9_999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Fatalf("maximum(%v): получено %d, ожидалось %d", tt.data, got, tt.want)
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
		{
			name: "nil",
			data: nil,
			want: 0,
		},
		{
			name: "пустой слайс",
			data: []int{},
			want: 0,
		},
		{
			name: "один элемент",
			data: []int{5},
			want: 5,
		},
		{
			name: "меньше CHUNKS",
			data: []int{3, 7, 1, 9, 2},
			want: 9,
		},
		{
			name: "ровно CHUNKS",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want: 8,
		},
		{
			name: "не делится на CHUNKS",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100},
			want: 100,
		},
		{
			name: "все отрицательные",
			data: []int{-5, -1, -9, -2, -8, -3, -7, -4},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Fatalf("maxChunks(%v): получено %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	data := make([]int, 10_000)
	for i := range data {
		data[i] = (i*37)%1000 + 1
	}

	got := maxChunks(data)
	want := maximum(data)
	if got != want {
		t.Fatalf("maxChunks: получено %d, maximum: получено %d", got, want)
	}
}
