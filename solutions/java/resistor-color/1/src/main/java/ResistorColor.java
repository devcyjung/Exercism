import java.util.Arrays;

class ResistorColor {
    static private String[] COLORS = {"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"};
    
    static int colorCode(String color) {
        return Arrays.asList(COLORS).indexOf(color);
    }

    String[] colors() {
        return Arrays.copyOf(COLORS, COLORS.length);
    }
}
