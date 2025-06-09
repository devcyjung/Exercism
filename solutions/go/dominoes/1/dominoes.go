package dominoes

import "slices"

type Domino [2]int

func MakeChain(input []Domino) ([]Domino, bool) {
    if len(input) == 0 {
        return []Domino{}, true
    }
	rest := input
    chains := make([][]Domino, 0, len(input))
    for len(rest) > 0 {
        var ok bool
        var chain []Domino
        chain, rest, ok = getChain(rest)
        if !ok {
            return nil, false
        }
        chains = append(chains, chain)
    }
    result := chains[0]
    pruned := make([]bool, len(chains))
    pruned[0] = true
    prunedCnt := 1
    for prunedCnt < len(chains) {
        found := false
        for chainIdx, chain := range chains {
            if pruned[chainIdx] {
                continue
            }
            joint, ok := findJointChain(result, chain)
            if !ok {
                continue
            }
            found = true
            pruned[chainIdx] = true
            prunedCnt++
            result = joint
        }
        if !found {
            return nil, false
        }
    }
    return result, true
}

func findJointChain(chain1, chain2 []Domino) ([]Domino, bool) {
	for i1, domino1 := range chain1 {
        i2 := slices.IndexFunc(chain2, func(domino2 Domino) bool {
            return domino1[0] == domino2[0]
        })
        if i2 != -1 {
            return slices.Concat(chain1[:i1], chain2[i2:], chain2[:i2], chain1[i1:]), true
        }
    }
    return nil, false
}

func getChain(input []Domino) (chain []Domino, rest []Domino, ok bool) {
    startNum := input[0][0]
    chain = make([]Domino, 0, len(input) / 2)
    rest = make([]Domino, 0, len(input) / 2)
    pruned := make([]bool, len(input))
    chain = append(chain, input[0])
    pruned[0] = true
    for chain[len(chain) - 1][1] != startNum {
        cur := chain[len(chain) - 1][1]
        found := false
        for nextIdx, domino := range input {
            if pruned[nextIdx] {
                continue
            }
            if domino[0] == cur {
                chain = append(chain, domino)
                pruned[nextIdx] = true
                found = true
                break
            } else if domino[1] == cur {
                chain = append(chain, Domino{domino[1], domino[0]})
                pruned[nextIdx] = true
                found = true
                break
            }
        }
        if !found {
            ok = false
            return
        }
    }
    for i, domino := range input {
        if !pruned[i] {
            rest = append(rest, domino)
        }
    }
    ok = true
    return
}