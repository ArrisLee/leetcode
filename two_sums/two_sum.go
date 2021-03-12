package twosums

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i := range nums {
		if val, ok := numMap[target-nums[i]]; ok {
			return []int{val, i}
		}
		numMap[nums[i]] = i
	}
	return nil
}
