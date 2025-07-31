import java.util.Arrays;
import java.util.List;

class ResistorColor {
    static final private List<String> COLORS = List.of(
        new String[]{
            "black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"
        }
    );
    
    static int colorCode(String color) {
        return COLORS.indexOf(color);
    }

    static String[] colors() {
        return COLORS.toArray(new String[0]);
    }
}
