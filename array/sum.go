package array

func Sum(numbers [5]int) int {
	sum := 0

	for index := 0; index < 5; index++ {
		sum = sum + numbers[index]
	}
	return sum
}
