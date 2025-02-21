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
</pre>

