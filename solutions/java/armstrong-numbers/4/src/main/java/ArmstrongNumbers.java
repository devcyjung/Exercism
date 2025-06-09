import java.util.function.IntUnaryOperator;
import java.util.stream.IntStream;

class ArmstrongNumbers {
    static boolean isArmstrongNumber(int numberToCheck) {
        final var stringified = Integer.toString(numberToCheck);
        final var pow = ipow(stringified.length());
        return numberToCheck == stringified.chars()
            .map(Character::getNumericValue)
            .map(pow)
            .reduce(0, Integer::sum);
    }
    
    static IntUnaryOperator ipow(int exp) {
        if (exp < 0) {
            throw new IllegalArgumentException(
                "ipow only accepts positive int, got: " + exp
            );
        }
        return base -> IntStream.generate(() -> base)
            .limit(exp)
            .reduce(1, Math::multiplyExact);
    }
}
