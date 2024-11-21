ch := make(chan int)

total := 0

var wg sync.WaitGroup
for i := 0; i < len(data); i += part_length {
	wg.Add(1)
	go func(l int, r int) {
		defer wg.Done()
		partial_sum := 0
		for i := l; i < min(r, len(data)); i++ {
			partial_sum += data[i]
		}
		ch <- partial_sum
	}(i, i+part_length)
}

var wg_total sync.WaitGroup
wg_total.Add(1)
go func() {
	defer wg_total.Done()
	for ps := range ch {
		total += ps
	}
}()

wg.Wait()
close(ch)
wg_total.Wait()
