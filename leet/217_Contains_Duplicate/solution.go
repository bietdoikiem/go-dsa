package containsduplicate

// containsDuplicate check if there are duplicated values in the given array
// Time complexity: O(n)
// Space complexity: O(n) - Auxiliary Map data structure
func containsDuplicate(nums []int) bool {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if numMap[nums[i]] > 0 {
			return true
		}
		numMap[nums[i]] += 1
	}
	return false
}
