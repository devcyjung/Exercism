def is_paired(input_string):
    opening, closing, stack = '({[', ')}]', []
    for input in input_string:
        match (opening.find(input), closing.find(input)):
            case o, _ if o >= 0:
                stack.append(o)
            case _, c if c >= 0 and (not stack or stack.pop() != c):
                return False
    return not stack