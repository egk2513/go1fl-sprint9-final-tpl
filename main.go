package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements возвращает срез случайных положительных целых чисел заданного размера.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = rand.Intn(1_000_000) + 1
	}

	return result
}

// maximum возвращает максимальное число в срезе.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxVal := data[0]
	for _, v := range data[1:] {
		if v > maxVal {
			maxVal = v
		}
	}

	return maxVal
}

// maxChunks находит максимальное значение, разбивая срез на CHUNKS частей и обрабатывая их в параллельных горутинах.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	maxValues := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		go func(chunkIndex int) {
			defer wg.Done()

			startIdx := chunkIndex * chunkSize
			endIdx := startIdx + chunkSize

			// Обрабатываем последний чанк, чтобы включить оставшиеся элементы
			if chunkIndex == CHUNKS-1 {
				endIdx = len(data)
			}

			chunk := data[startIdx:endIdx]
			maxValues[chunkIndex] = maximum(chunk)
		}(i)
	}

	wg.Wait()

	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	startTime := time.Now()
	maxVal := maximum(data)
	elapsed := time.Since(startTime).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxVal, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	startTime = time.Now()
	maxChunksVal := maxChunks(data)
	chunksElapsed := time.Since(startTime).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxChunksVal, chunksElapsed)
}
