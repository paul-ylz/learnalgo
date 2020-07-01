package main

func insertionSort(items []int) {
	// assume list[0] is "sorted"
	for i := 1; i < len(items); i++ {
		// go 1 back and work backwards and swap with every item that is bigger
		for j := i; j > 0; j-- {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
		}

	}
}
