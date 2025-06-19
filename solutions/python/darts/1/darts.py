from math import hypot

_distance_score: tuple[tuple[int, int], ...] = ((1, 10), (5, 5), (10, 1))

def score(x: int, y: int) -> int:
    radius: Final[float] = hypot(x, y)
    return next((score for distance, score in _distance_score if radius <= distance), 0)