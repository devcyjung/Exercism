class Lasagna
{
    public int ExpectedMinutesInOven() => 40;

    public int RemainingMinutesInOven(int elapsed) => 40 - elapsed;

    public int PreparationTimeInMinutes(int layers) => 2 * layers;

    public int ElapsedTimeInMinutes(int layers, int elapsed) => 2 * layers + elapsed;
}
