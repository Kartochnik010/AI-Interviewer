package main

import "fmt"

func main() {
	a := make(chan int)
	stop := make(chan struct{})

	res := calculator(a, stop)
	go func() {
		for i := 1; i <= 10; i++ {
			a <- i
		}
		stop <- struct{}{}
	}()
	fmt.Println(<-res)

}
func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	resChan := make(chan int)

	go func() {
		a := &[]int{}
		defer close(resChan)
		for {
			select {
			case arg := <-arguments:
				*a = append(*a, arg)
			case <-done:
				res := 0
				for _, v := range *a {
					res += v
				}
				resChan <- res
				return
			}
		}

	}()

	return resChan
}
