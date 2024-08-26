package deduplicate

func removeDuplicates(nums []int) int {

	// If we have only one number it is
	// always unique
	if len(nums) == 1 {
		return 1
	}

	unique := 1
	for n := 1; n < len(nums); n++ {
		// If this number is larger than the previous one
		// we found a new unique value
		if nums[n-1] < nums[n] {
			// Copy the number to the first non-unique position
			nums[unique] = nums[n]

			// Count the new unique number
			unique++
		} else if nums[n-1] > nums[n] {
			// Stop once we find a smaller number
			// than the pervious one
			break
		}
	}

	return unique
}
