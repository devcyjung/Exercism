public static class Languages
{
    public static List<string> NewList() => new();

    public static List<string> GetExistingLanguages() => new() {"C#", "Clojure", "Elm"};

    public static List<string> AddLanguage(List<string> languages, string language)
    {
        languages.Add(language);
        return languages;
    }

    public static int CountLanguages(List<string> languages) => languages.Count;

    public static bool HasLanguage(List<string> languages, string language) =>
        languages.Contains(language);

    public static List<string> ReverseList(List<string> languages)
    {
        languages.Reverse();
        return languages;
    }

    public static bool IsExciting(List<string> languages) => languages.Count switch
    {
        > 0 when languages[0].Equals("C#") => true,
        2 or 3 when languages[1].Equals("C#") => true,
        _ => false,
    };

    public static List<string> RemoveLanguage(List<string> languages, string language)
    {
        languages.RemoveAll(l => l.Equals(language));
        return languages;
    }

    public static bool IsUnique(List<string> languages) =>
        languages.Count == new HashSet<string>(languages).Count;
}