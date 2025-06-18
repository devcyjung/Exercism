from datetime import date, timedelta

class MeetupDayException(ValueError):
    """Exception raised when the Meetup weekday and count do not result in a valid date.

    message: explanation of the error.

    """
    def __init__(self, message: object = None) -> None:
        super().__init__(message)


def meetup(year: int, month: int, week: str, day_of_week: str) -> date:
    if month < 1 or month > 12:
        raise MeetupDayException(f'Invalid month: {month}')
    start_date: date = _start_date(year, month, week)
    delta: int = _ord_weekday(day_of_week) - start_date.weekday()
    meetup_date: date = start_date + timedelta(days = (delta + 7) % 7)
    if month != meetup_date.month:
        raise MeetupDayException('That day does not exist.')
    return meetup_date

def _ord_weekday(weekday: str) -> int:
    match weekday.title():
        case 'Monday':
            return 0
        case 'Tuesday':
            return 1
        case 'Wednesday':
            return 2
        case 'Thursday':
            return 3
        case 'Friday':
            return 4
        case 'Saturday':
            return 5
        case 'Sunday':
            return 6
        case _:
            raise MeetupDayException(f'Invalid weekday: {weekday}')

def _start_date(year: int, month: int, week: str) -> date:
    match week.lower():
        case 'teenth':
            return date(year, month, 13)
        case 'first':
            return date(year, month, 1)
        case 'second':
            return date(year, month, 8)
        case 'third':
            return date(year, month, 15)
        case 'fourth':
            return date(year, month, 22)
        case 'fifth':
            return date(year, month, 22) + timedelta(days = 7)
        case 'last':
            if month == 12:
                year += 1
                month = 1
            else:
                month += 1
            return date(year, month, 1) - timedelta(days = 7)
        case _:
            raise MeetupDayException(f'Invalid week format: {week}')
