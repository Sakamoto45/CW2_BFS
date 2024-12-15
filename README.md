# CW2. BFS. Ткаченко Егор Андреевич

[Реализация](bfs.go)

[Бенчмарк для сравнения (с проверкой корректности)](bfs_test.go)

Результат запуска с командой `go test -bench=. -benchtime=1x -benchmem -cpu=4`:
```
goos: linux
goarch: amd64
pkg: bfs
cpu: AMD Ryzen 7 5800X3D 8-Core Processor           
BenchmarkBFS/Parallel-4                        1        6213832891 ns/op        2589465592 B/op   484205 allocs/op
BenchmarkBFS/Sequential-4                      1        17859192942 ns/op       2541177520 B/op      295 allocs/op
2.8741025 times faster
PASS
ok      bfs     30.143s
```

Достигнуто ускорение в 2.8741025 раза