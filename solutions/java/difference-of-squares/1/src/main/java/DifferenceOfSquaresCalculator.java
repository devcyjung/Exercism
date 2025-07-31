class DifferenceOfSquaresCalculator {

    static int computeSquareOfSumTo(int input) {
        int sum = 0;
        for (int i = 1; i <= input; ++i) {
            sum += i;
        }
        return sum * sum;
    }

    static int computeSumOfSquaresTo(int input) {
        int sum = 0;
        for (int i = 1; i <= input; ++i) {
            sum += i * i;
        }
        return sum;
    }

    static int computeDifferenceOfSquares(int input) {
        int sum = 0;
        for (int i = 1; i <= input; ++i) {
            sum += i;
        }
        int diff = 0;
        for (int i = 1; i <= input; ++i) {
            diff += i * (sum - i);
        }
        return diff;
    }

}
