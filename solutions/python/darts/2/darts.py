from math import hypot
from typing import Final

_DISTANCE_SCORE: Final[tuple[tuple[int, int], ...]] = ((1, 10), (5, 5), (10, 1))

def score(x: int, y: int) -> int:
    radius: Final[float] = hypot(x, y)
    return next((score for distance, score in _DISTANCE_SCORE if radius <= distance), 0)