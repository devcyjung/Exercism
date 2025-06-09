def is_armstrong_number(number):
    if number < 0:
        raise ValueError("Number out of range")
    s = str(number)
    pow = len(s)
    acc = 0
    for ch in s:
        acc += int(ch) ** pow
    return acc == number