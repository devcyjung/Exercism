import java.time.LocalDate
import java.time.LocalDateTime
import java.time.LocalTime
import java.time.temporal.ChronoUnit
import java.time.temporal.Temporal
import java.time.temporal.TemporalQueries

class Gigasecond(val now: Temporal) {
    val localDate: LocalDate? = now.query(TemporalQueries.localDate())
    val localTime: LocalTime? = now.query(TemporalQueries.localTime())
    val localDateTime: LocalDateTime =
        if (localDate != null && localTime != null) {
            localDate.atTime(localTime)
        } else if (localDate != null) {
            localDate.atStartOfDay()
        } else if (localTime != null) {
            localTime.atDate(LocalDate.now())
        } else {
            LocalDateTime.now()
        }
    val date: LocalDateTime
        get() = localDateTime.plus(1_000_000_000L, ChronoUnit.SECONDS)
}