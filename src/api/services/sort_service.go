package services

import "github.com/carloshjoaquim/golang-testing/src/api/utils/sort"

const (
	privateConst = "private"
)

func Sort(elements []int) {
	if len(elements) < 10000 {
		sort.BubbleSort(elements)
		return
	}
	sort.Sort(elements)
}
