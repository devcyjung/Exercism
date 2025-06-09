public class JedliksToyCar {
    public static JedliksToyCar buy() {
        return new JedliksToyCar();
    }

    private int distance = 0;
    private int battery = 100;
    private static final String distanceFormat = "Driven %d meters";
    private static final String batteryFormat = "Battery at %d%%";
    private static final String emptyBatteryMessage = "Battery empty";
    private static final int speed = 20;

    public String distanceDisplay() {
        return String.format(distanceFormat, distance);
    }

    public String batteryDisplay() {
        if (isDrained()) {
            return emptyBatteryMessage;
        }
        return String.format(batteryFormat, battery);
    }

    public boolean isDrained() {
        return battery <= 0;
    }
    
    public void drive() {
        if (isDrained()) {
            return;
        }
        --battery;
        distance += speed;
    }
}
