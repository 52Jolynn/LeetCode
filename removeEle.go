package main

func main() {
	println(removeElement([]int{3, 2, 2, 3}, 3))
	println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement(nums []int, val int) int {
	index := -1
	for i := 0; i < len(nums); {
		if nums[i] != val {
			index += 1
			nums[index] = nums[i]
		}
		i += 1
	}

	return index + 1
}
