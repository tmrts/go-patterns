package generator

func Range(start int, end int, step int) chan int {
	c := make(chan int)

	go func() {
		result := start
		for result < end {
			c <- result
			result = result + step
		}

		close(c)
	}()

	return c
}

func main() {
	// print the numbers from 3 through 47 with a step size of 2
	for i := range Range(3, 47, 2) {
		println(i)
	}
}
