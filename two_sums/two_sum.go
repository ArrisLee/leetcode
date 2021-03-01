package twosums

func twoSumSltnOne(nums []int, target int) []int {
	for i := range nums {
		for j := range nums {
			if j == i {
				continue
			}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumSltnTwo(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i := range nums {
		if val, ok := numMap[target-nums[i]]; ok {
			return []int{val, i}
		}
		numMap[nums[i]] = i
	}
	return nil
}
