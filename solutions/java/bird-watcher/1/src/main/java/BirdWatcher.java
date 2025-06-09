import java.util.Arrays;
class BirdWatcher {
    private final int[] birdsPerDay;

    public BirdWatcher(int[] birdsPerDay) {
        this.birdsPerDay = birdsPerDay.clone();
    }

    public int[] getLastWeek() {
        return birdsPerDay.clone();
    }

    public int getToday() {
        return birdsPerDay[birdsPerDay.length - 1];
    }

    public void incrementTodaysCount() {
        ++birdsPerDay[birdsPerDay.length - 1];
    }

    public boolean hasDayWithoutBirds() {
        return Arrays.stream(birdsPerDay).anyMatch(v -> v == 0);
    }

    public int getCountForFirstDays(int numberOfDays) {
        return Arrays.stream(birdsPerDay).limit(numberOfDays).reduce(0, Integer::sum);
    }

    public int getBusyDays() {
        return (int)Arrays.stream(birdsPerDay).filter(v -> v >= 5).count();
    }
}
