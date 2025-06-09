import java.time.temporal.*
import java.time.*

class Gigasecond(val now: Temporal) {
    private val dateTimeNow = when (now) {
        is Instant -> now.atZone(ZoneId.systemDefault()).toLocalDateTime()
        is LocalDate -> now.atStartOfDay()
        is LocalDateTime -> now
        is LocalTime -> now.atDate(LocalDate.now())
        is OffsetDateTime -> now.toLocalDateTime()
        is OffsetTime -> now.atDate(LocalDate.now()).toLocalDateTime()
        is Year -> now.atMonthDay(MonthDay.of(Month.JANUARY, 1)).atStartOfDay()
        is YearMonth -> now.atDay(1).atStartOfDay()
        is ZonedDateTime -> now.toLocalDateTime()
        else -> {
            if (now.isSupported(ChronoField.INSTANT_SECONDS)) {
                LocalDateTime.ofEpochSecond(
                    now.getLong(ChronoField.INSTANT_SECONDS),
                    0,
                    ZoneOffset.of(ZoneId.systemDefault().getId()),
                )
            } else if (now.isSupported(ChronoField.EPOCH_DAY)) {
                LocalDate.ofEpochDay(
                    now.getLong(ChronoField.EPOCH_DAY)
                ).atStartOfDay()
            } else if (now.isSupported(ChronoField.SECOND_OF_DAY)) {
                LocalTime.ofSecondOfDay(
                    now.getLong(ChronoField.SECOND_OF_DAY)
                ).atDate(LocalDate.now())
            } else {
                throw IllegalArgumentException(
                    "Unsupported time unit: ${now::class.simpleName}"
                )
            }
        }
    }
    val date: LocalDateTime
        get() = dateTimeNow.plus(1e9.toLong(), ChronoUnit.SECONDS)
}