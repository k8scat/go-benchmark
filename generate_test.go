// generate_test.go
package main

import (
	"math/rand"
	"testing"
	"time"
)

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func benchmarkGenerateWithCap(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateWithCap(i)
	}
}

func benchmarkGenerate(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate(i)
	}
}

func BenchmarkGenerate1000(b *testing.B)    { benchmarkGenerate(1000, b) }
func BenchmarkGenerate10000(b *testing.B)   { benchmarkGenerate(10000, b) }
func BenchmarkGenerate100000(b *testing.B)  { benchmarkGenerate(100000, b) }
func BenchmarkGenerate1000000(b *testing.B) { benchmarkGenerate(1000000, b) }

func BenchmarkGenerateWithCap1000(b *testing.B)    { benchmarkGenerateWithCap(1000, b) }
func BenchmarkGenerateWithCap10000(b *testing.B)   { benchmarkGenerateWithCap(10000, b) }
func BenchmarkGenerateWithCap100000(b *testing.B)  { benchmarkGenerateWithCap(100000, b) }
func BenchmarkGenerateWithCap1000000(b *testing.B) { benchmarkGenerateWithCap(1000000, b) }
