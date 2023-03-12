package main

func main() {
	plusOne([]int{9})

}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i:][0] != 9 {
			digits[i:][0] += 1
			return digits
		}
		digits[i:][0] = 0
		if i == 0 {
			digits = append([]int{1}, digits...)
			return digits
		}
	}
	return digits
}
