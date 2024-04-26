# stack

## benchmark

```
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^(BenchmarkSliceStackPush|BenchmarkLinkedStackPush)$ github.com/john-h3/godat/stack

goos: linux
goarch: amd64
pkg: github.com/john-h3/godat/stack
cpu: Intel(R) Celeron(R) CPU  J1800  @ 2.41GHz
BenchmarkSliceStackPush-2    	19715175	       108.1 ns/op	      48 B/op	       0 allocs/op
BenchmarkLinkedStackPush-2   	 6394189	       212.2 ns/op	      16 B/op	       1 allocs/op
PASS
ok  	github.com/john-h3/godat/stack	3.807s
```

## test

```
Running tool: /usr/local/go/bin/go test -timeout 30s -run ^(TestStack|TestLinkedStackPush|TestSliceStackPush|TestLinkedStackPop|TestSliceStackPop|TestLinkedStackPeek|TestSliceStackPeek|TestLinkedStackSize|TestSliceStackSize|TestLinkedStackClear|TestSliceStackClear)$ github.com/john-h3/godat/stack

=== RUN   TestStack
--- PASS: TestStack (0.00s)
=== RUN   TestLinkedStackPush
--- PASS: TestLinkedStackPush (0.00s)
=== RUN   TestSliceStackPush
--- PASS: TestSliceStackPush (0.00s)
=== RUN   TestLinkedStackPop
--- PASS: TestLinkedStackPop (0.00s)
=== RUN   TestSliceStackPop
--- PASS: TestSliceStackPop (0.00s)
=== RUN   TestLinkedStackPeek
--- PASS: TestLinkedStackPeek (0.00s)
=== RUN   TestSliceStackPeek
--- PASS: TestSliceStackPeek (0.00s)
=== RUN   TestLinkedStackSize
--- PASS: TestLinkedStackSize (0.00s)
=== RUN   TestSliceStackSize
--- PASS: TestSliceStackSize (0.00s)
=== RUN   TestLinkedStackClear
--- PASS: TestLinkedStackClear (0.00s)
=== RUN   TestSliceStackClear
--- PASS: TestSliceStackClear (0.00s)
PASS
ok      github.com/john-h3/godat/stack  0.010s
```
