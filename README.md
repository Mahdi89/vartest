# vartest

Understanding variables/channels in Go better in terms of performance and energy!

Generally variables are known to be dataflow unfriendly! as they usually infer several read-after-write and write-after-read hazard checks etc.

In this experiment I am seeing ~1000x performance degradation when I avoid using variables and use channels instead. I understand that by using `make chan` runtime gets involved and makes syscalls to the OS that eventually affect performance. However, when dealing with bigger memory chunks this performance gap decreases to 69x and 47x and so on. I believe this observation is something to be taken account when dealingwith big data.. 


try `make bench`

```
go test -run=5 -bench=.
goos: linux
goarch: amd64
BenchmarkWithVariable-4      	2000000000	         0.32 ns/op
BenchmarkWithoutVariable-4   	 5000000	       268 ns/op
BenchmarkWithArray-4         	100000000	        13.0 ns/op
BenchmarkWithoutArray-4      	 2000000	       907 ns/op
BenchmarkWithBigArray-4      	  100000	     15710 ns/op
BenchmarkWithoutBigArray-4   	    2000	    739518 ns/op
PASS
ok  	_/vartest	9.659s

```
