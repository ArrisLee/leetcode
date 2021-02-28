package twosums

func twoSum(nums []int, target int) []int {
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
