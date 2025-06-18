from sys import float_info

def square_root(number: float) -> float:
    number = abs(number)
    guess: float = number
    result: float = guess * guess
    while abs(number - result) > float_info.epsilon * number:
        guess += (number - result) / (guess * 2)
        result = guess * guess
    return guess