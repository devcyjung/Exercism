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

    String[] colors() {
        return COLORS.toArray(new String[0]);
    }
}

class ResistorColorDuo {
    static int value(String[] colors) {
        if (colors.length < 2) {
            throw new IllegalArgumentException("Less than 2 colors: " + colors.toString());
        }
        int first = ResistorColor.colorCode(colors[0]);
        int second = ResistorColor.colorCode(colors[1]);
        if (first != -1 && second != -1) {
            return first * 10 + second;
        }
        if (first == -1) {
            throw new IllegalArgumentException("Invalid first color: " + colors[0]);
        }
        throw new IllegalArgumentException("Invalid second color: " + colors[1]);
    }
}
