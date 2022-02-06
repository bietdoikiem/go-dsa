package twosum

// twoSum returns the indices of two elements whose sum is equal to target
// Time complexity: O(n)
// Space complexity: O(n) (for auxiliary Map)
func twoSum(nums []int, target int) []int {
	records := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		if v, ok := records[target-nums[i]]; ok {
			return []int{i, v}
		}
		records[nums[i]] = i
	}
	return []int{}
}
