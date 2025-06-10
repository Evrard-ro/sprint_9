package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxChunksVsMaximum(t *testing.T) {
	slice := generateRandomElements(10000)
	maxSingle := maximum(slice)
	maxChunksVal := maxChunks(slice)

	assert.Equal(t, maxSingle, maxChunksVal)
}

func TestMaximum(t *testing.T) {

	sliceIsEmpty := maximum([]int{})
	assert.Equal(t, sliceIsEmpty, 0)

	sliceOneElement := maximum([]int{42})
	assert.Equal(t, sliceOneElement, 42)

	sliceNormal := maximum([]int{10, 50, 30, 20, 40})
	assert.Equal(t, sliceNormal, 50)

}

func TestGenerateRandomElements(t *testing.T) {

	zeroSize := generateRandomElements(0)
	assert.Equal(t, len(zeroSize), 0)

	normalSize := generateRandomElements(100)
	assert.Equal(t, len(normalSize), 100)

	negativeSize := generateRandomElements(-2)
	assert.Equal(t, len(negativeSize), 0)

}
