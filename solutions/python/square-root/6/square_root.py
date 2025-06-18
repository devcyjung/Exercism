from sys import float_info
from typing import Final

_EPSILON: Final[float] = float_info.epsilon

def square_root(number: float) -> float:
    number = abs(number)
    guess: float = number
    result: float = guess * guess
    while abs(number - result) > _EPSILON * number:
        guess += (number - result) / (guess * 2)
        result = guess * guess
    return guess