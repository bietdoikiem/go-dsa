package productofarrayexceptself

// productExceptSelf returns the arrays of multiplied result of all elements except nums[i]
// Time complexity: O(2n) => O(n)
// Space complexity: O(2n) (2 extra arrays excluding output array)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	outputArr := make([]int, n)
	prefixArr := make([]int, n)
	postfixArr := make([]int, n)
	// Initialize pre/post arrays for first/last element of those arrays
	prefixArr[0] = 1
	postfixArr[n-1] = 1
	// Calculate prefix arrray
	for i := 1; i < n; i++ {
		prefixArr[i] = nums[i-1] * prefixArr[i-1]
	}
	// Pre-compute first/last element of output array
	outputArr[n-1] = prefixArr[n-1] * postfixArr[n-1]
	// Calculate postfix array
	for i := n - 2; i >= 0; i-- {
		postfixArr[i] = nums[i+1] * postfixArr[i+1]
		outputArr[i] = postfixArr[i] * prefixArr[i] // Output array
	}
	return outputArr
}
