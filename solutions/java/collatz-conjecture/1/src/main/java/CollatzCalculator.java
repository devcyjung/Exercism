class CollatzCalculator {

    static int computeStepCount(int start) {
        if (start <= 0) {
            throw new IllegalArgumentException("Only positive integers are allowed");
        }
        int step = 0;
        while (start != 1) {
            if ((start & 1) == 1) {
                start += (start << 1) + 1;
                ++step;
            } else {
                int trailingZeros = Integer.numberOfTrailingZeros(start);
                start >>= trailingZeros;
                step += trailingZeros;
            }
        }
        return step;
    }

}
