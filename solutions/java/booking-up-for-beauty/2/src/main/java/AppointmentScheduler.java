import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.Month;
import java.time.MonthDay;
import java.time.format.DateTimeFormatter;

class AppointmentScheduler {
    private static final DateTimeFormatter appointmentFormat = DateTimeFormatter.ofPattern("MM/dd/yyyy HH:mm:ss");
    
    public static LocalDateTime schedule(String appointmentDateDescription) {
        return LocalDateTime.parse(appointmentDateDescription, appointmentFormat);
    }

    public static boolean hasPassed(LocalDateTime appointmentDate) {
        return appointmentDate.isBefore(LocalDateTime.now());
    }

    public boolean isAfternoonAppointment(LocalDateTime appointmentDate) {
        var hour = appointmentDate.getHour();
        return 12 <= hour && hour < 18;
    }

    private static final DateTimeFormatter descriptionFormat = DateTimeFormatter.ofPattern("'You have an appointment on' EEEE, MMMM d, yyyy, 'at' h:mm a.");
    
    public static String getDescription(LocalDateTime appointmentDate) {
        return descriptionFormat.format(appointmentDate);
    }

    private static final MonthDay anniversaryDate = MonthDay.of(Month.SEPTEMBER, 15);
    
    public static LocalDate getAnniversaryDate() {
        return LocalDate.now().with(anniversaryDate);
    }
}