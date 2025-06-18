from calendar import day_name, monthrange
from datetime import date as date_t

class MeetupDayException(ValueError):
    def __init__(self, message: object = 'That day does not exist.') -> None:
        super().__init__(message)

_DAY_NAMES = tuple(day_name)

def meetup(year: int, month: int, adjuster: str, day_of_week: str) -> date_t:
    result: tuple[int, int] = monthrange(year, month)
    first_weekday, days_in_month = result
    base_date: int = 1 + (_DAY_NAMES.index(day_of_week) - first_weekday + 7) % 7
    match adjuster:
        case 'first':
            return date_t(year, month, base_date)
        case 'second':
            return date_t(year, month, base_date + 7)
        case 'third':
            return date_t(year, month, base_date + 14)
        case 'fourth':
            return date_t(year, month, base_date + 21)
        case 'fifth':
            fifth_date = base_date + 28
            if fifth_date > days_in_month:
                raise MeetupDayException()
            return date_t(year, month, fifth_date)
        case 'last':
            fifth_date = base_date + 28
            return date_t(year, month,
                          fifth_date > days_in_month and base_date + 21 or fifth_date)
        case 'teenth':
            return date_t(year, month,
                          base_date >= 6 and base_date + 7 or base_date + 14)
        case _:
            raise MeetupDayException(f'Invalid adjuster format: {adjuster}')