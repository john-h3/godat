# deque

## benchmark

```
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^(BenchmarkDequeEnqueueFront|BenchmarkDequeEnqueueBack)$ github.com/john-h3/godat/deque

goos: linux
goarch: amd64
pkg: github.com/john-h3/godat/deque
cpu: Intel(R) Celeron(R) CPU  J1800  @ 2.41GHz
BenchmarkDequeEnqueueFront-2   	 6744954	       218.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkDequeEnqueueBack-2    	 7104967	       207.4 ns/op	      24 B/op	       1 allocs/op
PASS
ok  	github.com/john-h3/godat/deque	3.645s
```

## test

```
Running tool: /usr/local/go/bin/go test -timeout 30s -run ^(TestDequeEnqueueFront|TestDequeEnqueueBack|TestDequeDequeueFront|TestDequeDequeueBack|TestDequePeekFront|TestDequePeekBack|TestDequeClear|TestDequeSize)$ github.com/john-h3/godat/deque

=== RUN   TestDequeEnqueueFront
--- PASS: TestDequeEnqueueFront (0.00s)
=== RUN   TestDequeEnqueueBack
--- PASS: TestDequeEnqueueBack (0.00s)
=== RUN   TestDequeDequeueFront
--- PASS: TestDequeDequeueFront (0.00s)
=== RUN   TestDequeDequeueBack
--- PASS: TestDequeDequeueBack (0.00s)
=== RUN   TestDequePeekFront
--- PASS: TestDequePeekFront (0.00s)
=== RUN   TestDequePeekBack
--- PASS: TestDequePeekBack (0.00s)
=== RUN   TestDequeClear
--- PASS: TestDequeClear (0.00s)
=== RUN   TestDequeSize
--- PASS: TestDequeSize (0.00s)
PASS
ok      github.com/john-h3/godat/deque  0.009s
```