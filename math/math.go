package math

func Sum(numbers []int) int {
	s := 0
	for _, v := range numbers {
		s += v
	}
	return s
}

func Add(a, b int) int {
	return a + b
}
