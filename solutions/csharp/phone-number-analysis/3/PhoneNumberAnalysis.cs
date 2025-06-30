public static class PhoneNumber
{
    public static (bool isNY, bool isFake, string lastFour) Analyze(string number) =>
        (number[0..3] == "212", number[4..7] == "555", number[(number.Length - 4)..]);

    public static bool IsFake((bool, bool isFake, string) info) => info.isFake;
}