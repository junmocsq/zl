package monitor

func Parse(done chan struct{}) <-chan *Collect {
	ch := make(chan *Collect)
	go func() {
		defer close(ch)
		for i := range collectChan {
			select {
			case ch <- i:
			case <-done:
				return
			}
		}
	}()
	return ch
}
