public enum LogLevel {
    TRACE(1), DEBUG(2), INFO(4), WARNING(5), ERROR(6), FATAL(42), UNKNOWN(0);

    private final int shortNumber;

    private LogLevel(int shortNumber) {
        this.shortNumber = shortNumber;
    }

    @Override
    public String toString() {
        return Integer.toString(shortNumber);
    }

    public int valueOf() {
        return shortNumber;
    }
}