import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;
import java.time.temporal.ChronoUnit;
import java.time.temporal.Temporal;
import java.time.temporal.TemporalQueries;

public class Gigasecond {
    private final LocalDateTime localDateTime;
    
    public Gigasecond(Temporal now) {
        var localDate = now.query(TemporalQueries.localDate());
        var localTime = now.query(TemporalQueries.localTime());
        this.localDateTime = switch (localDate) {
            case null -> switch (localTime) {
                case null -> LocalDateTime.now();
                default -> localTime.atDate(LocalDate.now());
            };
            default -> switch (localTime) {
                case null -> localDate.atStartOfDay();
                default -> localDate.atTime(localTime);
            };
        };
    }

    public LocalDateTime getDateTime() {
        return localDateTime.plus(1_000_000_000L, ChronoUnit.SECONDS);
    }
}