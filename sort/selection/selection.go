package selection

func selectionSort(args []int) []int {
	list := make([]int, len(args))
	copy(list, args)

	for i := 0; i < len(list)-1; i++ {
		smallestIndex := i
		for j := i; j <= len(list)-1; j++ {
			if list[j] < list[smallestIndex] {
				smallestIndex = j
			}
		}
		list[i], list[smallestIndex] = list[smallestIndex], list[i]
	}
	return list
}

func selectionSortReverse(args []int) []int {
	list := make([]int, len(args))
	copy(list, args)

	for i := 0; i < len(list)-1; i++ {
		largestIndex := i
		for j := i; j <= len(list)-1; j++ {
			if list[j] > list[largestIndex] {
				largestIndex = j
			}
		}
		list[i], list[largestIndex] = list[largestIndex], list[i]
	}
	return list
}
