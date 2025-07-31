class NaturalNumber {
    final int number;
    final int aliquot;
    
    NaturalNumber(int number) {
        if (number <= 0) {
            throw new IllegalArgumentException("You must supply a natural number (positive integer)");
        }
        int aliquot = 0;
        for (int i = 1; i <= number / 2; ++i) {
            if (number % i == 0) {
                aliquot += i;
            }
        }
        this.number = number;
        this.aliquot = aliquot;
    }

    Classification getClassification() {
        if (aliquot == number) {
            return Classification.PERFECT;
        } else if (aliquot > number) {
            return Classification.ABUNDANT;
        } else {
            return Classification.DEFICIENT;
        }
    }
}
