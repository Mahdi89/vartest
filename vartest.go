package vartest

func WithVariable() (int, int) {

	a := 2 + 3
	b := a + 3
	c := a + 5
	return b, c
}

func WithoutVariable() (int, int) {

	a := make(chan int, 2)
	a <- 2 + 3
	a <- 2 + 3
	return (<-a + 3), (<-a + 5)
}
