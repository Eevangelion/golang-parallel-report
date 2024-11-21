c := make(chan string)
timeout := time.After(5 * time.Second)

for {
	select {
	case s := <-c:
		fmt.Println(s)
	case <-timeout:
		fmt.Println("Time is over")
		return
	}
}