package interpreter

import (
	"strconv"
	"strings"
)

func Calculate(operation string) (int, error) {
	stack := polishNotationStack{}
	operators := strings.Split(operation, " ")

	for _, operand := range operators {
		if isOperator(operand) {
			right := stack.Pop()
			left := stack.Pop()
			calculate := getOperationFunc(operand)
			stack.Push(calculate(left, right))
		} else {
			value, err := strconv.Atoi(operand)

			if err != nil {
				return 0, err
			}
			stack.Push(value)
		}
	}
	return stack.Pop(), nil
}

func isOperator(operand string) bool {
	if operand == SUM || operand == SUB || operand == MUL || operand == DIV {
		return true
	}
	return false
}

func getOperationFunc(operand string) func(a, b int) int {
	switch operand {
	case SUM:
		return func(a, b int) int {
			return a + b
		}
	case SUB:
		return func(a, b int) int {
			return a - b
		}
	case MUL:
		return func(a, b int) int {
			return a * b
		}
	case DIV:
		return func(a, b int) int {
			return a / b
		}
	}
	return nil
}

const (
	SUM = "sum"
	SUB = "sub"
	MUL = "mul"
	DIV = "div"
)

type polishNotationStack []int

func (pns *polishNotationStack) Push(n int) {
	*pns = append(*pns, n)
}

func (pns *polishNotationStack) Pop() int {
	length := len(*pns)

	if length > 0 {
		temp := (*pns)[length-1]
		*pns = (*pns)[:length-1]
		return temp
	}
	return 0
}
