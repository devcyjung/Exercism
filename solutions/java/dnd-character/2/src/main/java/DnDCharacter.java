import java.security.SecureRandom;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

record DnDCharacter(
    int getStrength, int getDexterity, int getConstitution,
    int getIntelligence, int getWisdom, int getCharisma
) {
    private static SecureRandom RNG = new SecureRandom();

    public DnDCharacter() {
        this(
            ability(), ability(), ability(),
            ability(), ability(), ability()
        );
    }

    public static int ability(List<Integer> scores) {
        return scores.stream().sorted().skip(1)
            .reduce(Integer::sum).orElseThrow();
    }

    public static int ability() {
        return ability(rollDice());
    }

    public static int roll() {
        return RNG.nextInt(1, 7);
    }

    public static List<Integer> rollDice() {
        return IntStream.generate(DnDCharacter::roll)
            .limit(4).boxed().collect(Collectors.toList());
    }

    public static int modifier(int input) {
        return Math.floorDiv(input - 10, 2);
    }

    public int getHitpoints() {
        return 10 + modifier(getConstitution());
    }
}