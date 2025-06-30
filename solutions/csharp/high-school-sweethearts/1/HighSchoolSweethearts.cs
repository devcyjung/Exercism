using System.Globalization;

public static class HighSchoolSweethearts
{
    public static string DisplaySingleLine(string studentA, string studentB) =>
        $"{studentA, 29} ♡ {studentB, -29}";

    public static string DisplayBanner(string studentA, string studentB) =>
        $@"     ******       ******
   **      **   **      **
 **         ** **         **
**            *            **
**                         **
**     {studentA} +  {studentB}    **
 **                       **
   **                   **
     **               **
       **           **
         **       **
           **   **
             ***
              *";

    private static readonly CultureInfo GermanLocale = new("de-DE");
    
    public static string DisplayGermanExchangeStudents(
        string studentA, string studentB, DateTime start, float hours
    ) => $"{studentA} and {studentB} have been dating since {
        start.ToString("d", GermanLocale)
    } - that's {
        hours.ToString("N2", GermanLocale)
    } hours";
}