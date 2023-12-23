package leetcode

import "strconv"

/*
150. Evaluate Reverse Polish Notation

You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the expression.

Note that:
* The valid operators are '+', '-', '*', and '/'.
* Each operand may be an integer or another expression.
* The division between two integers always truncates toward zero.
* There will not be any division by zero.
* The input represents a valid arithmetic expression in a reverse polish notation.
* The answer and all the intermediate calculations can be represented in a 32-bit integer.

Example 1:
Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9

Example 2:
Input: tokens = ["4","13","5","/","+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6

Example 3:
Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
Output: 22
Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22

Constraints:
* 1 <= tokens.length <= 10^4
* tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the range [-200, 200].
*/

func evaluateReversePolishNotation(tokens []string) int {
	type Operation func(int, int) int

	add := func(a1, a2 int) int { return a1 + a2 }
	subtract := func(a1, a2 int) int { return a1 - a2 }
	multiply := func(a1, a2 int) int { return a1 * a2 }
	divide := func(a1, a2 int) int { return a1 / a2 }

	var operations = map[string]Operation{
		"+": add,
		"-": subtract,
		"*": multiply,
		"/": divide,
	}

	var stack = new(Stack)

	for _, token := range tokens {
		if operation, ok := operations[token]; ok {
			a2, _ := stack.Pop()
			a1, _ := stack.Pop()
			stack.Push(operation(a1, a2))
		} else {
			v, _ := strconv.Atoi(token)
			stack.Push(v)
		}
	}

	res, _ := stack.Pop()
	return res
}

func evaluateReversePolishNotationV2(tokens []string) int {
	stack := make([]int, 0, 2)

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			operands := stack[len(stack)-2:]
			stack = stack[:len(stack)-2]
			var v int
			switch token {
			case "+":
				v = operands[0] + operands[1]
			case "-":
				v = operands[0] - operands[1]
			case "*":
				v = operands[0] * operands[1]
			case "/":
				v = operands[0] / operands[1]
			}
			stack = append(stack, v)
		} else {
			v, _ := strconv.Atoi(token)
			stack = append(stack, v)
		}
	}

	return stack[0]
}

func evaluateReversePolishNotationTopRated(tokens []string) int {
	stack := []int{}
	operators := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
	for _, token := range tokens {
		if calculate, exist := operators[token]; exist {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], calculate(a, b))
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}
