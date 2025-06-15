static class Appointment
{
    public static DateTime Schedule(string dateDescription) => DateTime.Parse(dateDescription);

    public static bool HasPassed(DateTime appointmentDate) => appointmentDate.CompareTo(DateTime.Now) < 0;

    public static bool IsAfternoonAppointment(DateTime appointmentDate) => appointmentDate.Hour switch
    {
        >= 12 and < 18 => true,
        _ => false,
    };

    public static string Description(DateTime appointmentDate) =>
        $"You have an appointment on {appointmentDate.ToString()}.";

    public static DateTime AnniversaryDate() => new DateTime(DateTime.Today.Year, 9, 15);
}