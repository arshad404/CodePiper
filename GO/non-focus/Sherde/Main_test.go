package main

import (
	"fmt"
	"testing"
)

func BenchmarkJSON(b *testing.B) {
	// Small, Medium, Large, Extra Large
	var input = []int{10, 50, 100, 400}

	for _, v := range input {
		albums := BuildResponse(v)
		b.Run(fmt.Sprintf("input_size_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MarshallingUnmarshallingJSON(albums)
			}
		})
	}
}

func BenchmarkSonic(b *testing.B) {
	// Small, Medium, Large, Extra Large
	var input = []int{10, 50, 100, 400}

	for _, v := range input {
		albums := BuildResponse(v)
		b.Run(fmt.Sprintf("input_size_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MarshallingUnmarshallingSONIC(albums)
			}
		})
	}
}
