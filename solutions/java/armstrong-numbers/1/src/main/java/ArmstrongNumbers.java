class ArmstrongNumbers {
    boolean isArmstrongNumber(int numberToCheck) {
        var stringified = Integer.toString(numberToCheck);
        var size = stringified.length();
        return numberToCheck == stringified.chars()
            .map(ch -> (int) Math.pow(ch - '0', size))
            .reduce(0, Integer::sum);
    }
}
