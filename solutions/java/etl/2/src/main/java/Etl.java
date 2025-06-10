import java.util.Collection;
import java.util.Map;
import java.util.stream.Collectors;

class Etl {
    Map<String, Integer> transform(Map<Integer, ? extends Collection<String>> old) {
        return old.entrySet().stream()
            .flatMap(e -> e.getValue().stream()
                .map(str -> Map.entry(str.toLowerCase(), e.getKey()))
            ).collect(Collectors.toMap(Map.Entry::getKey, Map.Entry::getValue));
    }
}