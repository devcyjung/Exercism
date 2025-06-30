using System.Globalization;
using System.Runtime.InteropServices;

public enum Location
{
    NewYork,
    London,
    Paris
}

public static class LocationExtension
{
    private static bool isWindows = RuntimeInformation.IsOSPlatform(OSPlatform.Windows);
    
    public static string GetTimeZoneId(this Location location) => location switch
    {
        Location.NewYork => isWindows ? "Eastern Standard Time" : "America/New_York",
        Location.London  => isWindows ? "GMT Standard Time" : "Europe/London",
        Location.Paris   => isWindows ? "W. Europe Standard Time": "Europe/Paris",
    };

    public static TimeZoneInfo GetTimeZoneInfo(this Location location) =>
        TimeZoneInfo.FindSystemTimeZoneById(location.GetTimeZoneId());

    public static CultureInfo GetCultureInfo(this Location location) => location switch
    {
        Location.NewYork => CultureInfo.GetCultureInfo("en-US"),
        Location.London  => CultureInfo.GetCultureInfo("en-GB"),
        Location.Paris   => CultureInfo.GetCultureInfo("fr-FR"),
    };
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
        
    public static DateTime Schedule(string appointmentDateDescription, Location location) =>
        TimeZoneInfo.ConvertTimeToUtc(DateTime.Parse(appointmentDateDescription), location.GetTimeZoneInfo());

    public static DateTime GetAlertTime(DateTime appointment, AlertLevel alertLevel) => alertLevel switch
    {
        AlertLevel.Early => appointment.AddDays(-1),
        AlertLevel.Standard => appointment.AddHours(-1.75),
        AlertLevel.Late => appointment.AddHours(-0.5),
    };

    public static bool HasDaylightSavingChanged(DateTime dt, Location location) =>
        location.GetTimeZoneInfo().IsDaylightSavingTime(DateTime.SpecifyKind(dt, DateTimeKind.Local)) !=
        location.GetTimeZoneInfo().IsDaylightSavingTime(DateTime.SpecifyKind(dt.AddDays(-7), DateTimeKind.Local));

    public static DateTime NormalizeDateTime(string dtStr, Location location) =>
        DateTime.TryParse(dtStr, location.GetCultureInfo(), out DateTime result) ? result : DateTime.MinValue;
}