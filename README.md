# vartest

Understanding variables/channels in Go better in terms of performance and energy!

Generally variables are known to be dataflow unfriendly! as they usually infer several read-after-write and write-after-read hazard checks etc.

In this experiment I am seeing ~1000x performance degradation when I avoid using variables and use channels instead. I understand that by using `make chan` runtime gets involved and makes syscalls to the OS that eventually affect performance. However, when dealing with bigger memory chunks this performance gap decreases to 69x, 47x, 33x and so on. I believe this observation is something to be taken into account when dealing with big data.. 


try `make bench`

```
go test -run=5 -bench=.
goos: linux
goarch: amd64
BenchmarkWithVariable-4         	2000000000	         0.33 ns/op
BenchmarkWithoutVariable-4      	 5000000	       251 ns/op
BenchmarkWithArray-4            	100000000	        12.7 ns/op
BenchmarkWithoutArray-4         	 2000000	       919 ns/op
BenchmarkWithBigArray-4         	  100000	     15600 ns/op
BenchmarkWithoutBigArray-4      	    2000	    724280 ns/op
BenchmarkWithBiggerArray-4      	    1000	   2109320 ns/op
BenchmarkWithoutBiggerArray-4   	      20	  71891180 ns/op
PASS
ok  	_/vartest	13.372s

```
