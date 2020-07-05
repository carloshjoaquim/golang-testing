package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBubbleSortOrderASC(t *testing.T) {
	elements := GetElements(10)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0])
	assert.EqualValues(t, 0, elements[len(elements)-1])

	timeoutChan := make(chan bool, 1)
	defer close(timeoutChan)

	go func() {
		BubbleSort(elements)
		timeoutChan <- false
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		timeoutChan <- true
	}()

	if <-timeoutChan {
		assert.Fail(t, "Bubble sort took more than 500 ms")
		return
	}

	BubbleSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])
}

func TestSortOrderASC(t *testing.T) {
	elements := GetElements(10)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0])
	assert.EqualValues(t, 0, elements[len(elements)-1])

	Sort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])
}

func TestBubbleSortAlreadySorted(t *testing.T) {
	elements := []int{0, 1, 2, 3, 4, 5}

	assert.NotNil(t, elements)
	assert.EqualValues(t, 6, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 5, elements[len(elements)-1])

	BubbleSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 6, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 5, elements[len(elements)-1])
}

func getElements(n int) []int {
	result := make([]int, n)
	j := 0

	for i := n - 1; i > 0; i-- {
		result[j] = i
		j++
	}
	return result
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := GetElements(10000)

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := GetElements(10000)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
