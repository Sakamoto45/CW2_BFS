package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const N = 5
const size = 500

func BenchmarkBFS(b *testing.B) {
	cubeEdges := genCube(size)
	var parTime time.Duration
	var seqTime time.Duration
	b.Run(fmt.Sprintf("Parallel"), func(b *testing.B) {
		for range N {
			b.StartTimer()
			dist := parallelBFS(cubeEdges, 0)
			b.StopTimer()

			require.True(b, validate(dist, size))
		}
		parTime = b.Elapsed()
	})
	b.Run(fmt.Sprintf("Sequential"), func(b *testing.B) {
		for range N {
			b.StartTimer()
			dist := sequentialBFS(cubeEdges, 0)
			b.StopTimer()
			require.True(b, validate(dist, size))
		}
		seqTime = b.Elapsed()

	})
	fmt.Println(float32(seqTime.Nanoseconds())/float32(parTime.Nanoseconds()), "times faster")
}

func cubeIdx(i, j, k, size int32) int32 {
	return (i*size+j)*size + k
}

func genCube(size int32) [][]int32 {
	result := make([][]int32, size*size*size)
	for i := range size {
		for j := range size {
			for k := range size {
				idx := cubeIdx(i, j, k, size)
				result[idx] = make([]int32, 0, 6)
				if i > 0 {
					result[idx] = append(result[idx], cubeIdx(i-1, j, k, size))
				}
				if j > 0 {
					result[idx] = append(result[idx], cubeIdx(i, j-1, k, size))
				}
				if k > 0 {
					result[idx] = append(result[idx], cubeIdx(i, j, k-1, size))
				}
				if i < size-1 {
					result[idx] = append(result[idx], cubeIdx(i+1, j, k, size))
				}
				if j < size-1 {
					result[idx] = append(result[idx], cubeIdx(i, j+1, k, size))
				}
				if k < size-1 {
					result[idx] = append(result[idx], cubeIdx(i, j, k+1, size))
				}
			}
		}
	}
	return result
}

func validate(dist []int32, size int32) bool {
	for i := range size {
		for j := range size {
			for k := range size {
				if dist[cubeIdx(i, j, k, size)] != i+j+k {
					return false
				}
			}
		}
	}
	return true
}

func printDist(dist []int32, size int32) {
	for i := range size {
		for j := range size {
			for k := range size {
				fmt.Print(dist[cubeIdx(i, j, k, size)], " ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
