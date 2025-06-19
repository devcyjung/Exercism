from itertools import zip_longest

def transpose(text):
    return '\n'.join((
        ''.join(column).rstrip(chr(0xDEAD)) for column in zip_longest(
            *text.split('\n'), fillvalue = chr(0xDEAD)
        ))).replace(chr(0xDEAD), ' ')