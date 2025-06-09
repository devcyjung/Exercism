package robotname

import (
    "errors"
    "math/rand/v2"
    "sync"
)

const MAX_UUID_COUNT int = 26 * 26 * 10 * 10 * 10

var (
    once sync.Once
    idTable []string
    NoAvailableNamesError = errors.New("All IDs are in use")
)

type Robot struct{
    name	string
}

func (r *Robot) Name() (string, error) {
    if r.name == "" {
        err := retrieveName(r)
    	return r.name, err
    }
    return r.name, nil
}

func (r *Robot) Reset() {
    recycleName(r)
}

func mapID(n int) string {
    letters := make([]rune, 5)
    for i := 0; i < 5; i++ {
        if i < 3 {
            letters[4 - i] = '0' + rune(n % 10)
            n /= 10
        } else {
            letters[4 - i] = 'A' + rune(n % 26)
            n /= 26
        }
    }
    return string(letters)
}

func shuffleNames() {
    rand.Shuffle(len(idTable), func(i, j int) {
        idTable[i], idTable[j] = idTable[j], idTable[i]
    })
}

func initialize() {
    idTable = make([]string, MAX_UUID_COUNT, MAX_UUID_COUNT)
    for i := range idTable {
        idTable[i] = mapID(i)
    }
    shuffleNames()
}

func retrieveName(r *Robot) error {
    once.Do(initialize)
    if len(idTable) == 0 {
        return NoAvailableNamesError
    }
    r.name = idTable[len(idTable) - 1]
    idTable = idTable[:len(idTable) - 1]
    return nil
}

func recycleName(r *Robot) {
    idTable = append(idTable, r.name)
    r.name = ""
    shuffleNames()
}