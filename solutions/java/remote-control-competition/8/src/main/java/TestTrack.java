import java.util.Collection;
import java.util.List;

public class TestTrack {
    public static void race(RemoteControlCar car) {
        car.drive();
    }

    public static <T extends Comparable<T> & RemoteControlCar>
    List<T> getRankedCars(Collection<T> cars) {
        return cars.stream().sorted().toList();
    }
}