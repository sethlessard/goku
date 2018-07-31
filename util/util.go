package util

import "sync"

// MergeGoroutines merges many goroutines into one
func MergeGoroutines(done <-chan bool, channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup

	wg.Add(len(channels))
	outgoingPackages := make(chan interface{})
	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case outgoingPackages <- i:
			}
		}
	}
	for _, c := range channels {
		go multiplex(c)
	}
	go func() {
		wg.Wait()
		close(outgoingPackages)
	}()
	return outgoingPackages
}
