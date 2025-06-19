def steps(number):
    if number <= 0:
        raise ValueError("Only positive integers are allowed")
    step = 0
    while number != 1:
        match (number & -number).bit_length() - 1:
            case 0:
                number += (number << 1) + 1
                step += 1
            case trailing_zero:
                number >>= trailing_zero
                step += trailing_zero
    return step