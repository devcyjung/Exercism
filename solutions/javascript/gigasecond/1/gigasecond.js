export const gigasecond = (date) => {
  const newDate = new Date(date)
  newDate.setTime(date.getTime() + 1_000_000_000_000)
  return newDate
};
