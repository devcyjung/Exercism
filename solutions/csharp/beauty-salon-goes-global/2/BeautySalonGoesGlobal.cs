using System.Globalization;
using System.Runtime.InteropServices;

public enum Location
{
    NewYork,
    London,
    Paris
}

public enum AlertLevel
{
    Early,
    Standard,
    Late
}

public static class Appointment
{
    public static DateTime ShowLocalTime(DateTime dtUtc) => dtUtc.ToLocalTime();

    private static bool isWindows = RuntimeInformation.IsOSPlatform(OSPlatform.Windows);
        
    private static string GetTimeZoneId(Location location) => location switch
    {
        Location.NewYork => isWindows ? "Eastern Standard Time" : "America/New_York",
        Location.London  => isWindows ? "GMT Standard Time" : "Europe/London",
        Location.Paris   => isWindows ? "W. Europe Standard Time": "Europe/Paris",
        _ => throw new ArgumentOutOfRangeException(nameof(location), "Unknown location")
    };

    private static TimeZoneInfo GetTimeZoneInfo(Location location) =>
        TimeZoneInfo.FindSystemTimeZoneById(GetTimeZoneId(location));

    public static DateTime Schedule(string appointmentDateDescription, Location location) =>
        TimeZoneInfo.ConvertTimeToUtc(DateTime.Parse(appointmentDateDescription), GetTimeZoneInfo(location));

    public static DateTime GetAlertTime(DateTime appointment, AlertLevel alertLevel) => alertLevel switch
    {
        AlertLevel.Early => appointment.AddDays(-1),
        AlertLevel.Standard => appointment.AddHours(-1.75),
        AlertLevel.Late => appointment.AddHours(-0.5),
        _ => throw new ArgumentOutOfRangeException(nameof(alertLevel), "Unknown alertLevel")
    };

    public static bool HasDaylightSavingChanged(DateTime dt, Location location) =>
        GetTimeZoneInfo(location).IsDaylightSavingTime(DateTime.SpecifyKind(dt, DateTimeKind.Local)) !=
        GetTimeZoneInfo(location).IsDaylightSavingTime(DateTime.SpecifyKind(dt.AddDays(-7), DateTimeKind.Local));

    private static CultureInfo GetCultureInfo(Location location) => location switch
    {
        Location.NewYork => CultureInfo.GetCultureInfo("en-US"),
        Location.London  => CultureInfo.GetCultureInfo("en-GB"),
        Location.Paris   => CultureInfo.GetCultureInfo("fr-FR"),
        _ => throw new ArgumentOutOfRangeException(nameof(location), "Unknown location")
    };

    public static DateTime NormalizeDateTime(string dtStr, Location location) =>
        DateTime.TryParse(dtStr, GetCultureInfo(location), out DateTime result) ? result : DateTime.MinValue;
}