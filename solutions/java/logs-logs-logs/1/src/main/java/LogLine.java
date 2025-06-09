import java.util.regex.Matcher;
import java.util.regex.Pattern;

public final class LogLine {
    private static final Pattern PATTERN = Pattern.compile("\\[(.+?)\\]:\\s(.+)");
    private final LogLevel level;
    private final String message;
    
    public LogLine(String logLine) {
        var matcher = PATTERN.matcher(logLine);
        if (matcher.matches()) {
            var lvl = matcher.group(1);
            this.message = matcher.group(2);
            this.level = switch (lvl) {
                case "TRC" -> LogLevel.TRACE;
                case "DBG" -> LogLevel.DEBUG;
                case "INF" -> LogLevel.INFO;
                case "WRN" -> LogLevel.WARNING;
                case "ERR" -> LogLevel.ERROR;
                case "FTL" -> LogLevel.FATAL;
                default -> LogLevel.UNKNOWN;
            };
        } else {
            this.level = LogLevel.UNKNOWN;
            this.message = logLine;
        }
    }

    public LogLevel getLogLevel() {
        return level;
    }

    public String getOutputForShortLog() {
        return level.toString() + ":" + message;
    }
}