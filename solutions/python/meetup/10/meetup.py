from calendar import day_name, monthrange
from datetime import date as date_t

class MeetupDayException(ValueError):
    def __init__(self, message: object = 'That day does not exist.') -> None:
        super().__init__(message)

_WEEKDAY_NAMES: tuple[str, ...] = tuple(day_name)
_WEEKS: tuple[str, ...] = ('first', 'second', 'third', 'fourth', 'fifth')

def meetup(year: int, month: int, week_str: str, day_of_week: str) -> date_t:
    result: tuple[int, int] = monthrange(year, month)
    first_weekday, last_day_in_month = result
    week_1: int = 1 + (_WEEKDAY_NAMES.index(day_of_week) - first_weekday + 7) % 7
    match week_str:
        case 'last':
            week_5 = week_1 + 28
            date = week_5 - 7 if week_5 > last_day_in_month else week_5
            return date_t(year, month, date)
        case 'teenth':
            week_2 = week_1 + 7
            date = week_2 if week_2 >= 13 else week_2 + 7
            return date_t(year, month, date)
        case _:
            date = week_1 + 7 * _WEEKS.index(week_str)
            if date > last_day_in_month:
                raise MeetupDayException()
            return date_t(year, month, date)