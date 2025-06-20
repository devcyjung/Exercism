def sum_of_multiples(limit, multiples):
    return sum({
        multiple * i for multiple in multiples
        if multiple != 0
        for i in range(1, (limit - 1) // multiple + 1)
    })