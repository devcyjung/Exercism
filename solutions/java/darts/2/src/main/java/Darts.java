class Darts {
    int score(double xOfDart, double yOfDart) {
        final var radius = Math.hypot(xOfDart, yOfDart);
        return (radius > 10) ? 0 :
               (radius > 5) ? 1 :
               (radius > 1) ? 5 : 10;
    }
}
