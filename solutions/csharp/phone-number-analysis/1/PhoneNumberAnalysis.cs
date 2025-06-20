public static class PhoneNumber
{
    public static (bool IsNewYork, bool IsFake, string LocalNumber) Analyze(string phoneNumber) =>
    (
        phoneNumber.Substring(0, 3).Equals("212"),
        phoneNumber.Substring(4, 3).Equals("555"),
        phoneNumber.Substring(phoneNumber.Length - 4)
    );

    public static bool IsFake((bool IsNewYork, bool IsFake, string LocalNumber) phoneNumberInfo) =>
        phoneNumberInfo.IsFake;
}