final class NeedForSpeed {
    private final int speed;
    private final int batteryDrain;
    private int distanceDriven = 0;
    private int batteryRemaining = 100;
    
    NeedForSpeed(int speed, int batteryDrain) {
        this.speed = speed;
        this.batteryDrain = batteryDrain;
    }

    public boolean batteryDrained() {
        return batteryRemaining < batteryDrain;
    }

    public int distanceDriven() {
        return distanceDriven;
    }

    public void drive() {
        if (batteryDrained()) {
            return;
        }
        distanceDriven += speed;
        batteryRemaining -= batteryDrain;
    }

    public int maximumDistance() {
        return batteryRemaining / batteryDrain * speed;
    }

    public static NeedForSpeed nitro() {
        return new NeedForSpeed(50, 4);
    }
}

final class RaceTrack {
    private final int distance;
    
    RaceTrack(int distance) {
        this.distance = distance;
    }

    public boolean canFinishRace(NeedForSpeed car) {
        return car.maximumDistance() >= distance;
    }
}
