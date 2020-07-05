package services

import (
	"github.com/carloshjoaquim/golang-testing/src/api/utils/sort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstants(t *testing.T) {
	if privateConst != "private" {
		t.Error("privateConst should be 'private'")
	}
}

func TestSort(t *testing.T) {
	elements := sort.GetElements(10)

	Sort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])
}

func TestSortMore10000(t *testing.T) {
	elements := sort.GetElements(10001)

	Sort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10001, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 10000, elements[len(elements)-1])
}
