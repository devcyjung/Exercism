import java.util.Objects;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public final class GameMaster {
    public static String describe(Character character) {
        return String.join(
            " ", "You're a level", Integer.toString(character.getLevel()),
            character.getCharacterClass(), "with",
            Integer.toString(character.getHitPoints()), "hit points."
        );
    }

    public static String describe(Destination destination) {
        return String.join(
            " ", "You've arrived at", destination.getName().concat(","),
            "which has", Integer.toString(destination.getInhabitants()),
            "inhabitants."
        );
    }

    public static String describe(TravelMethod travelMethod) {
        return String.join(
            " ", "You're traveling to your destination",
            (switch (travelMethod) {
                case WALKING -> "by walking";
                case HORSEBACK -> "on horseback";
            }).concat(".")
        );
    }

    public static String describe(
        Character character, Destination destination, TravelMethod travelMethod
    ) {
        return Stream.of(character, travelMethod, destination)
            .filter(Objects::nonNull)
            .map(param -> switch(param) {
                case Character c -> describe(c);
                case Destination d -> describe(d);
                case TravelMethod t -> describe(t);
                default -> throw new IllegalArgumentException("Unknown type");
            })
            .collect(Collectors.joining(" "));
    }

    public static String describe(Character character, Destination destination) {
        return describe(character, destination, TravelMethod.WALKING);
    }
}
