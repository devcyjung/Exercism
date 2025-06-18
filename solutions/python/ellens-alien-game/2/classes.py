"""Solution to Ellen's Alien Game exercise."""

from typing import ClassVar, Iterable, List, Tuple

class Alien:
    """Create an Alien object with location x_coordinate and y_coordinate.

    Attributes
    ----------
    (class)total_aliens_created: int
    x_coordinate: int - Position on the x-axis.
    y_coordinate: int - Position on the y-axis.
    health: int - Number of health points.

    Methods
    -------
    hit(): Decrement Alien health by one point.
    is_alive(): Return a boolean for if Alien is alive (if health is > 0).
    teleport(new_x_coordinate, new_y_coordinate): Move Alien object to new coordinates.
    collision_detection(other): Implementation TBD.
    """

    total_aliens_created: ClassVar[int] = 0

    def __init__(self: 'Alien', x_coordinate: int, y_coordinate: int) -> None:
        self.x_coordinate: int = x_coordinate
        self.y_coordinate: int = y_coordinate
        self.health: int = 3
        Alien.total_aliens_created += 1
    
    def hit(self: 'Alien') -> None:
        self.health -= 1

    def is_alive(self: 'Alien') -> bool:
        return self.health > 0

    def teleport(self: 'Alien', new_x_coordinate: int, new_y_coordinate: int) -> None:
        self.x_coordinate = new_x_coordinate
        self.y_coordinate = new_y_coordinate

    def collision_detection(self: 'Alien', other: 'Alien') -> None:
        pass


def new_aliens_collection(coordinates: Iterable[Tuple[int, int]]) -> List[Alien]:
    return [Alien(*coordinate) for coordinate in coordinates]
    