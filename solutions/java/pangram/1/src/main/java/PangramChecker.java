import java.util.HashMap;
import java.util.Map;
import java.util.stream.IntStream;

public class PangramChecker {

    public boolean isPangram(String input) {
        Map<Integer, Integer> map = input.toLowerCase().chars()
            .filter(Character::isLetter)
            .collect(
                HashMap::new,
                (m, c) -> m.put(c, m.getOrDefault(c, 0) + 1),
                HashMap::putAll
            );
        return IntStream.rangeClosed('a', 'z').allMatch(c -> map.containsKey(c));
    }

}
