package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

const notVisited = -1

func sequentialBFS(edges [][]int32, startVertex int32) []int32 {
	result := make([]int32, len(edges))
	for i := range result {
		result[i] = notVisited
	}
	result[startVertex] = 0

	curFrontier := []int32{startVertex}
	nextFrontier := make([]int32, 0)

	for len(curFrontier) != 0 {

		for _, from := range curFrontier {
			for _, to := range edges[from] {
				if result[to] == notVisited {
					result[to] = result[from] + 1
					nextFrontier = append(nextFrontier, to)
				}
			}
		}

		curFrontier, nextFrontier = nextFrontier, curFrontier
		nextFrontier = nextFrontier[:0]
	}

	return result
}

func parallelBFS(edges [][]int32, startVertex int32) []int32 {
	wg := sync.WaitGroup{}
	workers := runtime.GOMAXPROCS(0) * 16

	result := make([]int32, len(edges))
	for i := range result {
		result[i] = notVisited
	}
	result[startVertex] = 0

	curFrontier := []int32{startVertex}
	personalNextFrontier := make([][]int32, workers)

	for len(curFrontier) != 0 {
		wg.Add(workers)
		for worker := range workers {
			go func() {
				personalNextFrontier[worker] = personalNextFrontier[worker][:0]
				for _, from := range curFrontier[len(curFrontier)*worker/workers : len(curFrontier)*(worker+1)/workers] {
					for _, to := range edges[from] {
						if result[to] == notVisited {
							if atomic.CompareAndSwapInt32(&result[to], notVisited, result[from]+1) {
								personalNextFrontier[worker] = append(personalNextFrontier[worker], to)
							}
						}
					}
				}
				wg.Done()
			}()
		}
		wg.Wait()

		// currently it takes ~1% of function time
		curFrontier = curFrontier[:0]
		for worker := range workers {
			curFrontier = append(curFrontier, personalNextFrontier[worker]...)
		}
	}

	return result
}

// I tried
// use array of atomic instead of atomic to pointer
// preallocate memory for workers
// parallelize copy from personalNextFrontier to curFrontier or do it in goroutines
