func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	parts := 4
	part_length := len(data) / parts

	for i := 0; i < len(data); i += part_length {
		go func(l int, r int) {
			for i := l; i < min(r, len(data)); i++ {
				data[i]++
			}
		}(i, i+part_length)
	}

	fmt.Println("New array: ", data)
}