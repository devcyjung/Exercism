from itertools import zip_longest

def transpose(text):
    return '\n'.join((
        ''.join(column).rstrip(chr(0x1dead)) for column in zip_longest(
            *text.splitlines(), fillvalue = chr(0x1dead)
        ))).replace(chr(0x1dead), ' ')