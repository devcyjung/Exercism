package twobucket

import (
    "errors"
    "slices"
)

var (
    ErrInvalidInput = errors.New("Invalid input format")
    ErrImpossibleGoal = errors.New("Goal is unreachable")
)

type gameState struct {
    pair	[2]int
    move	int
}

func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (
    string, int, int, error,
) {
    if sizeBucketOne <= 0 || sizeBucketTwo <= 0 || goalAmount <= 0 {
        return "", 0, 0, ErrInvalidInput
    }
	var src int
    switch startBucket {
    case "one":
        src = 0
    case "two":
        src = 1
    default:
        return "", 0, 0, ErrInvalidInput
    }
    gen := nextStateGenerator(src, [2]int{sizeBucketOne, sizeBucketTwo})
    for {
        gs, ok := gen()
        if !ok {
            break
        }
        if gs.pair[0] == goalAmount {
            return "one", gs.move, gs.pair[1], nil
        }
        if gs.pair[1] == goalAmount {
            return "two", gs.move, gs.pair[0], nil
        }
    }
    return "", 0, 0, ErrImpossibleGoal
}

func nextStateGenerator(src int, sizes [2]int) func() (gameState, bool) {
    queue := make([]gameState, 0, 128)
    visited := make(map[[2]int]struct{})
    dst := 1 - src
    initialPair := [2]int{}
    initialPair[src] = sizes[src]
    queue = append(queue, gameState{initialPair, 1})
    isInvalid := func(gs gameState) bool {
        pair := gs.pair
		return pair[0] < 0 || pair[1] < 0 || pair[0] > sizes[0] || pair[1] > sizes[1] ||
        	(pair[src] == 0 && pair[dst] == sizes[dst])
    }
    nextCandidates := func(gs gameState) []gameState {
        move := gs.move + 1
        a, b := gs.pair[0], gs.pair[1]
        A, B := sizes[0], sizes[1]
        template := []gameState{
            {[2]int{0, a + b}, move}, {[2]int{a + b - B, B}, move}, {[2]int{a + b, 0}, move},
            {[2]int{A, a + b - A}, move}, {[2]int{0, b}, move}, {[2]int{a, 0}, move},
            {[2]int{A, b}, move}, {[2]int{a, B}, move},
        }
        return slices.DeleteFunc(template, func(g gameState) bool {
            _, exists := visited[g.pair]
            if exists || isInvalid(g) {
                return true
            }
            visited[g.pair] = struct{}{}
            return false
        })
    }
    return func() (g gameState, ok bool) {
        if len(queue) == 0 {
            return
        }
        g, ok = queue[0], true
        queue = append(queue[1:], nextCandidates(g)...)
        return
    }
}