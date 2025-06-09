import java.util.Locale;

class Badge {
    public static String print(Integer id, String name, String department) {
        if (name == null) {
            throw NO_NAME_EXCEPTION;
        }
        var dpt = switch (department) {
            case null -> OWNER_DPT;
            default -> department.toUpperCase();
        };
        return switch (id) {
            case null -> String.format(NEW_EMPLOYEE_FMT, name, dpt);
            default -> String.format(EMPLOYEE_FMT, id, name, dpt);
        };
    }

    private static final String EMPLOYEE_FMT = "[%d] - %s - %s";
    private static final String NEW_EMPLOYEE_FMT = "%s - %s";
    private static final String OWNER_DPT = "OWNER";
    private static final IllegalArgumentException NO_NAME_EXCEPTION =
        new IllegalArgumentException("Must provide non-null name");
}