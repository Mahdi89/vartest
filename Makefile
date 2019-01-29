NAME := Variable Test
DESC := Understanding Go channels better!

.PHONY: test bench

test:
	go test

bench: 	
	go test -run=5 -bench=.
