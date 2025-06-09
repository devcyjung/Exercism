import java.util.Locale;

class Badge {
    public static String print(Integer id, String name, String department) {
        if (name == null) {
            throw new IllegalArgumentException("Must provide non-null name");
        }
        var dpt = switch (department) {
            case null -> OWNER_DPT;
            default -> department.toUpperCase();
        };
        return switch (id) {
            case null -> name + " - " + dpt;
            default -> "[" + id + "]" + " - " + name + " - " + dpt;
        };
    }

    private static final String OWNER_DPT = "OWNER";
}