from math import hypot

score = lambda *xs: next((r for t, r in ((1, 10), (5, 5), (10, 1)) if hypot(*xs) <= t), 0)