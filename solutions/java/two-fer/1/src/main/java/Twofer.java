public class Twofer {
    public static String twofer(String name) {
        return switch (name) {
            case null -> DEFAULT_STRING;
            case "" -> DEFAULT_STRING;
            default -> "One for " + name + ", one for me.";
        };
    }

    private static String DEFAULT_STRING = "One for you, one for me.";
}