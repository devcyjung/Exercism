class BirdCount
{
    private int[] birdsPerDay;

    public BirdCount(int[] birdsPerDay) => this.birdsPerDay = birdsPerDay;

    public static int[] LastWeek() => new int[] {0, 2, 5, 3, 7, 8, 4};

    public int Today() => birdsPerDay[birdsPerDay.Length - 1];

    public void IncrementTodaysCount() => ++birdsPerDay[birdsPerDay.Length - 1];

    public bool HasDayWithoutBirds() => !Array.TrueForAll(birdsPerDay, bird => bird != 0);

    public int CountForFirstDays(int numberOfDays)
    {
        var sum = 0;
        for (var i = 0; i < numberOfDays && i < birdsPerDay.Length; ++i)
        {
            sum += birdsPerDay[i];
        }
        return sum;
    }

    public int BusyDays()
    {
        var count = 0;
        foreach (var birds in birdsPerDay)
        {
            if (birds >= 5)
            {
                ++count;
            }
        }
        return count;
    }
}
