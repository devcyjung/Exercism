class SqueakyClean {
    
    static boolean shouldCapitalize(StringBuilder sb) {
        return sb.length() == 0 || sb.charAt(sb.length()-1) == ' ';
    }
    
    static void appendCapitalized(StringBuilder sb, char ch) {
        if (shouldCapitalize(sb)) {
            sb.append(Character.toUpperCase(ch));
        } else {
            sb.append(ch);
        }
    }

    static void appendNormally(StringBuilder sb, char ch) {
        sb.append(ch);
    }
    
    static String clean(String identifier) {
        StringBuilder sb = new StringBuilder();
        for (int sourceIndex=0; sourceIndex<identifier.length(); ++sourceIndex) {
            switch (identifier.charAt(sourceIndex)) {
                case ' ':
                    sb.append('_');
                    break;
                case '-':
                    if (sourceIndex<identifier.length()-1) {
                        ++sourceIndex;
                        sb.append(Character.toUpperCase(identifier.charAt(sourceIndex)));
                    } else {
                        sb.append('-');
                    }
                    break;
                case '0':
                    // appendCapitalized(sb, 'o');
                    appendNormally(sb, 'o');
                    break;
                case '1':
                    // appendCapitalized(sb, 'l');
                    appendNormally(sb, 'l');
                    break;
                case '3':
                    // appendCapitalized(sb, 'e');
                    appendNormally(sb, 'e');
                    break;
                case '4':
                    // appendCapitalized(sb, 'a');
                    appendNormally(sb, 'a');
                    break;
                case '7':
                    // appendCapitalized(sb, 't');
                    appendNormally(sb, 't');
                    break;
                default:
                    char ch;
                    if (Character.isLetter(ch = identifier.charAt(sourceIndex))) {
                        // appendCapitalized(sb, ch);
                        appendNormally(sb, ch);
                    }
                    // else: omit appending
            }
        }
        return sb.toString();
    }
}
