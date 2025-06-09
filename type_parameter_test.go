package belajar_golang_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}


func TestLength(t *testing.T) {
	var resultString string = Length[string]("Belajar Golang Generics")

	assert.Equal(t, "Belajar Golang Generics", resultString)

	var resultInt int = Length[int](10)

	assert.Equal(t, 10, resultInt)
}
func TestSample(t *testing.T) {
	assert.True(t, true)
}