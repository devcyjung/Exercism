public class CarsAssemble {

    public double productionRatePerHour(int speed) {
        double successRatio = 1.0;
        if (speed > 4) successRatio = 0.9;
        if (speed == 9) successRatio = 0.8;
        if (speed == 10) successRatio = 0.77;
        return successRatio * speed * 221;
    }

    public int workingItemsPerMinute(int speed) {
        return (int) productionRatePerHour(speed) / 60;
    }
}
