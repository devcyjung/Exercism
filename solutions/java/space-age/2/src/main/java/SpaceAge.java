class SpaceAge {
    private final double seconds, earthYear, merYear, venYear, marYear, jupYear, satYear, uraYear, nepYear;
    private static final double EARTH_YEAR_SECS = 365.25 * 24 * 60 * 60;
    private static final double MER_ORBIT_PERIOD = 0.2408467;
    private static final double VEN_ORBIT_PERIOD = 0.61519726;
    private static final double MAR_ORBIT_PERIOD = 1.8808158;
    private static final double JUP_ORBIT_PERIOD = 11.862615;
    private static final double SAT_ORBIT_PERIOD = 29.447498;
    private static final double URA_ORBIT_PERIOD = 84.016846;
    private static final double NEP_ORBIT_PERIOD = 164.79132;
    
    SpaceAge(double seconds) {
        this.seconds = seconds;
        earthYear = seconds / EARTH_YEAR_SECS;
        merYear = earthYear / MER_ORBIT_PERIOD;
        venYear = earthYear / VEN_ORBIT_PERIOD;
        marYear = earthYear / MAR_ORBIT_PERIOD;
        jupYear = earthYear / JUP_ORBIT_PERIOD;
        satYear = earthYear / SAT_ORBIT_PERIOD;
        uraYear = earthYear / URA_ORBIT_PERIOD;
        nepYear = earthYear / NEP_ORBIT_PERIOD;
    }

    double getSeconds() {
        return seconds;
    }

    double onEarth() {
        return earthYear;
    }

    double onMercury() {
        return merYear;
    }

    double onVenus() {
        return venYear;
    }

    double onMars() {
        return marYear;
    }

    double onJupiter() {
        return jupYear;
    }

    double onSaturn() {
        return satYear;
    }

    double onUranus() {
        return uraYear;
    }

    double onNeptune() {
        return nepYear;
    }

}
