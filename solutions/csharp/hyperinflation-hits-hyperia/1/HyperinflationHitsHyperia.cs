using System.Numerics;

public static class CentralBank
{
    public static string MulWithDefault<T>(T a, T b, string defaultString = "*** Too Big ***")
        where T : INumber<T>
    {
        try
        {
            T result = checked(a * b);
            return T.IsFinite(result) ? result.ToString() : defaultString;
        }
        catch (OverflowException)
        {
            return defaultString;
        }
    }
    
    public static string DisplayDenomination(long @base, long multiplier) =>
        MulWithDefault(@base, multiplier);

    public static string DisplayGDP(float @base, float multiplier) =>
        MulWithDefault(@base, multiplier);

    public static string DisplayChiefEconomistSalary(decimal salaryBase, decimal multiplier) =>
        MulWithDefault(salaryBase, multiplier, "*** Much Too Big ***");
}
