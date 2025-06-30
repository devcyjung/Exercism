public class RemoteControlCar
{
    private int batteryPercentage = 100;
    private int distanceDrivenInMeters = 0;
    private string[] sponsors = new string[0];
    private int latestSerialNum = 0;

    public void Drive()
    {
        if (batteryPercentage > 0)
        {
            batteryPercentage -= 10;
            distanceDrivenInMeters += 2;
        }
    }

    public void SetSponsors(params string[] sponsors) => this.sponsors = sponsors;

    public string DisplaySponsor(int sponsorNum) => sponsors[sponsorNum];

    public bool GetTelemetryData(ref int serialNum, out int battery, out int distance)
    {
        if (latestSerialNum > serialNum)
        {
            serialNum = latestSerialNum;
            battery = distance = -1;
            return false;
        }
        latestSerialNum = serialNum;
        battery = batteryPercentage;
        distance = distanceDrivenInMeters;
        return true;
    }

    public static RemoteControlCar Buy() => new();
}

public class TelemetryClient
{
    private RemoteControlCar car;

    public TelemetryClient(RemoteControlCar car) => this.car = car;

    public string GetBatteryUsagePerMeter(int serialNum) =>
        !car.GetTelemetryData(ref serialNum, out int battery, out int distance) ? "no data"
            : distance == 0 ? "no data"
                : $"usage-per-meter={(100 - battery) / distance}";
}