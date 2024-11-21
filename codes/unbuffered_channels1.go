func countSum(data []int, int l, int r, wg *sync.WaitGroup) <- chan int {
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		partial_sum := 0
		for i := l; i < min(r, len(data)); i++ {
			partial_sum += data[i]
		}
		ch <- partial_sum
	}()
	return ch
}
sumChan := make(chan int)
total := 0

var wg sync.WaitGroup
for i := 0; i < len(data); i += part_length {
	sumChan <- <-countSum(data, i, i+part_length, wg)
}
go func() {
	defer wg_total.Done()
	for ps := range sumChan {
		total += ps
	}
}()