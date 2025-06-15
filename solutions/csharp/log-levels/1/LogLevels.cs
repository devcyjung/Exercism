static class LogLine
{   
    public static string Message(string logLine) => logLine.Substring(logLine.IndexOf("]:") + 2).Trim();

    public static string LogLevel(string logLine)
    {
        int begin = logLine.IndexOf("[") + 1;
        return logLine.Substring(begin, logLine.IndexOf("]:") - begin).ToLower();
    }

    public static string Reformat(string logLine) => $"{Message(logLine)} ({LogLevel(logLine)})";
}
