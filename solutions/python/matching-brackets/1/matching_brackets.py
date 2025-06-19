def is_paired(input_string):
    opening = ('(', '{', '[')
    closing = (')', '}', ']')
    stack = []
    for input in input_string:
        if input in closing:
            if not stack or closing.index(input) != stack.pop():
                return False
        if input in opening:
            stack.append(opening.index(input))
    return not stack