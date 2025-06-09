class ArmstrongNumbers {
    boolean isArmstrongNumber(int numberToCheck) {
        var stringified = Integer.toString(numberToCheck);
        var size = stringified.length();
        return numberToCheck == stringified.chars()
            .map(Character::getNumericValue)
            .map(i -> (int) Math.pow(i, size))
            .reduce(0, Integer::sum);
    }
}
