c := make(chan string)

for {
	select {
	case s := <-c:
		fmt.Println(s)
	case <-time.After(1 * time.Second):
		fmt.Println("You're too slow")
		return
	}
}