public class RemoteControlCar
{
    public RemoteControlCar() => telemetry = new(this);
    public string CurrentSponsor { get; private set; }
    public ITelemetry Telemetry => telemetry;
    public string GetSpeed() => currentSpeed.ToString();

    public interface ITelemetry
    {
        void Calibrate();
        bool SelfTest();
        void ShowSponsor(string sponsor);
        void SetSpeed(decimal amount, string units);
    }

    private readonly TelemetryImpl telemetry;
    private Speed currentSpeed;
    
    private class TelemetryImpl: ITelemetry
    {
        private RemoteControlCar car;
        internal TelemetryImpl(RemoteControlCar car) => this.car = car; 
        public void Calibrate() => _ = 0;
        public bool SelfTest() => true;
        public void ShowSponsor(string sponsor) => car.CurrentSponsor = sponsor;
        public void SetSpeed(decimal amount, string unitString) => car.currentSpeed = unitString switch
        {
            "cps" => new(amount, SpeedUnits.CentimetersPerSecond),
            "mps" => new(amount, SpeedUnits.MetersPerSecond),
            _ => throw new ArgumentOutOfRangeException(nameof(unitString), unitString, "Unexpected unit string"),
        };
    }

    private readonly record struct Speed(decimal Amount, SpeedUnits SpeedUnits)
    {
        public override string ToString() => SpeedUnits switch
        {
            SpeedUnits.CentimetersPerSecond => $"{Amount} centimeters per second",
            SpeedUnits.MetersPerSecond => $"{Amount} meters per second",
            _ => throw new ArgumentOutOfRangeException(nameof(SpeedUnits), SpeedUnits, "Unexpected speed unit"),
        };
    }

    private enum SpeedUnits { MetersPerSecond, CentimetersPerSecond }
}