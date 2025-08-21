import io.reactivex.Observable;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.stream.IntStream;

class Hangman {
    static final Part[] PARTS = Part.values();

    Observable<Output> play(Observable<String> words, Observable<String> letters) {
        var output = new Output[]{Output.empty()};
        return Observable.merge(words, letters).map(s -> {
            if (s.length() == 1) {
                var guess = output[0].guess;
                var misses = output[0].misses;
                var status = output[0].status;
                if (status != Status.PLAYING) {
                    return output[0];
                }
                if (guess.contains(s) || misses.contains(s)) {
                    throw new IllegalArgumentException("Letter " + s + " was already played");
                }
                var newGuess = s.charAt(0);
                var secret = output[0].secret;
                var discovered = IntStream.range(0, secret.length())
                    .mapToObj(i -> secret.charAt(i) == newGuess ? newGuess : output[0].discovered.charAt(i))
                    .collect(StringBuilder::new, StringBuilder::append, StringBuilder::append).toString();
                var parts = output[0].parts;
                if (secret.contains(s)) {
                    guess = new HashSet<>(guess);
                    guess.add(s);
                    status = discovered.contains("_") ? Status.PLAYING : Status.WIN;
                } else {
                    misses = new HashSet<>(misses);
                    misses.add(s);
                    parts = new ArrayList<>(parts);
                    parts.add(PARTS[parts.size()]);
                    status = parts.size() < PARTS.length ? Status.PLAYING : Status.LOSS;
                }
                output[0] = new Output(secret, discovered, guess, misses, parts, status);
            } else {
                output[0] = new Output(s, "_".repeat(s.length()), Set.of(), Set.of(), List.of(), Status.PLAYING);
            }
            return output[0];
        });
    }
}
