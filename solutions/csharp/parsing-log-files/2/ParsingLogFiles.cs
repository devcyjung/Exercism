using System.Text.RegularExpressions;

public class LogParser
{
    private readonly Regex validLine = new(@"^\[(TRC|DBG|INF|WRN|ERR|FTL)\]");
    public bool IsValidLine(string text) => validLine.IsMatch(text);

    private readonly Regex delimiter = new(@"<[\^*=-]+>");
    public string[] SplitLogLine(string text) => delimiter.Split(text);

    private readonly Regex password = new(@""".*password.*""", RegexOptions.IgnoreCase);
    public int CountQuotedPasswords(string lines) => lines.Split(Environment.NewLine)
        .Count(line => password.IsMatch(line));

    private readonly Regex eol = new(@"end-of-line[0-9]+");
    public string RemoveEndOfLineText(string line) => eol.Replace(line, "");

    private readonly Regex offending = new(@"\bpassword\S+", RegexOptions.IgnoreCase);
    public string[] ListLinesWithPasswords(string[] lines) => lines.Select(line =>
        {
            Match match = offending.Match(line);
            return match.Success ? $"{match.Value}: {line}" : $"--------: {line}";
        }).ToArray();
}