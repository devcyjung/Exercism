public static class LogAnalysis 
{
    public static string SubstringAfter(this string str, string begin) =>
        str.Substring(str.IndexOf(begin) + begin.Length);

    public static string SubstringBetween(this string str, string begin, string end)
    {
        var beginIdx = str.IndexOf(begin) + begin.Length;
        return str.Substring(beginIdx, str.IndexOf(end) - beginIdx);
    }
    
    public static string Message(this string str) => str.SubstringAfter("]: ").Trim();

    public static string LogLevel(this string str) => str.SubstringBetween("[", "]:");
}