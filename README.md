# vartest

Understanding variables/channels in Go better!

I genuinely don't like using variables in my code maybe it's because I am an engineer with hardware background and variables in hardware are not dataflow friendly! as they usually require several read-after-write and write-after-read hazard checks etc.

In this simple experiment I am seeing ~1000x performance degradation when I avoid using variables and use channels instead. I understand that by using `make chan` runtime gets involved and makes syscalls to the OS that eventually affect performance, but is there a way to write variable less code without worrying about performance too much?   

try `make bench`

```
go test -run=5 -bench=.
goos: linux
goarch: amd64
BenchmarkWithVariable-4      	2000000000	         0.32 ns/op
BenchmarkWithoutVariable-4   	 5000000	       239 ns/op
PASS
ok  	_/vartest	2.152s
```
