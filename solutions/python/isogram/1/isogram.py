def is_isogram(string):
    letters = tuple(filter(lambda c: not c.isspace() and c != '-', string.lower()))
    return len(set(letters)) == len(letters)