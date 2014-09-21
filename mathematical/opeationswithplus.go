package mathematical

func Multiplication(x int, y int) int {
	if x < y {
		return Multiplication(y, x)
	}

	result := 0
	for i := 0; i < Absolute(y); i++ {
		result += x
	}

	/*
		if y < 0 {
			return
		}
	*/
	return result
}

func Negate(x int) int {

	addition := 1
	if x < 0 {
		addition = -1
	}

	result := 0
	for i := 0; i < x; i++ {
		result += addition
	}

	return result
}

func Absolute(x int) int {
	if x < 0 {
		return Negate(x)
	}

	return x
}
