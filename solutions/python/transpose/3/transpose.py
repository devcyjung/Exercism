from itertools import zip_longest

def transpose(text):
    return '\n'.join((
        ''.join(column).rstrip(chr(0xdead)) for column in zip_longest(
            *text.splitlines(), fillvalue = chr(0xdead)
        ))).replace(chr(0xdead), ' ')