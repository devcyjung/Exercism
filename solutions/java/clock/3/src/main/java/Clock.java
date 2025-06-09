class Clock {
    private HourMinuteRecord hmRecord;

    Clock(int hours, int minutes) {
        hmRecord = new HourMinuteRecord(hours, minutes);
    }

    void add(int minutes) {
        hmRecord = hmRecord.plusMinutes(minutes);
    }

    @Override
    public String toString() {
        return hmRecord.toString();
    }

    @Override
    public boolean equals(Object obj) {
        return obj instanceof Clock clock
            && clock.hmRecord.equals(hmRecord);
    }

    @Override
    public int hashCode() {
        return hmRecord.hashCode();
    }
}

record HourMinuteRecord(int hours, int minutes) {
    private static final int MIN_PER_DAY = 24 * 60;
    
    public HourMinuteRecord {
        final int totalMinutes = Math.floorMod(hours * 60 + minutes, MIN_PER_DAY);
        hours = Math.floorDiv(totalMinutes, 60);
        minutes = Math.floorMod(totalMinutes, 60);
    }

    public HourMinuteRecord plusMinutes(int addedMinutes) {
        return new HourMinuteRecord(0, hours * 60 + minutes + addedMinutes);
    }

    @Override
    public String toString() {
        return String.format("%02d:%02d", hours, minutes);
    }
}