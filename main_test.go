package main

import (
	"math"
	"testing"
)

// generateRandomElements

func TestGenerateRandomElements_ZeroSize(t *testing.T) {
	result := generateRandomElements(0)

	if len(result) != 0 {
		t.Errorf("ожидался пустой срез при size = 0, получено len = %d", len(result))
	}
}

func TestGenerateRandomElements_NegativeSize(t *testing.T) {
	for _, size := range []int{-1, -100, math.MinInt} {
		result := generateRandomElements(size)

		if len(result) != 0 {
			t.Errorf("ожидался пустой срез при size = %d, получено len = %d", size, len(result))
		}
	}
}

func TestGenerateRandomElements_PositiveSize(t *testing.T) {
	const size = 1000

	result := generateRandomElements(size)

	if len(result) != size {
		t.Fatalf("ожидалась длина среза %d, получено %d", size, len(result))
	}

	// Все элементы должны быть положительными (от 1 до 1_000_000).
	for i, v := range result {
		if v <= 0 {
			t.Errorf("элемент [%d] = %d не является положительным", i, v)
		}
	}
}

func TestGenerateRandomElements_OneElement(t *testing.T) {
	result := generateRandomElements(1)

	if len(result) != 1 {
		t.Fatalf("ожидалась длина среза 1, получено %d", len(result))
	}
	if result[0] <= 0 {
		t.Errorf("элемент %d не является положительным", result[0])
	}
}

// maximum

func TestMaximum_EmptySlice(t *testing.T) {
	result := maximum([]int{})

	if result != 0 {
		t.Errorf("ожидалось 0 для пустого среза, получено %d", result)
	}
}

func TestMaximum_NilSlice(t *testing.T) {
	var data []int

	result := maximum(data)

	if result != 0 {
		t.Errorf("ожидалось 0 для nil-среза, получено %d", result)
	}
}

func TestMaximum_OneElement(t *testing.T) {
	result := maximum([]int{42})

	if result != 42 {
		t.Errorf("ожидалось 42, получено %d", result)
	}
}

func TestMaximum_NegativeNumbers(t *testing.T) {
	result := maximum([]int{-10, -5, -7, -100, -1})

	if result != -1 {
		t.Errorf("ожидалось -1, получено %d", result)
	}
}

func TestMaximum_AllEqual(t *testing.T) {
	result := maximum([]int{7, 7, 7, 7, 7})

	if result != 7 {
		t.Errorf("ожидалось 7, получено %d", result)
	}
}

func TestMaximum_MaxAtEdges(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"максимум в начале", []int{9, 1, 2, 3, 4}, 9},
		{"максимум в конце", []int{1, 2, 3, 4, 9}, 9},
		{"максимум в середине", []int{1, 2, 9, 3, 4}, 9},
		{"упорядоченный по возрастанию", []int{1, 2, 3, 4, 5}, 5},
		{"упорядоченный по убыванию", []int{5, 4, 3, 2, 1}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaximum_IntMax(t *testing.T) {
	result := maximum([]int{0, math.MaxInt, 1})

	if result != math.MaxInt {
		t.Errorf("ожидалось %d, получено %d", math.MaxInt, result)
	}
}

// maxChunks

func TestMaxChunks_EmptySlice(t *testing.T) {
	result := maxChunks([]int{})

	if result != 0 {
		t.Errorf("ожидалось 0 для пустого среза, получено %d", result)
	}
}

func TestMaxChunks_NilSlice(t *testing.T) {
	var data []int

	result := maxChunks(data)

	if result != 0 {
		t.Errorf("ожидалось 0 для nil-среза, получено %d", result)
	}
}

func TestMaxChunks_OneElement(t *testing.T) {
	result := maxChunks([]int{5})

	if result != 5 {
		t.Errorf("ожидалось 5, получено %d", result)
	}
}

func TestMaxChunks_ShorterThanChunks(t *testing.T) {
	// Длина среза меньше количества чанков: chunkSize = 0.
	data := []int{3, 9, 4}

	result := maxChunks(data)

	if result != 9 {
		t.Errorf("ожидалось 9, получено %d", result)
	}
}

func TestMaxChunks_NotDivisibleByChunks(t *testing.T) {
	// Длина не делится на CHUNKS без остатка: хвост должен быть учтён.
	const size = CHUNKS*3 + 1

	data := make([]int, size)
	for i := range data {
		data[i] = i + 1
	}

	result := maxChunks(data)

	if result != size {
		t.Errorf("ожидалось %d, получено %d", size, result)
	}
}

func TestMaxChunks_MaxInTail(t *testing.T) {
	// Максимум находится в «хвосте», который добавляется к последнему чанку.
	data := make([]int, CHUNKS*2+1)
	data[len(data)-1] = math.MaxInt

	result := maxChunks(data)

	if result != math.MaxInt {
		t.Errorf("ожидалось %d, получено %d", math.MaxInt, result)
	}
}

func TestMaxChunks_NegativeNumbers(t *testing.T) {
	data := make([]int, CHUNKS*2)
	for i := range data {
		data[i] = -(i + 1)
	}

	result := maxChunks(data)

	if result != -1 {
		t.Errorf("ожидалось -1, получено %d", result)
	}
}

func TestMaxChunks_MatchesSequentialMaximum(t *testing.T) {
	data := generateRandomElements(10_000)

	want := maximum(data)
	got := maxChunks(data)

	if got != want {
		t.Errorf("максимумы не совпали: maximum = %d, maxChunks = %d", want, got)
	}
}

func TestMaxChunks_ConsistencyAcrossRuns(t *testing.T) {
	data := generateRandomElements(1000)
	want := maximum(data)

	for i := 0; i < 10; i++ {
		if got := maxChunks(data); got != want {
			t.Fatalf("прогон %d: maxChunks = %d, ожидалось %d", i, got, want)
		}
	}
}
