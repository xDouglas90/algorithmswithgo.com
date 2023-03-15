package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) (total int) {
	total = 0
	for _, v := range numbers {
		total += v
	}
	return
}
