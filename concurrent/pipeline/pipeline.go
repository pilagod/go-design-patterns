package pipeline

func LaunchPipeline(amount int) int {
	return <-sum(power(generator(amount)))
}

func generator(max int) <-chan int {
	out := make(chan int, 100)

	go func() {
		for i := 1; i <= max; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func power(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		var res int

		for v := range in {
			res += v
		}
		out <- res
		close(out)
	}()
	return out
}
