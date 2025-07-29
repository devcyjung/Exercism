class CollatzCalculator {

    static int computeStepCount(int start) {
        if (start <= 0) {
            throw new IllegalArgumentException("Only positive integers are allowed");
        }
        int step = 0;
        while (start != 1) {
            int trailingZeros = Integer.numberOfTrailingZeros(start);
            if (trailingZeros == 0) {
                start += (start << 1) + 1;
                ++step;
            } else {
                start >>= trailingZeros;
                step += trailingZeros;
            }
        }
        return step;
    }

}
