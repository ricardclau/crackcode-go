package recursion

import (
	"fmt"
)

func PrintParenthesis(count int) {
	RecursivePrintParenthesis(count, count, "")
}

func RecursivePrintParenthesis(left int, right int, solution string) {
	//fmt.Printf("%v %v %v", left, right, solution)
	if left < 0 || right < left {
		panic("Something wrong happened")
	}

	if left == 0 && right == 0 {
		fmt.Printf("%v", solution)
		fmt.Println()
	}

	if left > 0 {
		RecursivePrintParenthesis(left-1, right, solution+"(")
	}

	if left < right {
		RecursivePrintParenthesis(left, right-1, solution+")")
	}
}
