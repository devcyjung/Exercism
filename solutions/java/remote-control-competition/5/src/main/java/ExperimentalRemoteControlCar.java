public class ExperimentalRemoteControlCar implements RemoteControlCar {
    private int distance = 0;
    private static int SPEED = 20;
    
    @Override
    public void drive() {
        distance += SPEED;
    }

    @Override
    public int getDistanceTravelled() {
        return distance;
    }
}