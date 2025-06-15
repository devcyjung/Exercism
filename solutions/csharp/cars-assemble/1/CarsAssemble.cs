static class AssemblyLine
{
    public static double SuccessRate(int speed) => speed switch
    {
        0 => 0.0,
        >= 1 and <= 4 => 1.0,
        >= 5 and <= 8 => 0.9,
        9 => 0.8,
        10 => 0.77,
        _ => throw new ArgumentOutOfRangeException(nameof(speed), "Speed must be between 0 and 10"),
    };
    
    public static double ProductionRatePerHour(int speed) => 221.0 * speed * SuccessRate(speed);

    public static int WorkingItemsPerMinute(int speed) => (int) ProductionRatePerHour(speed) / 60;
}
