from itertools import chain

def proverb(*args, qualifier = None):
    return [
        line for line in chain((
            f'For want of a {first} the {second} was lost.' for first, second in zip(
                args[:-1], args[1:])
        ), ('And all for the want of a ' + ' '.join(
            filter(lambda e: e, (qualifier, args[0]))
        ) + '.',))
    ] if args else []