c := make(chan int)
quit := make(chan bool)
for i := 0; i < 10; i++ {
	c <- i
}
quit <- true

for {
	select {
	case s := <-c:
		fmt.Println(s)
	case <-quit:
		fmt.Println("Job is done")
		return
	}
}