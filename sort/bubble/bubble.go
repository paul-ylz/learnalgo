package bubble

func bubbleSort(args []int) []int {
	list := make([]int, len(args))
	copy(list, args)

	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}
	}
	return list
}
