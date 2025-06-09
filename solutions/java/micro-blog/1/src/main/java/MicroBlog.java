import java.text.BreakIterator;

class MicroBlog {
    private static BreakIterator iter = BreakIterator.getCharacterInstance();
    
    public String truncate(String input) {
        iter.setText(input);
        var boundary = iter.next(5);
        if (boundary == BreakIterator.DONE) {
            return input;
        }
        return input.substring(0, boundary);
    }
}