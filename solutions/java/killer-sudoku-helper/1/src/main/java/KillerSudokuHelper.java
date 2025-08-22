import java.util.List;
import java.util.stream.IntStream;
import java.util.stream.Stream;

public class KillerSudokuHelper {
    static List<List<Integer>> combinationsInCage(Integer cageSum, Integer cageSize, List<Integer> exclude) {
        return combinationsInCage(cageSum, cageSize).stream()
            .filter(list -> exclude.stream().noneMatch(list::contains)).toList();
    }

    static List<List<Integer>> combinationsInCage(Integer cageSum, Integer cageSize) {
        return combination(IntStream.range(1, 10).boxed().toList(), cageSize).stream()
            .filter(list -> list.stream().reduce(0, Integer::sum).equals(cageSum)).toList();
    }

    private static <T> List<List<T>> combination(List<T> list, int count) {
        if (count < 0)
            throw new IllegalArgumentException("count must be non-negative");
        return switch (count) {
            case 0 -> List.of();
            case 1 -> list.stream().map(List::of).toList();
            default -> IntStream.range(0, list.size()).boxed()
                    .flatMap(i -> combination(list.subList(i + 1, list.size()), count - 1).stream()
                        .map(com -> Stream.concat(Stream.of(list.get(i)), com.stream()).toList())
                    ).toList();
        };
    }
}
