package stack

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			//数字入栈
			stack = append(stack, val)
		}
		if err != nil {
			x := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			y := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			z := cal(y, x, token)
			stack = append(stack, z)
		}
	}
	return stack[len(stack)-1]
}

func cal(a int, b int, op string) int {
	if op == "+" {
		return a + b
	}
	if op == "-" {
		return a - b
	}
	if op == "*" {
		return a * b
	}
	if op == "/" {
		return a / b
	}
	return 0
}

func evalRPN1(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else {
			y := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			x := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			z := cal(x, y, token)
			stack = append(stack, z)
		}
	}
	return stack[len(stack)-1]
}

func cal1(x int, y int, op string) int {
	if op == "+" {
		return x + y
	}
	if op == "-" {
		return x - y
	}
	if op == "*" {
		return x * y
	}
	if op == "/" {
		return x / y
	}
	return 0
}
