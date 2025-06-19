from math import isqrt

def classify(number):
    """ A perfect number equals the sum of its positive divisors.

    :param number: int a positive integer
    :return: str the classification of the input integer
    """
    if number <= 0:
        raise ValueError("Classification is only possible for positive integers.")
    match sum((
        0 if divisor == number else
            divisor if dividend == number or dividend == divisor else
            divisor + dividend
        for divisor, dividend in (
            (divisor, number / divisor)
            for divisor in range(1, isqrt(number) + 1)
            if number % divisor == 0
        )
    )):
        case aliquot if aliquot > number:
            print(aliquot)
            return "abundant"
        case aliquot if aliquot == number:
            print(aliquot)
            return "perfect"
        case aliquot if aliquot < number:
            print(aliquot)
            return "deficient"