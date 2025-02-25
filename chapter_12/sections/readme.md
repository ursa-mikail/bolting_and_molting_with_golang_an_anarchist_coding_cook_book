Golang sample_scaffold_working_structure

<pre>
.
├── go.mod
├── lib
│    ├── p0
│    │    └── p0.go
│    └── target.go
├── main.go
├── readme.md
├── tests
│   └── target_test.go
└── util
    └── util_00.go

% go test -v -run 'TestDefensePass|TestDefenseFail' -bench BenchmarkCacheDefense ./tests  
=== RUN   TestDefensePass
--- PASS: TestDefensePass (0.00s)
=== RUN   TestDefenseFail
--- PASS: TestDefenseFail (0.00s)
PASS
ok      example.com/demo/tests  0.219s

% go test ./tests
ok      example.com/demo/tests  0.381s

% go test -bench=. ./tests 
goos: darwin
goarch: arm64
pkg: example.com/demo/tests
cpu: Apple M1
BenchmarkDefense/Size_16-8          506158704            2.068 ns/op
BenchmarkDefense/Size_32-8          576874540            2.083 ns/op
BenchmarkDefense/Size_64-8          579767796            2.072 ns/op
BenchmarkDefense/Size_128-8         576997975            2.085 ns/op
BenchmarkDefense/Size_256-8         574993299            2.068 ns/op
BenchmarkDefense/Size_512-8         575291240            2.086 ns/op
BenchmarkDefense/Size_1024-8        579340242            2.067 ns/op
BenchmarkDefense/Size_2048-8        576988728            2.078 ns/op
BenchmarkDefense/Size_4096-8        580605243            2.096 ns/op
PASS
ok      example.com/demo/tests  12.809s

% go test -cover ./tests
ok      example.com/demo/tests  0.223s  coverage: [no statements]

</pre>

