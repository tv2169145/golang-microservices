package untils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestBubbleSortWorstCase(t *testing.T) {
	intSlice := []int{9, 7, 8, 5, 6}
	fmt.Println(intSlice)
	BubbleSort(intSlice)
	fmt.Println(intSlice)
	assert.NotNil(t, intSlice, "slice is nil")
	assert.EqualValues(t, 5, len(intSlice), "len not equal 5")
	assert.EqualValues(t, 5, intSlice[0])
	assert.EqualValues(t, 6, intSlice[1])
	assert.EqualValues(t, 7, intSlice[2])
	assert.EqualValues(t, 8, intSlice[3])
	assert.EqualValues(t, 9, intSlice[4])
}
func TestBubbleSortBestCase(t *testing.T) {
	intSlice := []int{5, 6, 7, 8, 9}
	fmt.Println(intSlice)
	BubbleSort(intSlice)
	fmt.Println(intSlice)
	assert.NotNil(t, intSlice, "slice is nil")
	assert.EqualValues(t, 5, len(intSlice), "len not equal 5")
	assert.EqualValues(t, 5, intSlice[0])
	assert.EqualValues(t, 6, intSlice[1])
	assert.EqualValues(t, 7, intSlice[2])
	assert.EqualValues(t, 8, intSlice[3])
	assert.EqualValues(t, 9, intSlice[4])
}

func TestBubbleSortNilSlice(t *testing.T) {
	BubbleSort(nil)
}


func getElement(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElement(t *testing.T) {
	result := getElement(5)
	assert.NotNil(t, result, "result is nil")
	assert.EqualValues(t, 4, result[0])
	assert.EqualValues(t, 3, result[1])
	assert.EqualValues(t, 2, result[2])
	assert.EqualValues(t, 1, result[3])
	assert.EqualValues(t, 0, result[4])
}

func BenchmarkBubbleSort10(b *testing.B) {
	els := getElement(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkSort10(b *testing.B) {
	els := getElement(10)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	els := getElement(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkSort1000(b *testing.B) {
	els := getElement(1000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkBubbleSort10000(b *testing.B) {
	els := getElement(10000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkSort10000(b *testing.B) {
	els := getElement(10000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkSort100000(b *testing.B) {
	els := getElement(10000)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

func TestSortUseBubbleSort(t *testing.T) {
	els := getElement(5)
	Sort(els)
	assert.NotNil(t, els)
	assert.EqualValues(t, 0, els[0])
	assert.EqualValues(t, 1, els[1])
	assert.EqualValues(t, 2, els[2])
	assert.EqualValues(t, 3, els[3])
	assert.EqualValues(t, 4, els[4])
}

func TestSortUseSort(t *testing.T) {
	els := getElement(1001)
	Sort(els)
	assert.NotNil(t, els)
	assert.EqualValues(t, 0, els[0])
	assert.EqualValues(t, 1, els[1])
	assert.EqualValues(t, 2, els[2])
	assert.EqualValues(t, 3, els[3])
	assert.EqualValues(t, 4, els[4])
}
