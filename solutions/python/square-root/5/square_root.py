from sys import float_info
from typing import Final

_EPSILON: Final[float] = float_info.epsilon * 2

def square_root(number: float) -> float:
    guess: float = number
    result: float = guess * guess
    while abs(number - result) > _EPSILON:
        guess += (number - result) / (guess * 2)
        result = guess * guess
    return guess