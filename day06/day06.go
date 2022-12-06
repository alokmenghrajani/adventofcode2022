package day06

func Part1(input string) int {
	return common(input, 4)
}

func Part2(input string) int {
	return common(input, 14)
}

func common(input string, length int) int {
	for i := length - 1; i < len(input); i++ {
		m := map[byte]int{}
		for j := 0; j < length; j++ {
			m[input[i-j]]++
		}
		if len(m) == length {
			return i + 1
		}
	}
	panic("failed to find a solution")
}
