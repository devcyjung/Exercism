const combination = (list, count) => {
    if (count < 0)
        throw new Error(`count must be non-negative but is: $count`)
    switch (count) {
        case 0:
            return []
        case 1:
            return list.map(elem => [elem])
        default:
            return list.flatMap((elem, index) => combination(list.slice(index + 1), count - 1).map(com => [elem, ...com]))
    }
}

export const combinations = cage => combination(Array.from({length: 9}, (_, i) => i + 1), cage.size)
    .filter(list => list.reduce((a, b) => a + b, 0) === cage.sum && list.every(elem => cage.exclude.indexOf(elem) === -1))
