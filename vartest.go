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

func WithArray() (int, int) {

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := 0
	c := 0

	for i := 0; i < 10; i++ {

		b += a[i]
		c -= a[i]
	}
	return b, c
}

func WithoutArray() (int, int) {

	a := make(chan int, 10)
	b := 0
	c := 0
	temp := 0

	for i := 1; i < 11; i++ {

		a <- i
	}

	for j := 0; j < 10; j++ {

		temp = <-a
		b += temp
		c -= temp
	}

	return (b), (c)
}

func WithBigArray() (int, int) {

	a := [10000]int{}
	b := 0
	c := 0

	for i := 0; i < 10000; i++ {

		a[i] = i

	}
	for j := 0; j < 10000; j++ {

		b += a[j]
		c -= a[j]
	}
	return b, c
}

func WithoutBigArray() (int, int) {

	a := make(chan int, 10000)
	b := 0
	c := 0
	temp := 0

	for i := 0; i < 10000; i++ {

		a <- i
	}

	for j := 0; j < 10000; j++ {

		temp = <-a
		b += temp
		c -= temp
	}

	return (b), (c)
}
