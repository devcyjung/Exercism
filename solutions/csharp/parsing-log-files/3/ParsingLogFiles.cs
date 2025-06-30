using System.Text.RegularExpressions;

public class LogParser
{
    private readonly Regex validLine = new(@"^\[(TRC|DBG|INF|WRN|ERR|FTL)\]", RegexOptions.Compiled);
    public bool IsValidLine(string text) => validLine.IsMatch(text);

    private readonly Regex delimiter = new(@"<[\^*=-]+>", RegexOptions.Compiled);
    public string[] SplitLogLine(string text) => delimiter.Split(text);

    private readonly Regex password = new(@""".*password.*""", RegexOptions.IgnoreCase | RegexOptions.Compiled);
    public int CountQuotedPasswords(string lines) => lines.Split(Environment.NewLine)
        .Count(line => password.IsMatch(line));

    private readonly Regex eol = new(@"end-of-line[0-9]+", RegexOptions.Compiled);
    public string RemoveEndOfLineText(string line) => eol.Replace(line, "");

    private readonly Regex offending = new(@"\bpassword\S+", RegexOptions.IgnoreCase | RegexOptions.Compiled);
    public string[] ListLinesWithPasswords(string[] lines) => lines.Select(line =>
        {
            Match match = offending.Match(line);
            return match.Success ? $"{match.Value}: {line}" : $"--------: {line}";
        }).ToArray();
}