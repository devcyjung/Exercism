public static class BafflingBirthdays
{
    private static readonly Random Rng = Random.Shared;

    public static DateOnly[] RandomBirthdates(int numberOfBirthdays) =>   
        Enumerable.Range(0, numberOfBirthdays).Select(_ =>
        {
            int month = Rng.Next(1, 13);
            int day = Rng.Next(1, DateTime.DaysInMonth(2001, month) + 1);
            return new DateOnly(2001, month, day);
        }).ToArray();
    
    public static bool SharedBirthday(DateOnly[] birthdays) =>
        birthdays.GroupBy(d => (d.Month, d.Day)).Any(g => g.Count() > 1);

    public static double EstimatedProbabilityOfSharedBirthday(int numberOfBirthdays) =>
        100.0 * (1.0 - Enumerable.Range(365 - numberOfBirthdays + 1, numberOfBirthdays)
                 .Aggregate(1.0, (acc, cur) => acc * cur / 365.0));
}
