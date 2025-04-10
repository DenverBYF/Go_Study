package language

import (
	"fmt"
	"sync"
	"time"
)

func TestChan() {
	ch1 := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			fmt.Println("send", i)
			time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case val, ok := <-ch1:
				if !ok {
					ch1 = nil
					fmt.Println("close ch")
					return
				}
				fmt.Println("rcv:", val)
			}
		}
	}()

	go func() {
		time.Sleep(15 * time.Second)
		close(ch1)
	}()

	wg.Wait()
}

func PrintNumAndLetter() {
	numCh := make(chan bool)
	letterCh := make(chan bool)

	wg := sync.WaitGroup{}
	go func() {
		i := 0
		for {
			select {
			case <-numCh:
				fmt.Println(i)
				i++
				letterCh <- true
			}
		}
	}()

	wg.Add(1)
	go func() {
		i := 'A'
		for {
			select {
			case <-letterCh:
				if i > 'Z' {
					wg.Done()
					return
				}
				fmt.Println(string(i))
				i++
				numCh <- true
			}
		}
	}()

	numCh <- true
	wg.Wait()
}
