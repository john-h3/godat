# queue

## benchmark

```
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^(BenchmarkLinkedQueueEnqueue|BenchmarkSliceQueueEnqueue|BenchmarkCircularQueueEnqueue)$ github.com/john-h3/godat/queue

goos: linux
goarch: amd64
pkg: github.com/john-h3/godat/queue
cpu: Intel(R) Celeron(R) CPU  J1800  @ 2.41GHz
BenchmarkLinkedQueueEnqueue-2     	 6419304	       183.9 ns/op	      16 B/op	       1 allocs/op
BenchmarkSliceQueueEnqueue-2      	17386848	        74.40 ns/op	      44 B/op	       0 allocs/op
BenchmarkCircularQueueEnqueue-2   	 7152586	       181.3 ns/op	      16 B/op	       1 allocs/op
PASS
ok  	github.com/john-h3/godat/queue	4.296s

```

## test

```
Running tool: /usr/local/go/bin/go test -timeout 30s -run ^(TestQueue|TestLinkedQueueEnqueue|TestSliceQueueEnqueue|TestCircularLinkedQueue|TestLinkedQueueDequeue|TestSliceQueueDequeue|TestCircularLinkedQueueDequeue|TestLinkedQueuePeek|TestSliceQueuePeek|TestCircularLinkedQueuePeek|TestLinkedQueueClear|TestSliceQueueClear|TestCircularLinkedQueueClear|TestLinkedQueueSize|TestSliceQueueSize|TestCircularLinkedQueueSize)$ github.com/john-h3/godat/queue

=== RUN   TestQueue
--- PASS: TestQueue (0.00s)
=== RUN   TestLinkedQueueEnqueue
--- PASS: TestLinkedQueueEnqueue (0.00s)
=== RUN   TestSliceQueueEnqueue
--- PASS: TestSliceQueueEnqueue (0.00s)
=== RUN   TestCircularLinkedQueue
--- PASS: TestCircularLinkedQueue (0.00s)
=== RUN   TestLinkedQueueDequeue
--- PASS: TestLinkedQueueDequeue (0.00s)
=== RUN   TestSliceQueueDequeue
--- PASS: TestSliceQueueDequeue (0.00s)
=== RUN   TestCircularLinkedQueueDequeue
--- PASS: TestCircularLinkedQueueDequeue (0.00s)
=== RUN   TestLinkedQueuePeek
--- PASS: TestLinkedQueuePeek (0.00s)
=== RUN   TestSliceQueuePeek
--- PASS: TestSliceQueuePeek (0.00s)
=== RUN   TestCircularLinkedQueuePeek
--- PASS: TestCircularLinkedQueuePeek (0.00s)
=== RUN   TestLinkedQueueClear
--- PASS: TestLinkedQueueClear (0.00s)
=== RUN   TestSliceQueueClear
--- PASS: TestSliceQueueClear (0.00s)
=== RUN   TestCircularLinkedQueueClear
--- PASS: TestCircularLinkedQueueClear (0.00s)
=== RUN   TestLinkedQueueSize
--- PASS: TestLinkedQueueSize (0.00s)
=== RUN   TestSliceQueueSize
--- PASS: TestSliceQueueSize (0.00s)
=== RUN   TestCircularLinkedQueueSize
--- PASS: TestCircularLinkedQueueSize (0.00s)
PASS
ok      github.com/john-h3/godat/queue  0.010s

```