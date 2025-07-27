public static class Grains
{
    public static ulong Square(int n)
    {
        if (n < 1 || n > 64)
        {
            throw new ArgumentOutOfRangeException("n must be between 1 and 64");
        }
        return 1UL << (n - 1);
    }

    public static ulong Total() => ~0UL;
}