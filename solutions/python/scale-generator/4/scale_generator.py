from collections.abc import Iterable
from itertools import accumulate
from functools import lru_cache
from typing import ClassVar, Literal, Union

class Scale:
    def __init__(self: 'Scale', tonic: str) -> None:
        if tonic in Scale.__sharps:
            self.__series: tuple[Note] = Scale.__sharp_series
            self.__base_idx: int = Scale.__sharp_series.index(tonic.title())
            return
        if tonic in Scale.__flats:
            self.__series: tuple[Note] = Scale.__flat_series
            self.__base_idx: int = Scale.__flat_series.index(tonic.title())
            return
        raise ValueError('Invalid tonic')

    @lru_cache(maxsize = 1)
    def chromatic(self: 'Scale'):
        return self.__generate_sequence(range(len(self.__series)))

    def interval(self: 'Scale', intervals: str):
        return self.__generate_sequence(accumulate(map(
            Scale.__convert_interval, intervals), initial = 0))

    def __generate_sequence(self: 'Scale', iterable: Iterable[int]):
        return [self.__series[(self.__base_idx + delta) % len(self.__series)] \
                for delta in iterable]

    @staticmethod
    def __convert_interval(interval: str) -> int:
        match interval:
            case 'A':
                return 3
            case 'M':
                return 2
            case 'm':
                return 1
            case _:
                raise ValueError(f'Invalid interval: {interval}')

    Sharp = Literal['C', 'a', 'G', 'D', 'A', 'E', 'B', 'F#', 'e', 'b', 'f#', 'c#', 'g#', 'd#']
    Flat = Literal['F', 'Bb', 'Eb', 'Ab', 'Db', 'Gb', 'd', 'g', 'c', 'f', 'bb', 'eb']
    Note = Union[Sharp, Flat]
    
    __sharps: ClassVar[frozenset[Sharp]] = frozenset(
        ('C', 'a', 'G', 'D', 'A', 'E', 'B', 'F#', 'e', 'b', 'f#', 'c#', 'g#', 'd#')
    )
    __flats: ClassVar[frozenset[Flat]] = frozenset(
        ('F', 'Bb', 'Eb', 'Ab', 'Db', 'Gb', 'd', 'g', 'c', 'f', 'bb', 'eb')
    )
    __sharp_series: ClassVar[tuple[Note]] = (
        'A', 'A#', 'B', 'C', 'C#', 'D', 'D#', 'E', 'F', 'F#', 'G', 'G#'
    )
    __flat_series: ClassVar[tuple[Note]] = (
        'A', 'Bb', 'B', 'C', 'Db', 'D', 'Eb', 'E', 'F', 'Gb', 'G', 'Ab'
    )