from calendar import day_name, monthrange
from datetime import date as date_t

class MeetupDayException(ValueError):
    def __init__(self, message: object = 'That day does not exist.') -> None:
        super().__init__(message)

_DAY_NAMES = tuple(day_name)
_OFFSETS = ('first', 'second', 'third', 'fourth', 'fifth')

def meetup(year: int, month: int, adjuster: str, day_of_week: str) -> date_t:
    result: tuple[int, int] = monthrange(year, month)
    first_weekday, days_in_month = result
    base_date: int = 1 + (_DAY_NAMES.index(day_of_week) - first_weekday + 7) % 7
    match adjuster:
        case 'last':
            fifth_date = base_date + 28
            date = fifth_date - 7 if fifth_date > days_in_month else fifth_date
            return date_t(year, month, date)
        case 'teenth':
            second_date = base_date + 7
            date = second_date if second_date >= 13 else second_date + 7
            return date_t(year, month, date)
        case _:
            date = base_date + 7 * _OFFSETS.index(adjuster)
            if date > days_in_month:
                raise MeetupDayException()
            return date_t(year, month, date)