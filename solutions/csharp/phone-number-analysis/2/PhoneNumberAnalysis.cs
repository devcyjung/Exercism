public static class PhoneNumber
{
    public static (bool IsNewYork, bool IsFake, string LocalNumber) Analyze(string phoneNumber) =>
    (
        phoneNumber[0..3] == "212",
        phoneNumber[4..7] == "555",
        phoneNumber[(phoneNumber.Length - 4)..]
    );

    public static bool IsFake((bool IsNewYork, bool IsFake, string LocalNumber) phoneNumberInfo) =>
        phoneNumberInfo.IsFake;
}