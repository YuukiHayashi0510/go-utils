package unisort

import "sort"

// UniqueSortNaturalInts sorts a slice of natural integers and removes duplicates.
func UniqueSortNaturalInts(arr []int) []int {
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)
	sort.Ints(result)

	if result[0] < 0 {
		return result
	}

	maxVal := 0
	for _, v := range result {
		if v > maxVal {
			maxVal = v
		}
	}

	sorted := make([]int, maxVal+1)
	for _, v := range result {
		sorted[v] = v
	}

	uniqueList := make([]int, 0)
	for k, v := range sorted {
		if v > 0 {
			uniqueList = append(uniqueList, k)
		}
	}

	return uniqueList
}
