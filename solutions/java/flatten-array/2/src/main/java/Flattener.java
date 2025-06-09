import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public final class Flattener {

    public static final List<Object> flatten(List<?> list) {
        return list.stream().flatMap(e -> switch(e) {
            case null -> Stream.empty();
            case List<?> l -> flatten(l).stream();
            default -> Stream.of(e);
        }).collect(Collectors.toList());
    }

}