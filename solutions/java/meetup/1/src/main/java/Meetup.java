import java.time.DayOfWeek;
import java.time.LocalDate;

record Meetup(int month, int year) {
    LocalDate day(DayOfWeek dayOfWeek, MeetupSchedule schedule) {
        var base = startDate(schedule);
        var delta = Math.floorMod(dayOfWeek.getValue() - base.getDayOfWeek().getValue(), 7);
        return base.plusDays(delta);
    }

    private LocalDate startDate(MeetupSchedule schedule) {
        return switch (schedule) {
            case FIRST -> LocalDate.of(year(), month(), 1);
            case SECOND -> LocalDate.of(year(), month(), 8);
            case THIRD -> LocalDate.of(year(), month(), 15);
            case FOURTH -> LocalDate.of(year(), month(), 22);
            case TEENTH -> LocalDate.of(year(), month(), 13);
            case LAST -> LocalDate.of(year(), month(), 1).plusMonths(1).minusDays(7);
        };
    }
}