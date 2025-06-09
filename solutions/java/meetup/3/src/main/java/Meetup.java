import java.time.DayOfWeek;
import java.time.LocalDate;
import java.time.temporal.TemporalAdjusters;

record Meetup(int month, int year) {
    LocalDate day(DayOfWeek dayOfWeek, MeetupSchedule schedule) {
        var base = switch(schedule) {
            case TEENTH -> LocalDate.of(year, month, 13);
            default -> LocalDate.of(year, month, 1);
        };
        var modifier = switch(schedule) {
            case FIRST -> TemporalAdjusters.dayOfWeekInMonth(1, dayOfWeek);
            case SECOND -> TemporalAdjusters.dayOfWeekInMonth(2, dayOfWeek);
            case THIRD -> TemporalAdjusters.dayOfWeekInMonth(3, dayOfWeek);
            case FOURTH -> TemporalAdjusters.dayOfWeekInMonth(4, dayOfWeek);
            case LAST -> TemporalAdjusters.dayOfWeekInMonth(-1, dayOfWeek);
            case TEENTH -> TemporalAdjusters.nextOrSame(dayOfWeek);
        };
        return base.with(modifier);
    }
}