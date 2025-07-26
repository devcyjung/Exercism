public static class Leap
{
    public static bool IsLeapYear(int year) => (year & 3) == 0 && year % 100 != 0 || year % 400 == 0;
}