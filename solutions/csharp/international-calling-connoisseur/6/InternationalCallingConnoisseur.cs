using CodeName = System.Collections.Generic.Dictionary<int, string>;

public static class DialingCodes
{
    public static CodeName GetEmptyDictionary() => new();

    public static CodeName GetExistingDictionary() =>
        new(){[1]="United States of America", [55]="Brazil", [91]="India"};

    public static CodeName AddCountryToEmptyDictionary(int code, string name) => new(){[code]=name};
    
    public static CodeName AddCountryToExistingDictionary(CodeName dict, int code, string name) =>
        (dict[code] = name) is string ? dict : dict;

    public static string GetCountryNameFromDictionary(CodeName dict, int code) =>
        dict.GetValueOrDefault(code, string.Empty);

    public static bool CheckCodeExists(CodeName dict, int code) => dict.ContainsKey(code);

    public static CodeName UpdateDictionary(CodeName dict, int code, string name) =>
        (dict.ContainsKey(code)) ? AddCountryToExistingDictionary(dict, code, name) : dict;

    public static CodeName RemoveCountryFromDictionary(CodeName dict, int code) =>
        dict.Remove(code) is bool ? dict : dict;
    
    public static string FindLongestCountryName(CodeName dict) =>
        dict.Values.MaxBy(name => name.Length) ?? string.Empty;
}