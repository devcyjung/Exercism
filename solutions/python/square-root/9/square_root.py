from math import ulp
from typing import Final

def square_root(input: float) -> float:
    number: float = abs(input)
    epsilon: Final[float] = ulp(number)
    guess: float = number
    result: float = guess * guess
    while abs(number - result) > epsilon:
        guess += (number - result) / (guess * 2)
        result = guess * guess
    return guess