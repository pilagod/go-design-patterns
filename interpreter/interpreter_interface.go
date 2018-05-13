package interpreter

import (
	"strconv"
	"strings"
)

func InterpreterCalculate(operation string) (int, error) {
	stack := interpreterPolishNotationStack{}
	operators := strings.Split(operation, " ")

	for _, operand := range operators {
		if isOperator(operand) {
			right := stack.Pop()
			left := stack.Pop()
			calculate := operatorFactory(operand, left, right)
			res := value(calculate.Read())
			stack.Push(&res)
		} else {
			val, err := strconv.Atoi(operand)

			if err != nil {
				return 0, err
			}
			temp := value(val)
			stack.Push(&temp)
		}
	}
	return stack.Pop().Read(), nil
}

type interpreterPolishNotationStack []Interpreter

func (ipns *interpreterPolishNotationStack) Push(i Interpreter) {
	*ipns = append(*ipns, i)
}

func (ipns *interpreterPolishNotationStack) Pop() Interpreter {
	length := len(*ipns)

	if length > 0 {
		temp := (*ipns)[length-1]
		*ipns = (*ipns)[:length-1]
		return temp
	}
	return nil
}

type Interpreter interface {
	Read() int
}

type value int

func (v *value) Read() int {
	return int(*v)
}

type operation struct {
	Left  Interpreter
	Right Interpreter
}

type operationSum struct {
	operation
}

func (os *operationSum) Read() int {
	return os.Left.Read() + os.Right.Read()
}

type operationSubtract struct {
	operation
}

func (os *operationSubtract) Read() int {
	return os.Left.Read() - os.Right.Read()
}

type operationMultiply struct {
	operation
}

func (om *operationMultiply) Read() int {
	return om.Left.Read() * om.Right.Read()
}

type operationDivide struct {
	operation
}

func (od *operationDivide) Read() int {
	return od.Left.Read() / od.Right.Read()
}

func operatorFactory(operand string, left, right Interpreter) Interpreter {
	operation := operation{left, right}
	switch operand {
	case SUM:
		// sum := &operationSum{}
		// sum.Left = left
		// sum.Right = right
		// or
		// sum := &operationSum{operation{left, right}}
		return &operationSum{operation}
	case SUB:
		return &operationSubtract{operation}
	case MUL:
		return &operationMultiply{operation}
	case DIV:
		return &operationDivide{operation}
	}
	return nil
}
