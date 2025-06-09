class IsogramChecker {

    boolean isIsogram(String phrase) {
        var seen = new boolean[26];
        return phrase.toLowerCase().chars()
            .filter(c -> c != ' ' && c != '-')
            .allMatch(c -> {
                if (!Character.isLetter(c)) {
                    return false;
                }
                if (seen[c - 'a']) {
                    return false;
                }
                seen[c - 'a'] = true;
                return true;
            });
    }

}
