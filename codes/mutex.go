var count int
var lock sync.Mutex
increment := func() {
	lock.Lock()
	defer lock.Unlock()
	count++
}
decrement := func() {
	lock.Lock()
	defer lock.Unlock()
	count--
}
var arithmetic sync.WaitGroup
for i := 0; i <= 5; i++ { // Increment
	arithmetic.Add(1)
	go func() {
		defer arithmetic.Done()
		increment()
	}()
}
for i := 0; i <= 5; i++ { // Decrement
	arithmetic.Add(1)
	go func() {
		defer arithmetic.Done()
		decrement()
	}()
}
arithmetic.Wait()