package main

import "testing"

func BenchmarkFibRecursive20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibRecursive(20)
	}
}

func BenchmarkFibIterative20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibIterative(20)
	}
}

func BenchmarkFibRecursive30(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibRecursive(30)
	}
}

func BenchmarkFibIterative30(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibIterative(30)
	}
}

