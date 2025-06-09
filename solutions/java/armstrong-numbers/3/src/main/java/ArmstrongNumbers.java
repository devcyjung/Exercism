class ArmstrongNumbers {
    boolean isArmstrongNumber(int numberToCheck) {
        final var stringified = Integer.toString(numberToCheck);
        final var size = stringified.length();
        return numberToCheck == stringified.chars()
            .map(Character::getNumericValue)
            .map(i -> (int) Math.pow(i, size))
            .reduce(0, Integer::sum);
    }
}
