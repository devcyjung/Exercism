from collections.abc import Iterable

def flatten(iterable):
    return [item for items in iterable for item in (
        flatten(items) if isinstance(items, Iterable) else (items,)
    ) if item is not None]