public static class SwiftScheduling
{
    public static DateTime DeliveryDate(DateTime meetingStart, string description)
    {
        switch (description)
        {
            case "NOW":
                return meetingStart.AddHours(2);
            case "ASAP":
                return meetingStart.Hour < 13 ? meetingStart.Date.AddHours(17) : meetingStart.Date.AddDays(1).AddHours(13);
            case "EOW":
            {
                DayOfWeek weekday = meetingStart.DayOfWeek;
                DateTime weekBeginning = meetingStart.Date.AddDays(- (double) weekday);
                switch (weekday)
                {
                    case DayOfWeek.Monday or DayOfWeek.Tuesday or DayOfWeek.Wednesday:
                        return weekBeginning.AddDays(5).AddHours(17);
                    case DayOfWeek.Thursday or DayOfWeek.Friday:
                        return weekBeginning.AddDays(7).AddHours(20);
                }
                break;
            }
        }

        if (description.EndsWith('M'))
        {
            int nth = int.Parse(description[..^1]);
            int month = meetingStart.Month;
            int year = meetingStart.Year;
            DateTime deliveryDay = month < nth ? new DateTime(year, nth, 1) : new DateTime(year + 1, nth, 1);
            while (deliveryDay.DayOfWeek is DayOfWeek.Saturday or DayOfWeek.Sunday)
            {
                deliveryDay = deliveryDay.AddDays(1);
            }
            return deliveryDay.AddHours(8);
        }
        if (description.StartsWith('Q'))
        {
            int nth = int.Parse(description[1..]);
            int nthAsMonth = (nth - 1) * 3 + 1;
            int month = meetingStart.Month;
            int year = meetingStart.Year;
            int quarter = (month - 1) / 3 + 1;
            DateTime deliveryDay = quarter > nth ? new DateTime(year + 1, nthAsMonth, 1) : new DateTime(year, nthAsMonth, 1);
            deliveryDay = deliveryDay.AddMonths(2);
            deliveryDay = deliveryDay.AddDays(DateTime.DaysInMonth(deliveryDay.Year, deliveryDay.Month) - 1);
            while (deliveryDay.DayOfWeek is DayOfWeek.Saturday or DayOfWeek.Sunday)
            {
                deliveryDay = deliveryDay.AddDays(-1);
            }
            return deliveryDay.AddHours(8);
        }
        throw new InvalidDataException($"Invalid input: meetingStart - {meetingStart}, description - {description}");
    }
}
