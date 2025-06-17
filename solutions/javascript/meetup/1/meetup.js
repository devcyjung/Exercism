const START_DATES = {teenth: 13, first: 1, second: 8, third: 15, fourth: 22, last: -6}
const WEEKDAYS = {Sunday: 0, Monday: 1, Tuesday: 2, Wednesday: 3, Thursday: 4, Friday: 5, Saturday: 6}

export const meetup = (year, month, week, weekday) => {
  const meetupDate = new Date(year, week !== 'last' ? month - 1 : month, START_DATES[week])
  meetupDate.setDate(meetupDate.getDate() + (WEEKDAYS[weekday] - meetupDate.getDay() + 7) % 7)
  return meetupDate
}