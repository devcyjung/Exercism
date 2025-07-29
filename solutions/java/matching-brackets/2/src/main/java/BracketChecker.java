import java.util.ArrayDeque;

class BracketChecker {

    private final boolean isWellformed;
    private final String expression;
    
    BracketChecker(String expression) {
        this.expression = expression;
        var stack = new ArrayDeque<Integer>();
        for (char ch: expression.toCharArray()) {
            int encoding = encode(ch);
            if (encoding > 0) {
                stack.addLast(encoding);
            } else if (encoding < 0) {
                Integer last = stack.pollLast();
                if (last == null || encoding + last != 0) {
                    isWellformed = false;
                    return;
                }
            }
        }
        isWellformed = stack.isEmpty();
    }

    private static int encode(char ch) {
        return switch (ch) {
            case '[' -> 1;
            case ']' -> -1;
            case '{' -> 2;
            case '}' -> -2;
            case '(' -> 3;
            case ')' -> -3;
            default -> 0;
        };
    }

    boolean areBracketsMatchedAndNestedCorrectly() {
        return isWellformed;
    }

}