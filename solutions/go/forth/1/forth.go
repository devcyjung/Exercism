package forth

// Checkout forth playground at: https://skilldrick.github.io/easyforth/#adding-some-numbers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Operation = func([]int) ([]int, error)
type OperationMap map[string]Operation

type InstructionKind string

const (
	Definition InstructionKind = "Definition"
	Execution  InstructionKind = "Execution"
)

type ErrorKind string

const (
	StackUnderflow       ErrorKind = "Stack Underflow"
	UnsupportedOperation ErrorKind = "Unsupported Operation"
	DivisionbyZero       ErrorKind = "Division by Zero"
	OperationError       ErrorKind = "Operation Error"
	SyntaxError          ErrorKind = "Syntax Error"
)

type Error struct {
	Kind  ErrorKind
	Trace string
}

func (e Error) Error() string {
	return fmt.Sprintf("Error: %s\nTraceBack: %s\n", e.Kind, e.Trace)
}

var plus = func(stack []int) (r []int, e error) {
	size := len(stack)
	if size < 2 {
		e = Error{
			Kind:  StackUnderflow,
			Trace: fmt.Sprintf("plus(%v)", stack),
		}
		return
	}
	stack = append(stack[:size-2], stack[size-2]+stack[size-1])
	r = stack
	return
}

var minus = func(stack []int) (r []int, e error) {
	size := len(stack)
	if size < 2 {
		e = Error{
			Kind:  StackUnderflow,
			Trace: fmt.Sprintf("minus(%v)", stack),
		}
		return
	}
	stack = append(stack[:size-2], stack[size-2]-stack[size-1])
	r = stack
	return
}

var mul = func(stack []int) (r []int, e error) {
	size := len(stack)
	if size < 2 {
		e = Error{
			Kind:  StackUnderflow,
			Trace: fmt.Sprintf("mul(%v)", stack),
		}
		return
	}
	stack = append(stack[:size-2], stack[size-2]*stack[size-1])
	r = stack
	return
}

var div = func(stack []int) (r []int, e error) {
	size := len(stack)
	if size < 2 {
		e = Error{
			Kind:  StackUnderflow,
			Trace: fmt.Sprintf("div(%v)", stack),
		}
		return
	}
	if stack[size-1] == 0 {
		e = Error{
			Kind:  DivisionbyZero,
			Trace: fmt.Sprintf("div(%v)", stack),
		}
		return
	}
	stack = append(stack[:size-2], stack[size-2]/stack[size-1])
	r = stack
	return
}

var dup = func(stack []int) (r []int, e error) {
	size := len(stack)
	if size < 1 {
		e = Error{
			Kind:  StackUnderflow,
			Trace: fmt.Sprintf("dup(%v)", stack),
		}
		return
	}
	stack = append(stack, stack[size-1])
	r = stack
	return
}

var drop = func(stack []int) (r []int, e error) {
	size := len(stack)
	if size < 1 {
		e = Error{
			Kind:  StackUnderflow,
			Trace: fmt.Sprintf("drop(%v)\n", stack),
		}
		return
	}
	stack = stack[:size-1]
	r = stack
	return
}

var swap = func(stack []int) (r []int, e error) {
	size := len(stack)
	if size < 2 {
		e = Error{
			Kind:  StackUnderflow,
			Trace: fmt.Sprintf("swap(%v)", stack),
		}
		return
	}
	stack = append(stack[:size-2], stack[size-1], stack[size-2])
	r = stack
	return
}

var over = func(stack []int) (r []int, e error) {
	size := len(stack)
	if size < 2 {
		e = Error{
			Kind:  StackUnderflow,
			Trace: fmt.Sprintf("over(%v)", stack),
		}
		return
	}
	stack = append(stack, stack[size-2])
	r = stack
	return
}

func GetBaseMap() OperationMap {
	return OperationMap{
		"+":    plus,
		"-":    minus,
		"*":    mul,
		"/":    div,
		"dup":  dup,
		"drop": drop,
		"swap": swap,
		"over": over,
	}
}

func ParseInput(input string) (s struct {
	kind  InstructionKind
	steps []string
	name  string
}) {
	fields := strings.Fields(strings.ToLower(input))
	if len(fields) >= 4 && fields[0] == ":" && fields[len(fields)-1] == ";" {
		s.kind = Definition
		s.name = fields[1]
		s.steps = fields[2 : len(fields)-1]
		return
	}
	s.kind = Execution
	s.steps = fields
	return
}

func (m OperationMap) Execute(stack []int, steps []string) (r []int, e error) {
	for _, step := range steps {
		n, err := strconv.Atoi(step)
		if err != nil {
			f, ok := m[step]
			if !ok {
				e = Error{
					Kind:  UnsupportedOperation,
					Trace: fmt.Sprintf("Execute(%v)\nOperationMap: %v\nOperation: %s", steps, m, step),
				}
				return
			}
			stack, err = f(stack)
			if err != nil {
				e = errors.Join(err, Error{
					Kind:  OperationError,
					Trace: fmt.Sprintf("Execute(%v)\nOperationMap: %v\nOperation: %s", steps, m, step),
				})
				return
			}
		} else {
			stack = append(stack, n)
		}
	}
	r = stack
	return
}

func (m OperationMap) New(steps []string) Operation {
	copied := make(OperationMap)
	for k, v := range m {
		copied[k] = v
	}
	return func(stack []int) (r []int, e error) {
		return copied.Execute(stack, steps)
	}
}

func Forth(inputs []string) (r []int, e error) {
	fmt.Println(inputs)
	var stack []int
	m := GetBaseMap()
	for _, input := range inputs {
		switch s := ParseInput(input); s.kind {
		case Definition:
			_, err := strconv.Atoi(s.name)
			if err == nil {
				e = errors.Join(e, Error{
					Kind:  UnsupportedOperation,
					Trace: fmt.Sprintf("Forth(%v) Number cannot be redefined as a command %s", inputs, s.name),
				})
			}
			m[s.name] = m.New(s.steps)
		case Execution:
			stack, e = m.Execute(stack, s.steps)
			if e != nil {
				e = errors.Join(e, Error{
					Kind:  OperationError,
					Trace: fmt.Sprintf("Forth(%v)", inputs),
				})
				return
			}
		default:
			e = Error{
				Kind:  SyntaxError,
				Trace: fmt.Sprintf("Forth(%v)", inputs),
			}
			return
		}
	}
	r = stack
	return
}
