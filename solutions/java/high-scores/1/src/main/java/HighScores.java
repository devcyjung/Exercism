import java.util.Comparator;
import java.util.List;
import java.util.stream.Collectors;

class HighScores {
    private List<Integer> highScores;

    public HighScores(List<Integer> highScores) {
        this.highScores = highScores;
    }

    List<Integer> scores() {
        return highScores;
    }

    Integer latest() {
        return highScores.getLast();
    }

    Integer personalBest() {
        return highScores.stream().max(Comparator.naturalOrder()).orElse(0);
    }

    List<Integer> personalTopThree() {
        return highScores.stream().sorted(Comparator.reverseOrder()).limit(3).collect(Collectors.toList());
    }

}
