# queue

## benchmark

```
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^BenchmarkLinkedQueueEnqueue$ github.com/john-h3/godat/queue

goos: linux
goarch: amd64
pkg: github.com/john-h3/godat/queue
cpu: Intel(R) Celeron(R) CPU  J1800  @ 2.41GHz
BenchmarkLinkedQueueEnqueue-2   	 6717310	       165.1 ns/op	      16 B/op	       1 allocs/op
PASS
ok  	github.com/john-h3/godat/queue	1.313s
```

## test

```
Running tool: /usr/local/go/bin/go test -timeout 30s -run ^(TestQueue|TestLinkedQueueEnqueue|TestLinkedQueueDequeue|TestPeek|TestLinkedQueueClear|TestLinkedQueueSize)$ github.com/john-h3/godat/queue

=== RUN   TestQueue
--- PASS: TestQueue (0.00s)
=== RUN   TestLinkedQueueEnqueue
--- PASS: TestLinkedQueueEnqueue (0.00s)
=== RUN   TestLinkedQueueDequeue
--- PASS: TestLinkedQueueDequeue (0.00s)
=== RUN   TestPeek
--- PASS: TestPeek (0.00s)
=== RUN   TestLinkedQueueClear
--- PASS: TestLinkedQueueClear (0.00s)
=== RUN   TestLinkedQueueSize
--- PASS: TestLinkedQueueSize (0.00s)
PASS
ok      github.com/john-h3/godat/queue  (cached)
```