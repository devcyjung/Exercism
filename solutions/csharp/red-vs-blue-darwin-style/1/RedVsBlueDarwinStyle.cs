namespace RedRemoteControlCarTeam
{
    public class RemoteControlCar(Motor m, Chassis c, Telemetry t, RunningGear r);
    public class RunningGear;
    public class Telemetry;
    public class Chassis;
    public class Motor;
}

namespace BlueRemoteControlCarTeam
{
    public class RemoteControlCar(Motor m, Chassis c, Telemetry t);
    public class Telemetry;
    public class Chassis;
    public class Motor;
}

namespace Combined
{
    using Red = RedRemoteControlCarTeam;
    using Blue = BlueRemoteControlCarTeam;
    public static class CarBuilder
    {   
        public static Red.RemoteControlCar BuildRed() => new(new(), new(), new(), new());
        public static Blue.RemoteControlCar BuildBlue() => new(new(), new(), new());
    }   
}