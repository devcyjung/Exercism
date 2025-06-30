public class WeatherStation
{
    private Reading reading;
    private List<DateTime> recordDates = new();
    private List<decimal> temperatures = new();

    public void AcceptReading(Reading reading)
    {
        this.reading = reading;
        recordDates.Add(DateTime.Now);
        temperatures.Add(reading.Temperature);
    }

    public void ClearAll()
    {
        reading = new Reading();
        recordDates.Clear();
        temperatures.Clear();
    }

    public decimal LatestTemperature => reading.Temperature;

    public decimal LatestPressure => reading.Pressure;

    public decimal LatestRainfall => reading.Rainfall;

    public bool HasHistory => recordDates.Count > 1;

    public Outlook ShortTermOutlook => reading switch
    {
        Reading r when r == new Reading() => throw new ArgumentException(),
        { Pressure: < 10m } and { Temperature: < 30m } => Outlook.Cool,
        { Temperature: > 50m } => Outlook.Good,
        _ => Outlook.Warm,
    };
    
    public Outlook LongTermOutlook => reading.WindDirection switch
    {
        WindDirection.Southerly => Outlook.Good,
        WindDirection.Northerly => Outlook.Cool,
        WindDirection.Easterly => reading.Temperature > 20 ? Outlook.Good : Outlook.Warm,
        WindDirection.Westerly => Outlook.Rainy,
        _ => throw new ArgumentException(),
    };

    public State RunSelfTest() => reading == new Reading() ? State.Bad : State.Good;
}

public readonly record struct Reading(decimal Temperature, decimal Pressure, decimal Rainfall, WindDirection WindDirection);

public enum State { Good, Bad }

public enum Outlook { Cool, Rainy, Warm, Good }

public enum WindDirection { Unknown, Northerly, Easterly, Southerly, Westerly }