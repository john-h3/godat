# queue

## benchmark

```
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^(BenchmarkLinkedQueueEnqueue|BenchmarkSliceQueueEnqueue)$ github.com/john-h3/godat/queue

goos: linux
goarch: amd64
pkg: github.com/john-h3/godat/queue
cpu: Intel(R) Celeron(R) CPU  J1800  @ 2.41GHz
BenchmarkLinkedQueueEnqueue-2   	 6572722	       193.6 ns/op	      16 B/op	       1 allocs/op
BenchmarkSliceQueueEnqueue-2    	20397871	        96.59 ns/op	      47 B/op	       0 allocs/op
PASS
ok  	github.com/john-h3/godat/queue	3.621s
```

## test

```
Running tool: /usr/local/go/bin/go test -timeout 30s -run ^(TestQueue|TestLinkedQueueEnqueue|TestSliceQueueEnqueue|TestLinkedQueueDequeue|TestSliceQueueDequeue|TestLinkedQueuePeek|TestSliceQueuePeek|TestLinkedQueueClear|TestSliceQueueClear|TestLinkedQueueSize|TestSliceQueueSize)$ github.com/john-h3/godat/queue

=== RUN   TestQueue
--- PASS: TestQueue (0.00s)
=== RUN   TestLinkedQueueEnqueue
--- PASS: TestLinkedQueueEnqueue (0.00s)
=== RUN   TestSliceQueueEnqueue
--- PASS: TestSliceQueueEnqueue (0.00s)
=== RUN   TestLinkedQueueDequeue
--- PASS: TestLinkedQueueDequeue (0.00s)
=== RUN   TestSliceQueueDequeue
--- PASS: TestSliceQueueDequeue (0.00s)
=== RUN   TestLinkedQueuePeek
--- PASS: TestLinkedQueuePeek (0.00s)
=== RUN   TestSliceQueuePeek
--- PASS: TestSliceQueuePeek (0.00s)
=== RUN   TestLinkedQueueClear
--- PASS: TestLinkedQueueClear (0.00s)
=== RUN   TestSliceQueueClear
--- PASS: TestSliceQueueClear (0.00s)
=== RUN   TestLinkedQueueSize
--- PASS: TestLinkedQueueSize (0.00s)
=== RUN   TestSliceQueueSize
--- PASS: TestSliceQueueSize (0.00s)
PASS
ok      github.com/john-h3/godat/queue  0.010s

```