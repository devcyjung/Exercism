import java.util.Map;
import java.util.stream.Collectors;

class RnaTranscription {

    private static Map<Integer, Integer> codePointMapping = Map.ofEntries(
        Map.entry("A".codePointAt(0), "U".codePointAt(0)),
        Map.entry("T".codePointAt(0), "A".codePointAt(0)),
        Map.entry("C".codePointAt(0), "G".codePointAt(0)),
        Map.entry("G".codePointAt(0), "C".codePointAt(0))
    );

    String transcribe(String dnaStrand) {
        return dnaStrand.codePoints()
            .map(cp -> codePointMapping.getOrDefault(cp, cp))
            .mapToObj(cp -> new String(Character.toChars(cp)))
            .collect(Collectors.joining());
    }

}
