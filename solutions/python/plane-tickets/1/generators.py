"""Functions to automate Conda airlines ticketing system."""

from collections.abc import Iterable, Iterator, Sized
from itertools import count, cycle, islice
from typing import Protocol, TypeVar

T = TypeVar('T')

class SizedIterable(Protocol[T]):
    def __iter__(self) -> Iterable[T]: ...
    def __len__(self) -> int: ...

def generate_seat_letters(number: int) -> Iterator[str]:
    """Generate a series of letters for airline seats.

    :param number: int - total number of seat letters to be generated.
    :return: generator - generator that yields seat letters.

    Seat letters are generated from A to D.
    After D it should start again with A.

    Example: A, B, C, D

    """

    return (letter for letter in islice(cycle('ABCD'), number))


def generate_seats(number: int) -> Iterator[str]:
    """Generate a series of identifiers for airline seats.

    :param number: int - total number of seats to be generated.
    :return: generator - generator that yields seat numbers.

    A seat number consists of the row number and the seat letter.

    There is no row 13.
    Each row has 4 seats.

    Seats should be sorted from low to high.

    Example: 3C, 3D, 4A, 4B

    """

    return (f'{row}{seat}' for row, seat in islice(
        zip((row for row in count(1) for _ in range(4) if row != 13), cycle('ABCD')),
        number))


def assign_seats(passengers: SizedIterable[str]) -> dict[str, str]:
    """Assign seats to passengers.

    :param passengers: list[str] - a list of strings containing names of passengers.
    :return: dict - with the names of the passengers as keys and seat numbers as values.

    Example output: {"Adele": "1A", "Björk": "1B"}

    """

    return {name: seat for name, seat in zip(passengers, generate_seats(len(passengers)))}


def generate_codes(seat_numbers: Iterable[str], flight_id: str) -> Iterator[str]:
    """Generate codes for a ticket.

    :param seat_numbers: list[str] - list of seat numbers.
    :param flight_id: str - string containing the flight identifier.
    :return: generator - generator that yields 12 character long ticket codes.

    """

    return (f'{seat + flight_id:0<12}' for seat in seat_numbers)
