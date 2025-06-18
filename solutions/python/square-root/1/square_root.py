def square_root(number: float) -> float:
    guess: float = number
    result: float = guess * guess
    while abs(number - result) > _EPSILON:
        guess += (number - result) / (guess * 2)
        result = guess * guess
    return guess

_EPSILON: float = 1e-32