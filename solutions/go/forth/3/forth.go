package forth

import (
	"errors"
	"strconv"
	"strings"
)

func Forth(input []string) (stack []int, err error) {
	inter := newForthInterpreter()
	for _, line := range input {
		for _, token := range strings.Fields(strings.ToUpper(line)) {
			if inter.state, err = inter.state.dispatch(token); err != nil {
				return
			}
		}
	}
	stack = inter.stack
	return
}

type forthInterpreter struct {
	stack       []int
	state       forthState
	definitions map[string][]string
}

func newForthInterpreter() *forthInterpreter {
	inter := &forthInterpreter{
		stack:       make([]int, 0),
		definitions: make(map[string][]string),
	}
	inter.state = &interpretState{
		forthInterpreter: inter,
	}
	return inter
}

func (inter *forthInterpreter) interpret(token string) error {
	var err error
	if intToken, err := strconv.Atoi(token); err == nil {
		inter.stack = append(inter.stack, intToken)
		return nil
	}
	if instructions, ok := inter.definitions[token]; ok {
		for _, instruction := range instructions {
			if err = inter.interpret(instruction); err != nil {
				return err
			}
		}
		return nil
	}
	if builtin, ok := builtins[token]; ok {
		if inter.stack, err = builtin(inter.stack); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("undefined operation")
	}
}

type forthState interface {
	dispatch(token string) (forthState, error)
}
type newWordState struct {
	*forthInterpreter
}

func (state *newWordState) dispatch(token string) (forthState, error) {
	if _, err := strconv.Atoi(token); err == nil {
		return state, errors.New("illegal operation")
	}
	return &defineWordState{
		forthInterpreter: state.forthInterpreter,
		word:             token,
	}, nil
}

type defineWordState struct {
	*forthInterpreter
	word         string
	instructions []string
}

func (state *defineWordState) isDefined(token string) bool {
	if _, ok := state.definitions[token]; ok {
		return true
	}
	if _, ok := builtins[token]; ok {
		return true
	}
	_, err := strconv.Atoi(token)
	return err == nil
}

func (state *defineWordState) dispatch(token string) (forthState, error) {
	if token == ";" {
		state.definitions[state.word] = state.instructions
		return &interpretState{
			forthInterpreter: state.forthInterpreter,
		}, nil
	}
	if !state.isDefined(token) {
		return state, errors.New("undefined operation")
	}
	if predefined, ok := state.definitions[token]; ok {
		state.instructions = append(state.instructions, predefined...)
	} else {
		state.instructions = append(state.instructions, token)
	}
	return state, nil
}

type interpretState struct {
	*forthInterpreter
}

func (state *interpretState) dispatch(token string) (forthState, error) {
	if token == ":" {
		return &newWordState{
			forthInterpreter: state.forthInterpreter,
		}, nil
	}
	err := state.interpret(token)
	return state, err
}

var builtins = map[string]func([]int) ([]int, error){
	"+":    plus,
	"-":    minus,
	"*":    multiply,
	"/":    divide,
	"SWAP": swap,
	"DROP": drop,
	"DUP":  dup,
	"OVER": over,
}

func plus(stack []int) ([]int, error) {
	return binaryOperation(stack, func(a, b int) []int {
		return []int{a + b}
	})
}

func minus(stack []int) ([]int, error) {
	return binaryOperation(stack, func(a, b int) []int {
		return []int{a - b}
	})
}

func multiply(stack []int) ([]int, error) {
	return binaryOperation(stack, func(a, b int) []int {
		return []int{a * b}
	})
}

func divide(stack []int) ([]int, error) {
	if len(stack) > 0 && stack[len(stack)-1] == 0 {
		return stack, errors.New("divide by zero")
	}
	return binaryOperation(stack, func(a, b int) []int {
		return []int{a / b}
	})
}

func swap(stack []int) ([]int, error) {
	return binaryOperation(stack, func(a, b int) []int {
		return []int{b, a}
	})
}

func drop(stack []int) ([]int, error) {
	return unaryOperation(stack, func(a int) []int {
		return []int{}
	})
}

func dup(stack []int) ([]int, error) {
	return unaryOperation(stack, func(a int) []int {
		return []int{a, a}
	})
}

func over(stack []int) ([]int, error) {
	return binaryOperation(stack, func(a, b int) []int {
		return []int{a, b, a}
	})
}

func unaryOperation(stack []int, op func(int) []int) ([]int, error) {
	if len(stack) < 1 {
		return stack, errors.New("empty stack")
	}
	result := op(stack[len(stack)-1])
	stack = stack[:len(stack)-1]
	return append(stack, result...), nil
}

func binaryOperation(stack []int, op func(int, int) []int) ([]int, error) {
	if len(stack) < 1 {
		return stack, errors.New("empty stack")
	}
	if len(stack) < 2 {
		return stack, errors.New("only one value on the stack")
	}
	result := op(stack[len(stack)-2], stack[len(stack)-1])
	stack = stack[:len(stack)-2]
	return append(stack, result...), nil
}
