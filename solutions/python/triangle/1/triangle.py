def equilateral(sides):
    if not triangle(sides):
        return False
    return sides[0] == sides[2]

def isosceles(sides):
    if not triangle(sides):
        return False
    return sides[0] == sides[1] or sides[1] == sides[2]

def scalene(sides):
    if not triangle(sides):
        return False
    return sides[0] != sides[1] and sides[1] != sides[2]

def triangle(sides):
    sides.sort()
    return sides[2] < sides[0] + sides[1]