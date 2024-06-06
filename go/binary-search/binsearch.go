package binsearch

func BinSearch(list []int, target int) int {
	var left int = 0
	var right int = len(list) - 1
	var mid int

	for left <= right {
		mid = int((left + right) / 2)
		if list[mid] < target {
			left = mid + 1
		} else if list[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
