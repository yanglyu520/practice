package main

import (
	"math/rand"
	"testing"
)

func generateArray(n int) []int32 {
	a := make([]int32, n)

	for i, _ := range a {
		a[i] = int32(rand.Int())
	}
	return a
}

func BenchmarkReverseArray1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseArray1(generateArray(1000))
	}
}

func BenchmarkReverseArrayInplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseArrayInplace(generateArray(1000))
	}
}
