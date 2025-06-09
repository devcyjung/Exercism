import java.time.temporal.Temporal
import java.time.temporal.ChronoUnit
import java.time.LocalDate
import java.time.LocalTime
import java.time.Instant
import java.time.ZoneId
import java.time.ZonedDateTime
import java.time.OffsetDateTime
import java.time.OffsetTime
import java.time.LocalDateTime
import java.time.Year
import java.time.Month
import java.time.YearMonth
import java.time.MonthDay

class Gigasecond(val now: Temporal) {
    private val dateTimeNow = when (now) {
        is LocalDate -> now.atStartOfDay()
        is LocalDateTime -> now
        is LocalTime -> now.atDate(LocalDate.now())
        is Instant -> now.atZone(ZoneId.of("Z")).toLocalDateTime()
        is ZonedDateTime -> now.toLocalDateTime()
        is OffsetDateTime -> now.toLocalDateTime()
        is OffsetTime -> now.atDate(LocalDate.now()).toLocalDateTime()
        is Year -> now.atMonthDay(MonthDay.of(Month.JANUARY, 1)).atStartOfDay()
        is YearMonth -> now.atDay(1).atStartOfDay()
        else -> throw IllegalArgumentException(
            "Unsupported Temporal type: ${now::class.simpleName}"
        )
    }
    val date: LocalDateTime
        get() = dateTimeNow.plus(1_000_000_000, ChronoUnit.SECONDS)
}