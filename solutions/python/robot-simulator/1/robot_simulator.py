EAST = (1, 0)
NORTH = (0, 1)
WEST = (-1, 0)
SOUTH = (0, -1)

class Robot:
    _DIRECTIONS = (NORTH, EAST, SOUTH, WEST)
    
    def __init__(self, direction=NORTH, x_pos=0, y_pos=0):
        self.direction = direction
        self.coordinates = (x_pos, y_pos)

    def move(self, instructions):
        for instruction in instructions:
            match instruction:
                case "A":
                    self.coordinates = tuple((
                        x + dx
                        for x, dx in
                        zip(self.coordinates, self.direction)
                    ))
                case "L":
                    self.direction = Robot._DIRECTIONS[
                        (Robot._DIRECTIONS.index(self.direction) - 1) % 4
                    ]
                case "R":
                    self.direction = Robot._DIRECTIONS[
                        (Robot._DIRECTIONS.index(self.direction) + 1) % 4
                    ]
                case _:
                    raise ValueError(f"Invalid command: {instruction}")