import java.util.Arrays;
import java.util.List;

class HandshakeCalculator {

    private static Signal[] SIGNALS = Signal.values();

    List<Signal> calculateHandshake(int number) {
        Signal[] signals = new Signal[5];
        int index = 0;
        for (int i = 0; i < SIGNALS.length; ++i) {
            if (((number >> i) & 1) == 1) {
                signals[index++] = SIGNALS[i];
            }
        }
        if (((number >> SIGNALS.length) & 1) == 1) {
            for (int start = 0, end = index - 1; start < end; ++start, --end) {
                Signal temp = signals[end];
                signals[end] = signals[start];
                signals[start] = temp;
            }
        }
        return Arrays.asList(Arrays.copyOf(signals, index));
    }

}
